package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jotar910/buzzer-cms/internal/components"
	components_core "github.com/jotar910/buzzer-cms/internal/components/core"
	"github.com/jotar910/buzzer-cms/internal/models"
	"github.com/jotar910/buzzer-cms/internal/services"
	cerrors "github.com/jotar910/buzzer-cms/pkg/errors"
	"net/http"
	"strconv"
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

		homepage := components.Homepage(articlesTableTempl)
		pageLayout := components_core.DefaultPageLayout(homepage)
		render(c, pageLayout)
	})

	r.POST("/search", func(c *gin.Context) {
		filters, err := new(models.ArticleListFilters).Decode(c)
		if err != nil {
			handleError(c, err)
			return
		}

		articlesList, err := ph.postsService.GetList(filters)
		if err != nil {
			handleError(c, err)
			return
		}
		articlesTableTempl := components.ArticlesTableSearchResult(articlesList)
		render(c, articlesTableTempl)
	})

	r.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			handleError(c, cerrors.Wrap(err, cerrors.NotFound))
			return
		}

		post, err := ph.postsService.GetPostById(id)
		if err != nil {
			handleError(c, err)
			return
		}

		article := components.ArticleDetails(post)
		pageLayout := components_core.DefaultPageLayout(article)
		render(c, pageLayout)
	})

	r.GET("/:id/content", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			handleError(c, cerrors.Wrap(err, cerrors.NotFound))
			return
		}
		typ := c.DefaultQuery("type", "post")

		post, err := ph.postsService.GetPostById(id)
		if err != nil {
			handleError(c, err)
			return
		}

		article := components.ArticleDetailsContent(post.ID, post.Filename, typ == "html")
		render(c, article)
	})

	r.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			handleError(c, cerrors.Wrap(err, cerrors.NotFound))
			return
		}
		payload, err := new(models.ArticlePatch).Decode(c)
		if err != nil {
			handleError(c, err)
			return
		}

		_, err = ph.postsService.PatchPost(id, payload)
		if err != nil {
			handleError(c, err)
			return
		}

		c.Status(http.StatusOK)
	})
}
