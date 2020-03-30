package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title   string
	Content string
	UserID  *uint
	User    User
}
