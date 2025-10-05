import React, { useState, useEffect } from 'react';
import CategoryManager from './components/CategoryManager';
import NoteForm from './components/NoteForm';
import NoteCategoryModal from './components/NoteCategoryModal';
import { Container, Navbar, Nav, Button, Dropdown, Row, Col, Card } from 'react-bootstrap';
import * as noteService from './services/noteService';
import * as categoryService from './services/categoryService';

function App() {
  const [notes, setNotes] = useState([]);
  const [categories, setCategories] = useState([]);
  const [showArchived, setShowArchived] = useState(false);
  const [selectedCategory, setSelectedCategory] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showNoteModal, setShowNoteModal] = useState(false);
  const [currentNote, setCurrentNote] = useState(null); // Nota que se está editando o creando
  const [showCategoryModal, setShowCategoryModal] = useState(false);
  const [currentNoteForCategories, setCurrentNoteForCategories] = useState(null); // Nota para gestionar categorías

  const handleCloseNoteModal = () => {
    setShowNoteModal(false);
    setCurrentNote(null);
  };

  const handleCloseCategoryModal = () => {
    setShowCategoryModal(false);
    setCurrentNoteForCategories(null);
  };

  const fetchNotes = async () => {
    setLoading(true);
    setError(null);
    try {
      const fetchedNotes = await noteService.getNotes(showArchived, selectedCategory);
      setNotes(fetchedNotes);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const fetchCategories = async () => {
    try {
      const response = await categoryService.getCategories();
      setCategories(response.data);
    } catch (err) {
      console.error("Error fetching categories:", err);
    }
  };

  // Fetch notes
  useEffect(() => {
    fetchNotes();
  }, [showArchived, selectedCategory]);

  // Fetch categories
  useEffect(() => {
    fetchCategories();
  }, []);

  const handleCategoryFilter = (categoryId) => {
    setSelectedCategory(categoryId === selectedCategory ? null : categoryId); // Toggle filter
  };

  const handleOpenCreateNoteModal = () => {
    setCurrentNote(null);
    setShowNoteModal(true);
  };

  const handleOpenEditNoteModal = (note) => {
    setCurrentNote(note);
    setShowNoteModal(true);
  };

  const handleSaveNote = async (noteData) => {
    try {
      if (noteData.ID) {
        await noteService.updateNote(noteData.ID, { Title: noteData.Title, Content: noteData.Content });
      } else {
        await noteService.createNote({ Title: noteData.Title, Content: noteData.Content });
      }
      handleCloseNoteModal();
      fetchNotes(); // Refresh notes after save
    } catch (err) {
      console.error("Error saving note:", err);
      setError("Error al guardar la nota.");
    }
  };

  const handleToggleArchiveNote = async (id) => {
    try {
      await noteService.toggleArchiveNote(id);
      fetchNotes(); // Refresh notes after archiving
    } catch (err) {
      console.error("Error toggling archive status:", err);
      setError("Error al cambiar el estado de archivado.");
    }
  };

  const handleDeleteNote = async (id) => {
    console.log("Attempting to delete note with ID:", id);
    if (window.confirm("¿Estás seguro de que quieres eliminar esta nota?")) {
      try {
        await noteService.deleteNote(id);
        fetchNotes(); // Refresh notes after deletion
      } catch (err) {
        console.error("Error deleting note:", err);
        setError("Error al eliminar la nota.");
      }
    }
  };

  const handleOpenCategoryModal = (note) => {
    console.log("Note object passed to category modal:", note);
    setCurrentNoteForCategories(note);
    setShowCategoryModal(true);
  };

  const handleCategoryChangeForNote = async (noteId, newSelectedCategories) => {
    try {
      const noteToUpdate = notes.find(n => n.ID === noteId);
      if (!noteToUpdate) return;

      const currentCategoryIds = new Set(noteToUpdate.Categories.map(cat => cat.ID));
      const newCategoryIds = new Set(Object.keys(newSelectedCategories).filter(key => newSelectedCategories[key]).map(Number));

      console.log("Note ID:", noteId);
      console.log("Current Category IDs:", Array.from(currentCategoryIds));
      console.log("New Selected Category IDs:", Array.from(newCategoryIds));

      // Categories to add
      for (const categoryId of newCategoryIds) {
        if (!currentCategoryIds.has(categoryId)) {
          console.log("Adding category", categoryId, "to note", noteId);
          await noteService.addCategoryToNote(noteId, categoryId);
        }
      }

      // Categories to remove
      for (const categoryId of currentCategoryIds) {
        if (!newCategoryIds.has(categoryId)) {
          console.log("Removing category", categoryId, "from note", noteId);
          await noteService.removeCategoryFromNote(noteId, categoryId);
        }
      }
      fetchNotes(); // Refresh notes after category changes
    } catch (err) {
      console.error("Error updating note categories:", err); // Log the full error object
      setError("Error al actualizar las categorías de la nota.");
    }
  };

  return (
    <div>
      <Navbar bg="dark" variant="dark" expand="lg">
        <Container>
          <Navbar.Brand href="#home">Note Manager - Ensolvers</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link onClick={() => setShowArchived(false)} active={!showArchived}>Active Notes</Nav.Link>
              <Nav.Link onClick={() => setShowArchived(true)} active={showArchived}>Archived Notes</Nav.Link>
            </Nav>
            <Dropdown className="ms-auto">
              <Dropdown.Toggle variant="secondary" id="dropdown-basic">
                {selectedCategory ? categories.find(cat => cat.ID === selectedCategory)?.Name : 'Filter by Category'}
              </Dropdown.Toggle>
              <Dropdown.Menu>
                <Dropdown.Item onClick={() => handleCategoryFilter(null)}>All Categories</Dropdown.Item>
                {categories.map(category => (
                  <Dropdown.Item key={category.ID} onClick={() => handleCategoryFilter(category.ID)}>
                    {category.Name}
                  </Dropdown.Item>
                ))}
              </Dropdown.Menu>
            </Dropdown>
            <Button variant="primary" className="ms-3" onClick={handleOpenCreateNoteModal}>Create New Note</Button>
          </Navbar.Collapse>
        </Container>
      </Navbar>

      <Container className="mt-4">
        <h2 className="mb-4">{showArchived ? 'Archived Notes' : 'Active Notes'}</h2>
        {loading && <p>Loading notes...</p>}
        {error && <p className="text-danger">Error: {error}</p>}
        {!loading && !error && notes.length === 0 && <p>No notes to display.</p>}
        
        <Row xs={1} md={2} lg={3} className="g-4">
          {!loading && !error && notes.map(note => (
            <Col key={note.ID}>
              <Card>
                <Card.Body>
                  <Card.Title>{note.Title}</Card.Title>
                  <Card.Text>{note.Content}</Card.Text>
                  {note.Categories && note.Categories.length > 0 && (
                    <div>
                      {note.Categories.map(cat => (
                        <span key={cat.ID} className="badge bg-info text-dark me-1">{cat.Name}</span>
                      ))}
                    </div>
                  )}
                  <div className="mt-3">
                    <Button variant="info" size="sm" className="me-2" onClick={() => handleOpenEditNoteModal(note)}>Edit</Button>
                    <Button variant="warning" size="sm" className="me-2" onClick={() => handleToggleArchiveNote(note.ID)}>{note.Archived ? 'Unarchive' : 'Archive'}</Button>
                    <Button variant="danger" size="sm" className="me-2" onClick={() => handleDeleteNote(note.ID)}>Delete</Button>
                    <Button variant="success" size="sm" className="me-2" onClick={() => handleOpenCategoryModal(note)}>Manage Categories</Button>
                  </div>
                </Card.Body>
              </Card>
            </Col>
          ))}
        </Row>

        {/* For now, we keep CategoryManager here. We will integrate it better later. */}
        <CategoryManager onCategoriesUpdated={fetchCategories} />
      </Container>

      <NoteForm
        show={showNoteModal}
        handleClose={handleCloseNoteModal}
        handleSubmit={handleSaveNote}
        initialNote={currentNote}
      />

      <NoteCategoryModal
        show={showCategoryModal}
        handleClose={handleCloseCategoryModal}
        note={currentNoteForCategories}
        allCategories={categories}
        onCategoryChange={handleCategoryChangeForNote}
      />
    </div>
  );
}

export default App;




