package storage

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jotar910/htmx-templ/internal/models"
	cerrors "github.com/jotar910/htmx-templ/pkg/errors"
	"github.com/jotar910/htmx-templ/pkg/logger"
	"github.com/pkg/errors"
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
	if filters.Term == "" {
		logger.L.Debug("getting posts list without filtering")
		return sqldb.GetPostsList()
	}
	logger.L.Debugf("getting posts list with filtering: %+v", filters)

	rows, err := sqldb.db.Queryx(`
		select *
		from articles
		where lower(title) like ?
		order by title`,
		fmt.Sprintf("%%%s%%", filters.Term),
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

func (sqldb *SQLiteDatabase) GetMostSeenPosts() ([]models.ArticleItem, error) {
	logger.L.Debug("getting most seen posts")

	rows, err := sqldb.db.Queryx(`
		select articles.*
		from articles
		inner join (
		    select article_id
		    from article_views
		    order by views desc, article_id asc
			limit 10
		) article_views
		on articles.id = article_views.article_id`,
	)
	if err != nil {
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	items, err := scanPostItems(rows)
	if err != nil {
		return nil, errors.Wrap(err, "getting most seen posts")
	}
	return items, nil
}

func (sqldb *SQLiteDatabase) GetHighlightPosts() ([]models.ArticleItem, error) {
	logger.L.Debug("getting highlight posts")

	rows, err := sqldb.db.Queryx(`
		select articles.*
		from articles
		inner join (
		    select article_highlights.article_id, article_views.views
    		from article_highlights
             left outer join article_views
			 on article_highlights.article_id = article_views.article_id
    		order by article_views.views desc, article_highlights.article_id asc
    		limit 3
		) article_highlights
		on articles.id = article_highlights.article_id`,
	)
	if err != nil {
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	items, err := scanPostItems(rows)
	if err != nil {
		return nil, errors.Wrap(err, "getting highlight posts")
	}
	return items, nil
}

func (sqldb *SQLiteDatabase) GetRecentPosts() ([]models.ArticleItem, error) {
	logger.L.Debug("getting recent posts")

	rows, err := sqldb.db.Queryx(`
		select articles.*
		from articles
		order by timestamp desc
		limit 10`,
	)
	if err != nil {
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	items, err := scanPostItems(rows)
	if err != nil {
		return nil, errors.Wrap(err, "getting recent posts")
	}
	return items, nil
}

func (sqldb *SQLiteDatabase) GetCarouselPosts() ([]models.ArticleItem, error) {
	logger.L.Debug("getting carousel posts")

	rows, err := sqldb.db.Queryx(`
		select articles.*
		from articles
		where id in (
		    select article_id
		    from article_carousel
		)
		order by timestamp desc`,
	)
	if err != nil {
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	items, err := scanPostItems(rows)
	if err != nil {
		return nil, errors.Wrap(err, "getting carousel posts")
	}
	return items, nil
}

func (sqldb *SQLiteDatabase) GetRelatedPosts(id int) ([]models.ArticleItem, error) {
	logger.L.Debugf("getting related posts to %d", id)

	rows, err := sqldb.db.Queryx(`
		select articles.*
		from articles
		where
			id != $0
			and id in (
				select distinct article_tags.article_id
				from article_tags
				inner join (
					select article_tags.*
					from articles
					inner join article_tags
					on articles.id = article_tags.article_id
					where articles.id = $0
				) as related_article_tags
				on article_tags.tag_name = related_article_tags.tag_name
				limit 10
			)`,
		id,
	)
	if err != nil {
		return nil, cerrors.Wrap(err, cerrors.InternalServerError)
	}
	items, err := scanPostItems(rows)
	if err != nil {
		return nil, errors.Wrap(err, "getting related posts")
	}
	return items, nil
}

func (sqldb *SQLiteDatabase) GetPostById(id int) (*models.Article, error) {
	logger.L.Debugf("getting post by id %d", id)

	entity := new(models.ArticleEntity)
	err := sqldb.db.Get(
		entity,
		`select articles.*
		from articles
		where id = ?
		LIMIT 1`,
		id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, cerrors.Wrap(errors.Wrap(err, "getting post by id"), cerrors.NotFound)
		}
		return nil, cerrors.Wrap(errors.Wrap(err, "getting post by id"), cerrors.InternalServerError)
	}
	return models.FromArticleEntityToArticle(entity), nil
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
