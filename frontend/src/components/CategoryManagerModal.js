import React from 'react';
import { Modal, Button } from 'react-bootstrap';
import CategoryManager from './CategoryManager';

/**
 * Componente Modal para gestionar categorías.
 * Envuelve el componente CategoryManager en un modal de Bootstrap.
 * @param {object} props - Propiedades del componente.
 * @param {boolean} props.show - Controla la visibilidad del modal.
 * @param {function} props.handleClose - Función para cerrar el modal.
 * @param {function} props.onCategoriesUpdated - Función de callback cuando las categorías son actualizadas.
 */
function CategoryManagerModal({ show, handleClose, onCategoriesUpdated }) {
  return (
    <Modal show={show} onHide={handleClose} size="lg">
      {/* Encabezado del modal */}
      <Modal.Header closeButton>
        <Modal.Title>Manage Categories</Modal.Title>
      </Modal.Header>
      {/* Cuerpo del modal que contiene el CategoryManager */}
      <Modal.Body>
        <CategoryManager onCategoriesUpdated={onCategoriesUpdated} />
      </Modal.Body>
      {/* Pie de página del modal con un botón para cerrar */}
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          Close
        </Button>
      </Modal.Footer>
    </Modal>
  );
}

export default CategoryManagerModal;
