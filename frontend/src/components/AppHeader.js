import React from 'react';
import { Navbar, Container, Nav, Dropdown, Button } from 'react-bootstrap';

/**
 * Componente para la cabecera de la aplicación, incluyendo la barra de navegación y los controles de filtro.
 * Recibe los estados y funciones de manejo de filtros y creación de notas como props.
 */
function AppHeader({
  showArchived,
  setShowArchived,
  selectedCategory,
  categories,
  handleCategoryFilter,
  handleOpenCreateNoteModal,
}) {
  return (
    <Navbar bg="dark" variant="dark" expand="lg">
      <Container>
        <Navbar.Brand href="#home">Note Manager - Ensolvers</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link onClick={() => setShowArchived(false)} active={!showArchived}>
              Active Notes
            </Nav.Link>
            <Nav.Link onClick={() => setShowArchived(true)} active={showArchived}>
              Archived Notes
            </Nav.Link>
          </Nav>
          <Dropdown className="ms-auto">
            <Dropdown.Toggle variant="secondary" id="dropdown-basic">
              {selectedCategory
                ? categories.find((cat) => cat.ID === selectedCategory)?.Name
                : 'Filter by Category'}
            </Dropdown.Toggle>
            <Dropdown.Menu>
              <Dropdown.Item onClick={() => handleCategoryFilter(null)}>
                All Categories
              </Dropdown.Item>
              {categories.map((category) => (
                <Dropdown.Item
                  key={category.ID}
                  onClick={() => handleCategoryFilter(category.ID)}
                >
                  {category.Name}
                </Dropdown.Item>
              ))}
            </Dropdown.Menu>
          </Dropdown>
          <Button
            variant="primary"
            className="ms-3"
            onClick={handleOpenCreateNoteModal}
          >
            Create New Note
          </Button>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default AppHeader;
