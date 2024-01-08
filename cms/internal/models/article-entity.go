package models

type ArticleEntity struct {
	ID        int64
	Title     string
	Filename  string
	Image     string
	Summary   string
	Timestamp int64
	// Joins
	CarouselID *int64 `db:"carousel_id"`
}
