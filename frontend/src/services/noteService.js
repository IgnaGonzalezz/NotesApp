const API_BASE_URL = 'http://localhost:8080/notes'; 

const handleResponse = async (response) => {
  console.log("Response status:", response.status);
  console.log("Response OK:", response.ok);
  const contentType = response.headers.get("content-type");

  if (!response.ok) {
    let errorData = { error: 'Something went wrong' };
    if (contentType && contentType.includes("application/json")) {
      errorData = await response.json();
    }
    console.error("API Error Data:", errorData);
    throw new Error(errorData.error || 'Something went wrong');
  }


  if (response.status === 204 || !contentType || !contentType.includes("application/json")) {
    return null; 
  }

  return response.json();
};

export const getNotes = async (archived = false, categoryId = null) => {
  let url = API_BASE_URL;
  if (archived) {
    url = `${API_BASE_URL}/archived`;
  }

  if (categoryId) {
    url = `http://localhost:8080/categories/${categoryId}/notes`;
  }
  console.log("Fetching notes from URL:", url);
  const response = await fetch(url);
  return handleResponse(response);
};

export const createNote = async (note) => {
  const response = await fetch(API_BASE_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(note),
  });
  return handleResponse(response);
};

export const updateNote = async (id, note) => {
  const response = await fetch(`${API_BASE_URL}/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(note),
  });
  return handleResponse(response);
};

export const deleteNote = async (id) => {
  console.log("noteService: Deleting note with ID:", id);
  const response = await fetch(`${API_BASE_URL}/${id}`, {
    method: 'DELETE',
  });
  console.log("noteService: Delete response:", response);
  return handleResponse(response);
};

export const toggleArchiveNote = async (id) => {
 
  const response = await fetch(`${API_BASE_URL}/${id}/archive`, {
    method: 'PATCH',
  });
  return handleResponse(response);
};

export const addCategoryToNote = async (noteId, categoryId) => {
  
  const response = await fetch(`${API_BASE_URL}/${noteId}/category/${categoryId}`, {
    method: 'PUT',
  });
  return handleResponse(response);
};

export const removeCategoryFromNote = async (noteId, categoryId) => {
  
  const response = await fetch(`${API_BASE_URL}/${noteId}/category/${categoryId}`, {
    method: 'DELETE',
  });
  return handleResponse(response);
};
