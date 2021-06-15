package database

import (
	"github.com/gleo08/fresherOCG/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DbURL = "root:12345678@tcp(127.0.0.1:3306)/fiber_v1?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open(DbURL), &gorm.Config{})

	if err != nil {
		panic("Can not connect to database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Product{})
}
