package controllers

import (
	"net/http"
	"notesapp/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

// Crear categoría
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := c.service.CreateCategory(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

// Listar categorías
func (c *CategoryController) ListCategories(ctx *gin.Context) {
	categories, err := c.service.ListCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list categories"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// Listar notas de una categoría
func (c *CategoryController) ListNotesByCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	notes, err := c.service.ListNotesByCategory(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found or failed to fetch notes"})
		return
	}

	ctx.JSON(http.StatusOK, notes)
}

// Eliminar categoría
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := c.service.DeleteCategory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
