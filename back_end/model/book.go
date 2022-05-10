// model 包定义了系统的模型
package model

import (
	"github.com/Astronaut-X-X/GORM/config"
	"github.com/Astronaut-X-X/GORM/global"
	"github.com/Astronaut-X-X/GORM/utils"
	"gorm.io/gorm"
)

// Book 书籍对象的基本信息
type Book struct {
	gorm.Model
	Price       float32 `gorm:"default:0.00"`
	Introduce   string  `gorm:"default:"`
	DownloadUrl string  `gorm:"default:"`
	Image       string  `gorm:"default"`
	BookInfo    BookInfo
}

//
//	InsertBooK 添加书籍信息
// 	参数：
//		book 一个 Book 的引用类型
//	返回值：
//		bool 添加成功返回 true 添加失败返回 false
//
func InsertBook(book *Book) bool {
	result := global.GlobalDB.Create(book)
	utils.PrintResult(result)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

//
//	DeleteBook 根据 book 的 id 删除书籍
//	参数：
//		id 一个 uint 类型
//	返回值：
//		bool 删除成功返回 true 删除失败返回 false
//
func DeleteBook(id uint) bool {
	result := global.GlobalDB.Where("id = ?", id).Delete(&Book{})
	utils.PrintResult(result)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

//
//	UpdateBook 根据 book 的 id 更新书籍信息
//	参数：
//		book 一个 Book 的引用类型
//	返回值：
//		bool 修改成功返回 true 修改失败返回 false
//
func UpdateBook(book *Book) bool {
	result := global.GlobalDB.Save(book)
	utils.PrintResult(result)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

//
//	SelectBookById 根据 id 查询 book 的信息
//	参数：
//		id 一个 uint 类型
//	返回值：
//		book 类型
//
func SelectBookById(id uint) Book {
	var book Book
	result := global.GlobalDB.Where("id = ?", id).First(&book)
	utils.PrintResult(result)
	return book
}

//
//	SelectBooks 根据 pageNumber 查询10条书籍信息
//
func SelectBooks(pageNumber int) []Book {
	var books []Book
	result := global.GlobalDB.Select([]string{"id", "image", "name"}).Offset(pageNumber - 1*config.PAGE_QUANTITY).Limit(config.PAGE_QUANTITY).Find(&books)
	utils.PrintResult(result)
	return books
}
