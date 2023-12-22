package models

import (
	"path"
	"time"
)

func FromArticleEntityToArticleItem(entity *ArticleEntity) *ArticleItem {
	return &ArticleItem{
		ID:       entity.ID,
		Title:    entity.Title,
		Filename: entity.Filename,
		Date:     time.UnixMilli(entity.Timestamp),
	}
}

func FromArticleEntityToArticle(entity *ArticleEntity) *Article {
	return &Article{
		ID:      entity.ID,
		Title:   entity.Title,
		Summary: entity.Summary,
		Image: LocalFile{
			Name: entity.Image,
			Url:  path.Join(ArticlesFolder, entity.Filename, entity.Image),
		},
		Date:     time.UnixMilli(entity.Timestamp),
		Filename: entity.Filename,
	}
}
