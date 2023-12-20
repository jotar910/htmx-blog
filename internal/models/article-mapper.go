package models

import (
	"fmt"
	"path"
	"time"
)

func FromArticleEntityToArticleItem(entity *ArticleEntity) *ArticleItem {
	return &ArticleItem{
		ID:    entity.ID,
		Title: entity.Title,
		Image: LocalFile{
			Name: entity.Image,
			Url:  fmt.Sprintf(path.Join(ArticlesFolder, entity.Filename, entity.Image)),
		},
		Summary: entity.Summary,
		Date:    time.UnixMilli(entity.Timestamp),
	}
}

func FromArticleEntityToArticle(entity *ArticleEntity) *Article {
	return &Article{
		ID:    entity.ID,
		Title: entity.Title,
		Image: LocalFile{
			Name: entity.Image,
			Url:  fmt.Sprintf(path.Join(ArticlesFolder, entity.Filename, entity.Image)),
		},
		Date:     time.UnixMilli(entity.Timestamp),
		Filename: entity.Filename,
	}
}
