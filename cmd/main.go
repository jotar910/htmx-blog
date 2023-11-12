package main

import (
	_ "github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/jotar910/htmx-templ/internal/handlers"
	"github.com/jotar910/htmx-templ/internal/services"
	"github.com/jotar910/htmx-templ/internal/storage"
)

func main() {
	r := gin.Default()
	r.HTMLRender = &TemplRender{}

	postsDB := storage.NewInMemoryDatabase()
	postsService := services.NewPostsService(postsDB)
	postsHandler := handlers.NewPostsHandler(postsService)
	postsHandler.RegisterPosts(&r.RouterGroup)

	r.Static("/htmx", "./public/htmx")
	r.Static("/tw-elements", "./node_modules/tw-elements")
	r.Static("/assets", "./public/assets")
	r.Static("/dist", "./dist")
	r.Run()
}
