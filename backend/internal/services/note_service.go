package services

import (
	"errors"
	"notesapp/internal/repositories"
	"notesapp/pkg/models"
)

type NoteService struct {
	noteRepo     *repositories.NoteRepository
	categoryRepo *repositories.CategoryRepository
}

func NewNoteService(noteRepo *repositories.NoteRepository, categoryRepo *repositories.CategoryRepository) *NoteService {
	return &NoteService{noteRepo: noteRepo, categoryRepo: categoryRepo}
}

func (s *NoteService) CreateNote(title, content string) (*models.Note, error) {
	note := &models.Note{
		Title:    title,
		Content:  content,
		Archived: false,
	}
	err := s.noteRepo.Create(note)
	return note, err
}

func (s *NoteService) ListNotes(archived bool) ([]models.Note, error) {
	return s.noteRepo.FindAll(archived)
}

func (s *NoteService) UpdateNote(id uint, title, content string) (*models.Note, error) {
	note, err := s.noteRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("note not found")
	}

	note.Title = title
	note.Content = content

	if err := s.noteRepo.Update(note); err != nil {
		return nil, err
	}
	return note, nil
}

func (s *NoteService) DeleteNote(id uint) error {
	return s.noteRepo.Delete(id)
}

func (s *NoteService) ToggleArchiveNote(id uint) (*models.Note, error) {
	note, err := s.noteRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("note not found")
	}

	note.Archived = !note.Archived

	if err := s.noteRepo.Update(note); err != nil {
		return nil, err
	}
	return note, nil
}

func (s *NoteService) AddCategoryToNote(noteID, categoryID uint) error {
	note, err := s.noteRepo.FindByID(noteID)
	if err != nil {
		return errors.New("note not found")
	}

	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	return s.noteRepo.AddCategory(note, category)
}

func (s *NoteService) RemoveCategoryFromNote(noteID, categoryID uint) error {
	note, err := s.noteRepo.FindByID(noteID)
	if err != nil {
		return errors.New("note not found")
	}

	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	return s.noteRepo.RemoveCategory(note, category)
}
