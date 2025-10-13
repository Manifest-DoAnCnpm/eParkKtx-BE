package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("eparkktx.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối database: ", err)
	}

	DB = database
	log.Println("Đã kết nối database thành công!")
}
