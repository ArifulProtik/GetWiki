package db

import "getwiki/model"

// Storage Holds Everything of DB Package
type Storage interface {
	Pingdb()
	// Starts Article
	CreateArticle(*model.Article) (err error)
}
