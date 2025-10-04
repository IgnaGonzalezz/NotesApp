package controllers

import (
	"net/http"
	"notesapp/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

// Crear categoría
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Create(&category).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// Listar categorías
func (c *CategoryController) ListCategories(ctx *gin.Context) {
	var categories []models.Category
	if err := c.DB.Find(&categories).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

// Listar notas de una categoría
func (c *CategoryController) ListNotesByCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	var notes []models.Note
	if err := c.DB.Preload("Category").Where("category_id = ? AND archived = ?", categoryID, false).Find(&notes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notes)
}
