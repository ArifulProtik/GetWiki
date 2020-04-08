package model

import "github.com/jinzhu/gorm"

// Article Holds the Wiki Artcile
type Article struct {
	gorm.Model
	UID   uint   `gorm:"primary_key" json:"id"`
	Title string `gorm:"type:varchar(120)" json:"title"`
	Body  string `gorm:"type:varchar(4000)" json:"body"`
}
