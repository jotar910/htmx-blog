package services

import "github.com/jotar910/buzzer-cms/internal/models"

type PostsDatabase interface {
	GetPostsList() (*models.ArticleList, error)
	GetPostsListFiltered(filters *models.ArticleListFilters) (*models.ArticleList, error)
	GetPostById(id int) (*models.Article, error)
}

type PostsService struct {
	postsDB PostsDatabase
}

func NewPostsService(postsDB PostsDatabase) *PostsService {
	return &PostsService{
		postsDB: postsDB,
	}
}

func (ps *PostsService) GetList(filters *models.ArticleListFilters) (*models.ArticleList, error) {
	if filters == nil || filters.Empty() {
		return ps.postsDB.GetPostsList()
	}
	return ps.postsDB.GetPostsListFiltered(filters)
}

func (ps *PostsService) GetPostById(id int) (*models.Article, error) {
	return ps.postsDB.GetPostById(id)
}
