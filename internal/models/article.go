package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleList struct {
	Total int
	Items []ArticleItem
}

type ArticleItem struct {
	ID      int64
	Title   string
	Image   LocalFile
	Summary string
	Date    time.Time
}

type ArticleListFilters struct {
	Term string
}

func (lf *ArticleListFilters) FromRequest(c *gin.Context) *ArticleListFilters {
	lf.Term = c.Query("searchTerm")
	return lf
}

func (lf *ArticleListFilters) Empty() bool {
	return lf.Term == ""
}
