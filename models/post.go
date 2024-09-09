package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id    string
	Title string
	Body  string
}
