package models

import (
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
