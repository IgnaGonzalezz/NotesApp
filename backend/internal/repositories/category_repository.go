package repositories

import (
	"notesapp/pkg/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.Preload("Notes").First(&category, id).Error // Added .Preload("Notes")
	return &category, err
}

func (r *CategoryRepository) FindNotesByCategoryID(id uint) ([]models.Note, error) {
	var notes []models.Note
	err := r.DB.Model(&models.Note{}).
		Joins("JOIN note_categories ON note_categories.note_id = notes.id").
		Where("note_categories.category_id = ?", id).
		Where("notes.archived = ?", false).
		Find(&notes).Error
	return notes, err
}

// ClearAssociations elimina la relacion entre una categoria y sus notas
func (r *CategoryRepository) ClearAssociations(category *models.Category) error {
	return r.DB.Model(category).Association("Notes").Clear()
}

// DeleteNoteCategoryAssociations elimina explícitamente las entradas de la tabla intermedia
func (r *CategoryRepository) DeleteNoteCategoryAssociations(categoryID uint) error {
	return r.DB.Exec("DELETE FROM note_categories WHERE category_id = ?", categoryID).Error
}

func (r *CategoryRepository) Delete(category *models.Category) error {
	return r.DB.Delete(category).Error
}
