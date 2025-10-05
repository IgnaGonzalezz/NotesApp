package repositories

import (
	"notesapp/pkg/models"

	"gorm.io/gorm"
)

type NoteRepository struct {
	DB *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{DB: db}
}

func (r *NoteRepository) Create(note *models.Note) error {
	return r.DB.Create(note).Error
}

func (r *NoteRepository) FindAll(archived bool) ([]models.Note, error) {
	var notes []models.Note
	err := r.DB.Preload("Categories").Where("archived = ?", archived).Find(&notes).Error
	return notes, err
}

func (r *NoteRepository) FindByID(id uint) (*models.Note, error) {
	var note models.Note
	err := r.DB.Preload("Categories").First(&note, id).Error
	return &note, err
}

func (r *NoteRepository) Update(note *models.Note) error {
	return r.DB.Save(note).Error
}

func (r *NoteRepository) Delete(id uint) error {
	// Primero encontrar la nota para eliminar las relaciones
	var note models.Note
	if err := r.DB.First(&note, id).Error; err != nil {
		return err // No se encontro la nota u otro error
	}

	// Limpiar relaciones many-to-many con categorias
	if err := r.DB.Model(&note).Association("Categories").Clear(); err != nil {
		return err // Error limpiando relaciones
	}

	// Ahora borrar la nota
	return r.DB.Delete(&models.Note{}, id).Error
}

func (r *NoteRepository) AddCategory(note *models.Note, category *models.Category) error {
	return r.DB.Model(note).Association("Categories").Append(category)
}

func (r *NoteRepository) RemoveCategory(note *models.Note, category *models.Category) error {
	return r.DB.Model(note).Association("Categories").Delete(category)
}
