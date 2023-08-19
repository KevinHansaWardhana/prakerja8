package configs

import (
	"fmt"
	"os"
	"prakerja8/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDatabase(){
	dsn := fmt.Sprintf("root:@tcp(localhost:3306)/coba?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

func initMigration(){
	
	DB.AutoMigrate(&models.Book{}, &models.User{}, &models.CreditCard{})
}