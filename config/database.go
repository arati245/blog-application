package config

import (
	"blog_api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Replace with your MySQL details
	dsn := "root:@Arati2004@tcp(127.0.0.1:3306)/blogdb?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to MySQL: ", err)
	}

	// Auto-migrate tables
	database.AutoMigrate(&models.Post{})
	DB = database
}
