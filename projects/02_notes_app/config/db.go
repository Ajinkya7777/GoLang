package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
    "github.com/joho/godotenv"
	"02_notes_app/models"
)

var DB *gorm.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
        "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
        os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

    var err2 error
    DB, err2 = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err2 != nil {
        log.Fatal("Failed to connect to database: ", err2)
    }

    DB.AutoMigrate(&models.Note{}) 
}
