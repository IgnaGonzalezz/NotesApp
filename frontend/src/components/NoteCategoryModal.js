import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';

function NoteCategoryModal({ show, handleClose, note, allCategories, onCategoryChange }) {
  const [selectedCategories, setSelectedCategories] = useState({});
  const [error, setError] = useState('');

  useEffect(() => {
    if (note && allCategories) {
      const initialSelected = {};
      if (note.Categories) { // Add null check here
        note.Categories.forEach(cat => {
          initialSelected[cat.ID] = true;
        });
      }
      setSelectedCategories(initialSelected);
      setError('');
    }
  }, [note, allCategories, show]);

  const handleToggleCategory = (categoryId) => {
    setSelectedCategories(prev => ({
      ...prev,
      [categoryId]: !prev[categoryId],
    }));
  };

  const handleSubmit = () => {
    // Call the parent's onCategoryChange with the updated categories
    onCategoryChange(note.ID, selectedCategories);
    handleClose();
  };

  return (
    <Modal show={show} onHide={handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>Gestionar Categorías para "{note?.Title}"</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        {error && <Alert variant="danger">{error}</Alert>}
        <Form>
          {allCategories.map(category => (
            <Form.Check
              key={category.ID}
              type="checkbox"
              id={`category-${category.ID}`}
              label={category.Name}
              checked={!!selectedCategories[category.ID]}
              onChange={() => handleToggleCategory(category.ID)}
            />
          ))}
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          Cancelar
        </Button>
        <Button variant="primary" onClick={handleSubmit}>
          Guardar Cambios
        </Button>
      </Modal.Footer>
    </Modal>
  );
}

export default NoteCategoryModal;
