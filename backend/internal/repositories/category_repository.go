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
	err := r.DB.First(&category, id).Error
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

// ClearAssociations removes the link between a category and all its notes.
func (r *CategoryRepository) ClearAssociations(category *models.Category) error {
	return r.DB.Model(category).Association("Notes").Clear()
}

func (r *CategoryRepository) Delete(category *models.Category) error {
	return r.DB.Delete(category).Error
}
