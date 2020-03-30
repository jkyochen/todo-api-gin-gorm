package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	Name      string
	Password  string
	Lastlogin *time.Time
}
