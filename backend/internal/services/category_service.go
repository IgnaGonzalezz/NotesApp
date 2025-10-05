package services

import (
	"notesapp/internal/repositories"
	"notesapp/pkg/models"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(name string) (*models.Category, error) {
	category := &models.Category{Name: name}
	err := s.repo.Create(category)
	return category, err
}

func (s *CategoryService) ListCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) ListNotesByCategory(id uint) ([]models.Note, error) {
	return s.repo.FindNotesByCategoryID(id)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	// Primero encontrar la categoria
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err // Categoria no encontrada
	}

	// Eliminar explícitamente todas las asosiaciones en la tabla intermedia
	if err := s.repo.DeleteNoteCategoryAssociations(category.ID); err != nil {
		return err
	}

	// Finalmente borrar la categoria
	return s.repo.Delete(category)
}
