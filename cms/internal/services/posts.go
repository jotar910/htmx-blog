package services

import (
	"github.com/jotar910/buzzer-cms/internal/models"
	"github.com/jotar910/buzzer-cms/pkg/logger"
	"github.com/pkg/errors"
)

type PostsDatabase interface {
	GetPostsList() (*models.ArticleList, error)
	GetPostsListFiltered(filters *models.ArticleListFilters) (*models.ArticleList, error)
	GetPostById(id int) (*models.Article, error)
	AddCarouselArticle(id int64) error
	RemoveCarouselArticle(id int64) error
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

func (ps *PostsService) PatchPost(id int, patch *models.ArticlePatch) (*models.Article, error) {
	article, err := ps.postsDB.GetPostById(id)
	if err != nil {
		return nil, err
	}

	if err := ps.updatePostCarouselState(article, patch); err != nil {
		logger.L.Errorf("updating carousel state: %v", err)
	}

	return article, nil
}

func (ps *PostsService) updatePostCarouselState(article *models.Article, patch *models.ArticlePatch) error {
	id := article.ID

	if patch.Carousel == nil {
		logger.L.Debugf("skipping update for article %d on carousel", id)
		return nil
	}

	article.Carousel = *patch.Carousel
	logger.L.Debugf("updating article %d on carousel as %v", id, article.Carousel)
	if article.Carousel {
		if err := ps.postsDB.AddCarouselArticle(id); err != nil {
			return errors.Wrap(err, "adding article to carousel")
		}
	} else {
		if err := ps.postsDB.RemoveCarouselArticle(id); err != nil {
			return errors.Wrap(err, "removing article from carousel")
		}
	}

	return nil
}
