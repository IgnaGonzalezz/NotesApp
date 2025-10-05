import React from 'react';
import { Row, Col, Card, Button } from 'react-bootstrap';

/**
 * Componente funcional para mostrar una lista de notas.
 * Renderiza las notas en un formato de tarjeta, incluyendo opciones para editar, archivar/desarchivar, eliminar y gestionar categorías.
 * @param {object} props - Propiedades del componente.
 * @param {Array<object>} props.notes - Array de objetos de notas a mostrar.
 * @param {boolean} props.loading - Indica si las notas están cargando.
 * @param {string|null} props.error - Mensaje de error si ocurre uno durante la carga de notas.
 * @param {boolean} props.showArchived - Indica si se están mostrando notas archivadas (no usado directamente en el render, pero útil para contexto).
 * @param {function} props.handleOpenEditNoteModal - Función para abrir el modal de edición de una nota.
 * @param {function} props.handleToggleArchiveNote - Función para alternar el estado de archivado de una nota.
 * @param {function} props.handleDeleteNote - Función para eliminar una nota.
 * @param {function} props.handleOpenCategoryModal - Función para abrir el modal de gestión de categorías de una nota.
 */
function NoteList({
  notes,
  loading,
  error,
  showArchived, // Esta prop no se usa directamente en el renderizado de NoteList, pero se pasa desde App.js
  handleOpenEditNoteModal,
  handleToggleArchiveNote,
  handleDeleteNote,
  handleOpenCategoryModal,
}) {
  // Si las notas están cargando, muestra un mensaje de carga.
  if (loading) {
    return <p>Loading notes...</p>;
  }

  // Si hay un error, muestra el mensaje de error.
  if (error) {
    return <p className="text-danger">Error: {error}</p>;
  }

  // Si no hay notas para mostrar, indica que no hay notas.
  if (notes.length === 0) {
    return <p>No notes to display.</p>;
  }

  // Renderiza las notas en un diseño de cuadrícula utilizando componentes de React-Bootstrap.
  return (
    <Row xs={1} md={2} lg={3} className="g-4">
      {notes.map((note) => (
        <Col key={note.ID}> {/* Cada nota se renderiza en una columna */}
          <Card>
            <Card.Body>
              <Card.Title>{note.Title}</Card.Title> {/* Título de la nota */}
              <Card.Text>{note.Content}</Card.Text> {/* Contenido de la nota */}
              {/* Muestra las categorías asociadas a la nota, si existen */}
              {note.Categories && note.Categories.length > 0 && (
                <div>
                  {note.Categories.map((cat) => (
                    <span key={cat.ID} className="badge bg-info text-dark me-1">
                      {cat.Name}
                    </span>
                  ))}
                </div>
              )}
              {/* Botones de acción para cada nota */}
              <div className="mt-3">
                <Button variant="info" size="sm" className="me-2" onClick={() => handleOpenEditNoteModal(note)}>
                  Edit
                </Button>
                <Button
                  variant="warning"
                  size="sm"
                  className="me-2"
                  onClick={() => handleToggleArchiveNote(note.ID)}
                >
                  {note.Archived ? 'Unarchive' : 'Archive'} {/* Texto dinámico según el estado de archivado */}
                </Button>
                <Button variant="danger" size="sm" className="me-2" onClick={() => handleDeleteNote(note.ID)}>
                  Delete
                </Button>
                <Button
                  variant="success"
                  size="sm"
                  className="me-2"
                  onClick={() => handleOpenCategoryModal(note)}
                >
                  Manage Categories
                </Button>
              </div>
            </Card.Body>
          </Card>
        </Col>
      ))}
    </Row>
  );
}

export default NoteList;
