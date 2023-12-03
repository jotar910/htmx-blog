package models

import (
	"net/url"
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

func (lf *ArticleListFilters) Empty() bool {
	return lf.Term == ""
}

func (lf *ArticleListFilters) Encode() string {
	v := url.Values{}
	if lf.Term != "" {
		v.Set("searchTerm", lf.Term)
	}
	return v.Encode()
}

func (lf *ArticleListFilters) Decode(c *gin.Context) *ArticleListFilters {
	lf.Term = c.Query("searchTerm")
	return lf
}

type Article struct {
	ID       int64
	Title    string
	Date     time.Time
	Filename string
}
