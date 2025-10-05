# Full Stack Notes Application

This is a full-stack web application that allows users to create, edit, delete, archive, and categorize notes. It consists of a Go backend (REST API) and a React frontend (Single Page Application).

## Features

**Phase 1: Note Management**
- Create, edit, and delete notes.
- Archive/unarchive notes.
- List active notes.
- List archived notes.

**Phase 2: Category Management**
- Add/remove categories to notes.
- Filter notes by category.

## Technologies Used

### Backend
- **Language:** Go (Golang)
- **Web Framework:** Gin Gonic
- **ORM:** GORM
- **Database:** PostgreSQL
- **Dependency Management:** Go Modules

### Frontend
- **Framework:** React.js
- **Language:** JavaScript
- **Styling:** React-Bootstrap
- **HTTP Client:** Axios (for categories), Fetch API (for notes)
- **Dependency Management:** npm

### Other
- **Containerization:** Docker, Docker Compose

## Prerequisites

To run this application, you need to have the following installed:

- **Docker:** [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
- **Docker Compose:** (Usually comes with Docker Desktop installation, or can be installed separately: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/))

## How to Run the Application

1.  **Clone the repository:**
    ```bash
    git clone <repository_url>
    cd <repository_name>
    ```

2.  **Run the application:**
    Execute the provided `run.sh` script from the root directory of the project:
    ```bash
    ./run.sh
    ```

    This command will:
    - Build the Docker images for both the backend and frontend.
    - Start the PostgreSQL database service.
    - Start the backend API service.
    - Start the frontend development server.

3.  **Access the application:**
    - **Frontend:** Open your web browser and navigate to `http://localhost:3000`
    - **Backend API:** The backend API will be running on `http://localhost:8080`

## Project Structure

- `backend/`: Contains the Go backend application.
    - `Dockerfile`: Dockerfile for the backend.
    - `cmd/app/main.go`: Main entry point for the backend application, setting up routes, database connection, and CORS.
    - `internal/controllers/`: Handlers for API requests.
    - `internal/db/`: Database connection and migration logic.
    - `internal/repositories/`: Data access layer.
    - `internal/services/`: Business logic layer.
    - `pkg/models/`: Database models (Note, Category).
- `frontend/`: Contains the React frontend application.
    - `Dockerfile`: Dockerfile for the frontend.
    - `public/`: Public assets.
    - `src/`: Source code for the React application.
        - `components/`: Reusable React components.
        - `services/`: Functions for interacting with the backend API.
- `docker-compose.yml`: Defines the multi-container Docker application.
- `run.sh`: Script to build and run the application using Docker Compose.

## API Endpoints (Backend)

### Notes
- `POST /notes`: Create a new note.
- `GET /notes`: List active notes.
- `GET /notes/archived`: List archived notes.
- `PUT /notes/:id`: Update an existing note.
- `DELETE /notes/:id`: Delete a note.
- `PATCH /notes/:id/archive`: Toggle archive status of a note.
- `PUT /notes/:id/category/:categoryId`: Add a category to a note.
- `DELETE /notes/:id/category/:categoryId`: Remove a category from a note.

### Categories
- `POST /categories`: Create a new category.
- `GET /categories`: List all categories.
- `GET /categories/:id/notes`: List notes by category.
- `DELETE /categories/:id`: Delete a category.

