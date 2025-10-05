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
	// First, find the category
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err // Category not found
	}

	// Remove all associations from the join table
	if err := s.repo.ClearAssociations(category); err != nil {
		return err
	}

	// Finally, delete the category itself
	return s.repo.Delete(category)
}
