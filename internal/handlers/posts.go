package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jotar910/htmx-templ/internal/models"
	"github.com/jotar910/htmx-templ/internal/services"
	"github.com/jotar910/htmx-templ/views/components"
	components_articlescarousel "github.com/jotar910/htmx-templ/views/components/articles-carousel"
	components_articleslist "github.com/jotar910/htmx-templ/views/components/articles-list"
	components_highlights "github.com/jotar910/htmx-templ/views/components/highlights"
	components_mostseen "github.com/jotar910/htmx-templ/views/components/most-seen"
	components_recentlist "github.com/jotar910/htmx-templ/views/components/recent-list"
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
		articlesList := components_articleslist.ArticlesContainer(list, filters)
		mostSeen := components_mostseen.MostSeenContainer(list.Items)
		highlights := components_highlights.HighlightsContainer(
			&list.Items[0],
			&list.Items[1],
			&list.Items[2],
		)
		recentList := components_recentlist.RecentListContainer(list.Items)
		articlesCarousel := components_articlescarousel.ArticlesCarousel(list.Items)
		homepage := components.Homepage(
			articlesCarousel,
			recentList,
			highlights,
			mostSeen,
			articlesList,
		)
		component := components.Index(homepage)
		c.HTML(http.StatusOK, "", component)
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

		componentList := components_articleslist.ArticlesItemsResponse(list.Items)
		c.HTML(http.StatusOK, "", componentList)

		componentCounter := components_articleslist.ArticlesCountResponse(list.Total, filters.Term)
		c.HTML(http.StatusOK, "", componentCounter)
	})
}
