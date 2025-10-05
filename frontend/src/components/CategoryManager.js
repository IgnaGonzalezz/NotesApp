import React, { useState, useEffect } from 'react';
import { getCategories, createCategory, deleteCategory } from '../services/categoryService';
import { ListGroup, Button, Form, Container, Row, Col, Alert } from 'react-bootstrap';

// Este es un componente funcional de React para gestionar categorías.
function CategoryManager({ onCategoriesUpdated }) {
  // --- ESTADOS ---
  // Estado para guardar la lista de categorías que viene de la API.
  const [categories, setCategories] = useState([]);
  // Estado para guardar el nombre de la nueva categoría que el usuario escribe.
  const [newCategoryName, setNewCategoryName] = useState('');
  // Estado para mostrar mensajes de error al usuario.
  const [error, setError] = useState('');

  // --- EFECTOS ---
  // useEffect se ejecuta cuando el componente se monta por primera vez.
  // Es el lugar ideal para hacer llamadas a APIs.
  useEffect(() => {
    fetchCategories();
  }, []);

  // --- FUNCIONES ---
  // Función para obtener las categorías de la API y actualizar el estado.
  const fetchCategories = async () => {
    try {
      const response = await getCategories();
      setCategories(response.data); // Guardamos los datos de la respuesta en el estado.
      setError(''); // Limpiamos cualquier error anterior.
    } catch (err) {
      setError('Error loading categories.');
    }
  };

  // Función que se ejecuta cuando el usuario envía el formulario para crear una categoría.
  const handleCreateCategory = async (e) => {
    e.preventDefault(); // Prevenimos que la página se recargue.
    if (!newCategoryName.trim()) return; // No hacer nada si el nombre está vacío.

    try {
      await createCategory(newCategoryName);
      setNewCategoryName(''); // Limpiamos el input.
      fetchCategories(); // Volvemos a cargar la lista para que se vea la nueva categoría.
      if (onCategoriesUpdated) {
        onCategoriesUpdated(); // Notify parent to refresh global categories
      }
    } catch (err) {
      setError('Error creating category.');
    }
  };

  // Función para eliminar una categoría.
  const handleDeleteCategory = async (id) => {
    try {
      await deleteCategory(id);
      fetchCategories(); // Recargamos la lista para que desaparezca la categoría eliminada.
      setError(''); // Limpiamos cualquier error anterior.
      if (onCategoriesUpdated) {
        onCategoriesUpdated(); // Notify parent to refresh global categories
      }
    } catch (err) {
      console.error("Error deleting category:", err);
      // Intentar extraer un mensaje de error más específico del backend
      const errorMessage = err.response && err.response.data && err.response.data.error 
                           ? err.response.data.error 
                           : 'Error deleting category. Make sure it is not assigned to any note.';
      setError(errorMessage);
    }
  };

  // --- RENDERIZADO ---
  // Esto es lo que el componente dibuja en la pantalla.
  return (
    <Container className="mt-4">
      <Row>
        <Col md={{ span: 6, offset: 3 }}>
          <h2>Category Manager</h2>
          
          {/* Formulario para crear una nueva categoría */}
          <Form onSubmit={handleCreateCategory} className="mb-3">
            <Row>
              <Col>
                <Form.Control
                  type="text"
                  value={newCategoryName}
                  onChange={(e) => setNewCategoryName(e.target.value)}
                  placeholder="New category name"
                />
              </Col>
              <Col xs="auto">
                <Button type="submit">Create</Button>
              </Col>
            </Row>
          </Form>

          {/* Muestra un mensaje de error si existe */}
          {error && <Alert variant="danger">{error}</Alert>}

          {/* Lista de categorías existentes */}
          <ListGroup>
            {categories.map(category => (
              <ListGroup.Item key={category.ID} className="d-flex justify-content-between align-items-center">
                {category.Name}
                <Button variant="danger" size="sm" onClick={() => handleDeleteCategory(category.ID)}>
                  Delete
                </Button>
              </ListGroup.Item>
            ))}
          </ListGroup>
        </Col>
      </Row>
    </Container>
  );
}

export default CategoryManager;
