package handlers

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	cerrors "github.com/jotar910/htmx-templ/pkg/errors"
	"net/http"
	"strconv"

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
		list, err := ph.postsService.GetList(filters)
		if err != nil {
			handleError(c, err)
			return
		}
		articlesList := components.ArticlesListContainer(list, filters)

		mostSeenList, err := ph.postsService.GetMostSeenPosts()
		if err != nil {
			handleError(c, err)
			return
		}
		mostSeen := components.MostSeenContainer(mostSeenList)

		highlightsList, err := ph.postsService.GetHighlightPosts()
		if err != nil {
			handleError(c, err)
			return
		}
		var highlights templ.Component
		if len(highlightsList) > 2 {
			highlights = components.HighlightsContainer(
				&highlightsList[0],
				&highlightsList[1],
				&highlightsList[2],
			)
		}

		recent, err := ph.postsService.GetRecentPosts()
		if err != nil {
			handleError(c, err)
			return
		}
		recentList := components.RecentListContainer(recent)

		carousel, err := ph.postsService.GetCarouselPosts()
		if err != nil {
			handleError(c, err)
			return
		}
		articlesCarousel := components.ArticlesCarousel(carousel)

		homepage := components.Homepage(
			articlesCarousel,
			recentList,
			highlights,
			mostSeen,
			articlesList,
		)
		render(c, components_core.DefaultPageLayout(homepage))
	})

	r.GET("/filtered", func(c *gin.Context) {
		filters := new(models.ArticleListFilters).Decode(c)
		list, err := ph.postsService.GetList(filters)
		if err != nil {
			handleError(c, err)
			return
		}

		if filters.Empty() {
			c.Header("HX-Push-Url", "./")
		} else {
			c.Header("HX-Push-Url", "./?"+filters.Encode())
		}

		componentList := components.ArticlesListItemsResponse(list.Items)
		render(c, componentList)

		componentCounter := components.ArticlesListCountResponse(list.Total, filters.Term)
		render(c, componentCounter)
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

		pageOptions := make([]components.ArticleOption, 0)

		articleLinks, err := ph.postsService.GetHighlightPosts()
		if err != nil {
			handleError(c, err)
			return
		}
		if len(articleLinks) > 0 {
			pageOptions = append(pageOptions, components.ArticleOption{
				Component: components.ArticlesLinksList(articleLinks),
				Area:      "header",
			})
		}

		related, err := ph.postsService.GetRelatedPosts(id)
		if err != nil {
			handleError(c, err)
			return
		}
		if len(related) > 0 {
			pageOptions = append(pageOptions, components.ArticleOption{
				Component: components.RelatedVerticalContainer(related),
				Area:      "aside",
			})
		}

		article := components.ArticleDetails(post, pageOptions...)
		render(c, components_core.DefaultPageLayout(article))
	})
}
