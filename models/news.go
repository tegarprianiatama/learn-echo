package models

type News struct {
	Id      uint   `gorm:"primaryKey" json:"id" form:"id"`
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Content string `gorm:"type:varchar(100);not null" json:"content"`
	Author  string `gorm:"type:varchar(50);not null" json:"author"`
}
