package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jotar910/buzzer-cms/internal/components"
	components_core "github.com/jotar910/buzzer-cms/internal/components/core"
	"github.com/jotar910/buzzer-cms/internal/models"
	"github.com/jotar910/buzzer-cms/internal/services"
	"github.com/jotar910/buzzer-cms/pkg/logger"
)

type PostsHandler struct {
	postsService *services.PostsService
}

func NewPostsHandler(postsService *services.PostsService) *PostsHandler {
	return &PostsHandler{
		postsService: postsService,
	}
}

func (ph *PostsHandler) RegisterPosts(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		articlesList, err := ph.postsService.GetList(nil)
		if err != nil {
			handleError(c, err)
			return
		}
		articlesTableTempl := components.ArticlesTable(articlesList)

		pageLayout := components_core.DefaultPageLayout(articlesTableTempl)
		homepage := components.Homepage(pageLayout)
		render(c, homepage)
	})

	r.POST("/search", func(c *gin.Context) {
		filters, err := new(models.ArticleListFilters).Decode(c)
		if err != nil {
			handleError(c, err)
			return
		}
		logger.L.Debugf("%+v", filters)

		articlesList, err := ph.postsService.GetList(filters)
		if err != nil {
			handleError(c, err)
			return
		}
		articlesTableTempl := components.ArticlesTableSearchResult(articlesList)
		render(c, articlesTableTempl)
	})
}
