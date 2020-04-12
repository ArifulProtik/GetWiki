package db

import "getwiki/model"

// Storage Holds Everything of DB Package
type Storage interface {
	Migrator()
	// Starts Article
	CreateArticle(*model.Article) (err error)
	SearchArticle(string) ([]model.Article, error)
	Getbyid(uint64) (model.Article, error)
	Getall() ([]model.Article, error)
}
