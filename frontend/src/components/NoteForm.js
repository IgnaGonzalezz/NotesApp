import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';

function NoteForm({ show, handleClose, handleSubmit, initialNote }) {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [error, setError] = useState('');

  useEffect(() => {
    if (initialNote) {
      setTitle(initialNote.Title);
      setContent(initialNote.Content);
    } else {
      setTitle('');
      setContent('');
    }
    setError(''); // Clear error on modal open/note change
  }, [initialNote, show]);

  const onSubmit = (e) => {
    e.preventDefault();
    if (!title.trim()) {
      setError('El título no puede estar vacío.');
      return;
    }
    setError('');
    handleSubmit({ ...initialNote, Title: title, Content: content });
  };

  return (
    <Modal show={show} onHide={handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>{initialNote ? 'Editar Nota' : 'Crear Nueva Nota'}</Modal.Title>
      </Modal.Header>
      <Form onSubmit={onSubmit}>
        <Modal.Body>
          {error && <Alert variant="danger">{error}</Alert>}
          <Form.Group className="mb-3">
            <Form.Label>Título</Form.Label>
            <Form.Control
              type="text"
              placeholder="Título de la nota"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Label>Contenido</Form.Label>
            <Form.Control
              as="textarea"
              rows={3}
              placeholder="Contenido de la nota"
              value={content}
              onChange={(e) => setContent(e.target.value)}
            />
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Cancelar
          </Button>
          <Button variant="primary" type="submit">
            {initialNote ? 'Guardar Cambios' : 'Crear Nota'}
          </Button>
        </Modal.Footer>
      </Form>
    </Modal>
  );
}

export default NoteForm;
