// model 包定义了系统的模型
package model

import "time"

// BookInfo 书籍的详细信息
type BookInfo struct {
	Name string `gorm:"default:"`
	PublicationTime time.Time
	BookNumber string
	Price float32
	Pages uint16
	OriginalName string
	OriginalNumber string
}

/*
书名：程序员的英语
出版日期 2018-03-19
书号 978-7-115-47305-9
定价 49
页数 274
原书名 English for Developers
原书号 9788968482199
*/