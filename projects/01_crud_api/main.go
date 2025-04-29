package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"github.com/gin-contrib/cors"
    "github.com/joho/godotenv"
    "os"
    "log"
    "net/http"
)

type Book struct {
    ID     uint    `json:"id" gorm:"primaryKey"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}

var db *gorm.DB

func initDB() {
    _ = godotenv.Load()

    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
        "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
        os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

	//AutoMigrate is like spring.jpa.hibernate.ddl-auto=update , Automatically creates table if not exists
    db.AutoMigrate(&Book{})
}

func main() {
    initDB()
    r := gin.Default()

	// Enable CORS with default settings
    r.Use(cors.Default())

    r.GET("/books", getBooks)
    r.GET("/books/:id", getBook)
    r.POST("/books", createBook)
    r.PUT("/books/:id", updateBook)
    r.DELETE("/books/:id", deleteBook)

    r.Run(":8080")
}

func getBooks(c *gin.Context) {
    var books []Book
    db.Find(&books)
    c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
    var book Book
    if err := db.First(&book, c.Param("id")).Error; 
	err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
    var book Book
    if err := c.ShouldBindJSON(&book); 
	err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&book)
    c.JSON(http.StatusCreated, book)
}

func updateBook(c *gin.Context) {
    var book Book
    if err := db.First(&book, c.Param("id")).Error; 
	err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    if err := c.ShouldBindJSON(&book); 
	err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Save(&book)
    c.JSON(http.StatusOK, book)
}

func deleteBook(c *gin.Context) {
    var book Book
    if err := db.First(&book, c.Param("id")).Error; 
	err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    db.Delete(&book)
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
