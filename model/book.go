package model

type Book struct {
	Id      int64  `gorm:"pk autoincr int (64)" form:"id" json:"id"`
	Title   string `gorm:"varchar(40)" json:"title" form:"title"`
	Content string `gorm:"varchar(40)" json:"content" from:"content"`
}
