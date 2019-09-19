package infrastructure

import (
	"github.com/daleksprinter/clean_todo/server/interfaces/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:password@tcp(0.0.0.0:3306)/memo?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&database.Memo{})

	return db
}
