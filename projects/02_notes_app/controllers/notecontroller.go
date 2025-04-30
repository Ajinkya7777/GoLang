package controllers

import (
    "02_notes_app/config"
    "02_notes_app/models"
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
    var notes []models.Note
    config.DB.Preload("Category").Find(&notes)
    c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {
    var note models.Note
    if err := c.ShouldBindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fmt.Println("Received Note:", note)  
    if note.CategoryID == 0 {
    note.CategoryID = 1
}

    if err := config.DB.Create(&note).Error; err != nil {
        fmt.Println("DB error:", err)  
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
        return
    }

    c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
    var note models.Note
    id := c.Param("id")

    if err := config.DB.First(&note, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }

    if err := c.ShouldBindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Save(&note)
    c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
    var note models.Note
    id := c.Param("id")

    if err := config.DB.First(&note, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }

    config.DB.Delete(&note)
    c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
