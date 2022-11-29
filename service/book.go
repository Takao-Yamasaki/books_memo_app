package model

type Book struct {
	Id      int46  `xorm:"pk autoincr int64" from:"id" json:"id"`
	Title   string `xorm:"varchar(40)" json:"title" from:"title"`
	Content string `xorm:"varchar(40)" json:"content" from:"content"`
}
