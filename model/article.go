package model

import "time"

// Article Holds the Wiki Artcile
type Article struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;unique;not null" json:"title"`
	Body      string    `sql:"type:text" gorm:"not null" json:"body"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
