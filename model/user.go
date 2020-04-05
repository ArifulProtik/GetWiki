package model

import "github.com/jinzhu/gorm"

// User Holds The User Model
type User struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Username string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
	Name     string
	Status   bool
}
