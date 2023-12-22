package models

type ArticleEntity struct {
	ID        int64
	Title     string
	Filename  string
	Image     string
	Summary   string
	Timestamp int64
}
