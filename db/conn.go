package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open("mysql", "root:123456@/file_store?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
}

func Conn() *gorm.DB {
	return db
}