package controllers

import (
	"net/http"
	"notesapp/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	service *services.NoteService
}

func NewNoteController(service *services.NoteService) *NoteController {
	return &NoteController{service: service}
}

type NoteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

// Crear una nota
func (c *NoteController) CreateNote(ctx *gin.Context) {
	var req NoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := c.service.CreateNote(req.Title, req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	ctx.JSON(http.StatusCreated, note)
}

// Listar notas activas
func (c *NoteController) ListActiveNotes(ctx *gin.Context) {
	notes, err := c.service.ListNotes(false)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notes)
}

// Listar notas archivadas
func (c *NoteController) ListArchivedNotes(ctx *gin.Context) {
	notes, err := c.service.ListNotes(true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notes)
}

// Actualizar nota
func (c *NoteController) UpdateNote(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var req NoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := c.service.UpdateNote(uint(id), req.Title, req.Content)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, note)
}

// Borrar nota
func (c *NoteController) DeleteNote(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	if err := c.service.DeleteNote(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

// Archivar/Desarchivar nota
func (c *NoteController) ToggleArchiveNote(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	note, err := c.service.ToggleArchiveNote(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, note)
}

// Asignar categoria a una nota
func (c *NoteController) AddCategoryToNote(ctx *gin.Context) {
	noteID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	categoryID, err := strconv.ParseUint(ctx.Param("categoryId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := c.service.AddCategoryToNote(uint(noteID), uint(categoryID)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// Desasignar categoria de una nota
func (c *NoteController) RemoveCategoryFromNote(ctx *gin.Context) {
	noteID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	categoryID, err := strconv.ParseUint(ctx.Param("categoryId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := c.service.RemoveCategoryFromNote(uint(noteID), uint(categoryID)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
