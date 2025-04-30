package routes

import (
    "github.com/gin-gonic/gin"
    "02_notes_app/controllers"
)

func RegisterNoteRoutes(r *gin.Engine) {
    notes := r.Group("/api/notes")
    {
        notes.GET("/", controllers.GetNotes)
        notes.POST("/", controllers.CreateNote)
        notes.PUT("/:id", controllers.UpdateNote)
        notes.DELETE("/:id", controllers.DeleteNote)
    }
}
