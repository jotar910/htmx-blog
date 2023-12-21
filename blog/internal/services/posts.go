package services

import "github.com/jotar910/htmx-templ/internal/models"

type PostsDatabase interface {
	GetPostsList() (*models.ArticleList, error)
	GetPostsListFiltered(filters *models.ArticleListFilters) (*models.ArticleList, error)
	GetMostSeenPosts() ([]models.ArticleItem, error)
	GetHighlightPosts() ([]models.ArticleItem, error)
	GetRecentPosts() ([]models.ArticleItem, error)
	GetCarouselPosts() ([]models.ArticleItem, error)
	GetPostById(id int) (*models.Article, error)
	GetRelatedPosts(id int) ([]models.ArticleItem, error)
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
	if filters.Empty() {
		return ps.postsDB.GetPostsList()
	}
	return ps.postsDB.GetPostsListFiltered(filters)
}

func (ps *PostsService) GetMostSeenPosts() ([]models.ArticleItem, error) {
	return ps.postsDB.GetMostSeenPosts()
}

func (ps *PostsService) GetHighlightPosts() ([]models.ArticleItem, error) {
	return ps.postsDB.GetHighlightPosts()
}

func (ps *PostsService) GetRecentPosts() ([]models.ArticleItem, error) {
	return ps.postsDB.GetRecentPosts()
}

func (ps *PostsService) GetCarouselPosts() ([]models.ArticleItem, error) {
	return ps.postsDB.GetCarouselPosts()
}

func (ps *PostsService) GetPostById(id int) (*models.Article, error) {
	return ps.postsDB.GetPostById(id)
}

func (ps *PostsService) GetRelatedPosts(id int) ([]models.ArticleItem, error) {
	return ps.postsDB.GetRelatedPosts(id)
}
