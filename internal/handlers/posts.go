package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotar910/htmx-templ/internal/models"
	"github.com/jotar910/htmx-templ/internal/services"
	"github.com/jotar910/htmx-templ/views/components"
	components_articleslist "github.com/jotar910/htmx-templ/views/components/articles-list"
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
		filters := new(models.ArticleListFilters).Decode(c)
		list := ph.postsService.GetList(filters)
		component := components.Index(components_articleslist.ArticlesContainer(list, filters))
		c.HTML(http.StatusOK, "", component)
	})

	r.GET("/filtered", func(c *gin.Context) {
		term := c.Query("searchTerm")
		filters := &models.ArticleListFilters{Term: term}
		list := ph.postsService.GetList(filters)

		if filters.Empty() {
			c.Header("HX-Push-Url", "./")
		} else {
			c.Header("HX-Push-Url", "./?"+filters.Encode())
		}

		componentList := components_articleslist.ArticlesItemsResponse(list.Items)
		c.HTML(http.StatusOK, "", componentList)

		componentCounter := components_articleslist.ArticlesCountResponse(list.Total)
		c.HTML(http.StatusOK, "", componentCounter)
	})
}
