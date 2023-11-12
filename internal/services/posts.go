package services

import "github.com/jotar910/htmx-templ/internal/models"

type PostsDatabase interface {
	GetPostsList() *models.ArticleList
	GetPostsListFiltered(filters *models.ArticleListFilters) *models.ArticleList
}

type PostsService struct {
	postsDB PostsDatabase
}

func NewPostsService(postsDB PostsDatabase) *PostsService {
	return &PostsService{
		postsDB: postsDB,
	}
}

func (ps *PostsService) GetList(filters *models.ArticleListFilters) *models.ArticleList {
	if filters.Empty() {
		return ps.postsDB.GetPostsList()
	}
	return ps.postsDB.GetPostsListFiltered(filters)
}
