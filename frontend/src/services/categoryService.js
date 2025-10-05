import axios from 'axios';

const API_URL = 'http://localhost:8080';

/**
 * Obtiene todas las categorías desde el backend.
 * @returns {Promise<Array>} Una promesa que resuelve a un array de categorías.
 */
export const getCategories = () => {
  return axios.get(`${API_URL}/categories`);
};

/**
 * Elimina una categoría por su ID.
 * @param {number} id El ID de la categoría a eliminar.
 * @returns {Promise} Una promesa que se resuelve cuando la categoría es eliminada.
 */
export const deleteCategory = (id) => {
  return axios.delete(`${API_URL}/categories/${id}`);
};

/**
 * Crea una nueva categoría.
 * @param {string} name El nombre de la nueva categoría.
 * @returns {Promise<Object>} Una promesa que resuelve al objeto de la nueva categoría creada.
 */
export const createCategory = (name) => {
  return axios.post(`${API_URL}/categories`, { name });
};
