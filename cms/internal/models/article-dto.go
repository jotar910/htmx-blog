package models

import (
	cerrors "github.com/jotar910/buzzer-cms/pkg/errors"
	"github.com/jotar910/buzzer-cms/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleList struct {
	Total int
	Items []ArticleItem
}

type ArticleItem struct {
	ID       int64
	Title    string
	Filename string
	Date     time.Time
	Carousel bool
}

type ArticleListFilters struct {
	ID       string `form:"id"`
	Title    string `form:"title"`
	Filename string `form:"filename"`
	Date     string `form:"date"`
}

func (lf *ArticleListFilters) Empty() bool {
	return lf.ID == "" && lf.Title == "" && lf.Filename == "" && lf.Date == ""
}

func (lf *ArticleListFilters) Decode(c *gin.Context) (*ArticleListFilters, error) {
	if err := c.Bind(lf); err != nil {
		logger.L.Errorf("parsing article list filter: %v", err)
		return nil, cerrors.Wrap(err, cerrors.BadRequest)
	}
	return lf, nil
}

type Article struct {
	ID       int64
	Title    string
	Summary  string
	Image    LocalFile
	Filename string
	Date     time.Time
	// Customizations
	Carousel bool
}

type ArticlePatch struct {
	Carousel *bool `form:"carousel"`
}

func (lf *ArticlePatch) Decode(c *gin.Context) (*ArticlePatch, error) {
	if err := c.Bind(lf); err != nil {
		logger.L.Errorf("parsing article patch payload: %v", err)
		return nil, cerrors.Wrap(err, cerrors.BadRequest)
	}
	return lf, nil
}
