package handlers

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/jotar910/htmx-templ/internal/components"
	components_core "github.com/jotar910/htmx-templ/internal/components/core"
	"github.com/jotar910/htmx-templ/internal/models"
	"github.com/jotar910/htmx-templ/internal/services"
)

type PostsHandler struct {
	postsService *services.PostsService
}

func NewPostsHandler(postsService *services.PostsService) *PostsHandler {
	return &PostsHandler{
		postsService: postsService,
	}
}

func render(c *gin.Context, html templ.Component) {
	if c.GetHeader("HX-Request") == "" {
		// This means it's the initial full page load
		// Run your specific middleware logic here
		// For example, initializing session data, etc.

		// Log for demonstration purposes
		c.HTML(http.StatusOK, "", components_core.Index(html))
	} else {
		c.HTML(http.StatusOK, "", html)
	}
}

func (ph *PostsHandler) RegisterPosts(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		filters := new(models.ArticleListFilters).Decode(c)
		list := ph.postsService.GetList(filters)
		articlesList := components.ArticlesListContainer(list, filters)
		mostSeen := components.MostSeenContainer(list.Items)
		highlights := components.HighlightsContainer(
			&list.Items[0],
			&list.Items[1],
			&list.Items[2],
		)
		recentList := components.RecentListContainer(list.Items)
		articlesCarousel := components.ArticlesCarousel(list.Items)
		homepage := components.Homepage(
			articlesCarousel,
			recentList,
			highlights,
			mostSeen,
			articlesList,
		)
		render(c, homepage)
	})

	r.GET("/filtered", func(c *gin.Context) {
		term := c.Query("searchTerm")
		filters := &models.ArticleListFilters{Term: term}
		list := ph.postsService.GetList(filters)

		if filters.Empty() {
			c.Header("HX-Push-Url", "./#articles")
		} else {
			c.Header("HX-Push-Url", "./?"+filters.Encode()+"#articles")
		}

		componentList := components.ArticlesListItemsResponse(list.Items)
		c.HTML(http.StatusOK, "", componentList)

		componentCounter := components.ArticlesListCountResponse(list.Total, filters.Term)
		c.HTML(http.StatusOK, "", componentCounter)
	})

	r.GET("/:id", func(c *gin.Context) {
		arg := &models.Article{
			ID:       1,
			Title:    "Testing this",
			Date:     time.Now(),
			Filename: "markdown.md",
		}
		filters := new(models.ArticleListFilters).Decode(c)
		list := ph.postsService.GetList(filters)
		article := components.ArticleDetails(
			arg,
			components.ArticleOption{
				Component: components.ArticlesLinksList(list.Items[:3]),
				Area:      "header",
			},
			components.ArticleOption{
				Component: components.RelatedVerticalContainer(list.Items),
				Area:      "aside",
			},
		)
		render(c, article)
	})
}
