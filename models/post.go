package models

type Post struct {
	Id    string `gorm:"column:id"`
	Title string `gorm:"column:title"`
	Body  string `gorm:"column:body"`
}
