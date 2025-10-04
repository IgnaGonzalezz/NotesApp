package controllers

import (
	"net/http"
	"notesapp/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NoteController struct {
	DB *gorm.DB
}

func NewNoteController(db *gorm.DB) *NoteController {
	return &NoteController{DB: db}
}

// Crear una nota
func (c *NoteController) CreateNote(ctx *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := models.Note{
		Title:   input.Title,
		Content: input.Content,
	}

	if err := c.DB.Create(&note).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, note)
}

func (c *NoteController) ListActiveNotes(ctx *gin.Context) {
	var notes []models.Note
	if err := c.DB.Where("archived = ?", false).Find(&notes).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, notes)
}

func (c *NoteController) UpdateNote(ctx *gin.Context) {
	id := ctx.Param("id")
	var note models.Note

	if err := c.DB.First(&note, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Nota no encontrada"})
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	note.Title = input.Title
	note.Content = input.Content

	c.DB.Save(&note)
	ctx.JSON(200, note)

}

func (c *NoteController) DeleteNote(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.DB.Delete(&models.Note{}, id).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Nota eliminada"})
}

func (c *NoteController) ToggleArchiveNote(ctx *gin.Context) {
	id := ctx.Param("id")
	var note models.Note

	if err := c.DB.First(&note, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Nota no encontrada"})
		return
	}

	note.Archived = !note.Archived
	c.DB.Save(&note)

	ctx.JSON(200, note)
}

func (c *NoteController) ListArchivedNotes(ctx *gin.Context) {
	var notes []models.Note
	if err := c.DB.Where("archived = ?", true).Find(&notes).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, notes)
}
