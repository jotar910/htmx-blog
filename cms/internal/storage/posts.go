package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jotar910/buzzer-cms/internal/models"
	cerrors "github.com/jotar910/buzzer-cms/pkg/errors"
	"github.com/jotar910/buzzer-cms/pkg/logger"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (sqldb *SQLiteDatabase) GetPostsList() (*models.ArticleList, error) {
	logger.L.Debug("getting posts list")
	rows, err := sqldb.db.Queryx("select * from articles order by title")
	if err != nil {
		logger.L.Errorf("selecting posts list: %v", err)
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	items, err := scanPostItems(rows)
	if err != nil {
		return nil, errors.Wrap(err, "getting posts list")
	}
	return &models.ArticleList{
		Total: len(items),
		Items: items,
	}, nil
}

func (sqldb *SQLiteDatabase) GetPostsListFiltered(
	filters *models.ArticleListFilters,
) (*models.ArticleList, error) {
	logger.L.Debugf("getting posts list with filtering: %+v", filters)

	type WhereClause struct {
		query string
		arg   any
	}

	whereClauses := make([]WhereClause, 0)
	if filters.ID != "" {
		whereClauses = append(whereClauses, WhereClause{"cast(id as text) like ?", fmt.Sprintf("%%%s%%", filters.ID)})
	}
	if filters.Title != "" {
		whereClauses = append(whereClauses, WhereClause{"lower(title) like ?", fmt.Sprintf("%%%s%%", filters.Title)})
	}
	if filters.Filename != "" {
		whereClauses = append(whereClauses, WhereClause{"lower(filename) like ?", fmt.Sprintf("%%%s%%", filters.Filename)})
	}
	if filters.Date != "" {
		date, err := time.Parse("2006-01-02", filters.Date)
		if err != nil {
			return nil, cerrors.Wrap(err, cerrors.BadRequest)
		}
		whereClauses = append(whereClauses, WhereClause{"timestamp >= ?", date.Unix() * 1000})
		whereClauses = append(whereClauses, WhereClause{"timestamp < ?", date.Add(time.Hour*24).Unix() * 1000})
	}

	query := fmt.Sprintf(`
		select *
		from articles
		where %s
		order by title`,
		strings.Join(mapArray(whereClauses, func(v WhereClause, _ int) string { return v.query }), " and "),
	)

	rows, err := sqldb.db.Queryx(
		query,
		mapArray(whereClauses, func(v WhereClause, _ int) any { return v.arg })...,
	)
	if err != nil {
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	items, err := scanPostItems(rows)
	if err != nil {
		return nil, errors.Wrap(err, "getting posts list with filtering")
	}
	return &models.ArticleList{
		Total: len(items),
		Items: items,
	}, nil
}

func scanPostItems(rows *sqlx.Rows) ([]models.ArticleItem, error) {
	items := make([]models.ArticleItem, 0)
	for rows.Next() {
		item, err := scanPostItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}
	return items, nil
}

func scanPostItem(row StructScanner) (*models.ArticleItem, error) {
	var item models.ArticleEntity
	if err := row.StructScan(&item); err != nil {
		logger.L.Errorf("scanning post item: %v", err)
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	return models.FromArticleEntityToArticleItem(&item), nil
}

type StructScanner interface {
	StructScan(any) error
}

func mapArray[T any, R any](arr []T, cb func(T, int) R) []R {
	res := make([]R, len(arr))
	for i, v := range arr {
		res[i] = cb(v, i)
	}
	return res
}
