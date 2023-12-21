package main

import (
	_ "github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jotar910/htmx-templ/internal/handlers"
	"github.com/jotar910/htmx-templ/internal/services"
	"github.com/jotar910/htmx-templ/internal/storage"
	"github.com/jotar910/htmx-templ/pkg/logger"
)

func main() {
	logger.Init()

	r := gin.Default()
	r.HTMLRender = &TemplRender{}

	logger.L.Info("connecting to db...")
	db, err := sqlx.Open("sqlite3", "./articles.db")
	if err != nil {
		logger.L.Fatal(err.Error())
	}
	defer db.Close()

	logger.L.Info("setting up server...")
	postsDB := storage.NewSQLiteDatabase(db)
	postsService := services.NewPostsService(postsDB)
	postsHandler := handlers.NewPostsHandler(postsService)
	postsHandler.RegisterPosts(&r.RouterGroup)

	r.Static("/htmx", "./public/htmx")
	r.Static("/assets", "./public/assets")
	r.Static("/articles", "./public/articles")
	r.Static("/dist", "./dist")

	logger.L.Info("running server...")
	r.Run()
}
