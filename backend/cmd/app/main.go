package main

import (
	"fmt"
	"net/http"
	"notesapp/internal/controllers"
	"notesapp/internal/db"
	"notesapp/internal/repositories"
	"notesapp/internal/services"
	"notesapp/pkg/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creacion del router de Gin
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// Conectar a la DB
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	// Migrar tablas automáticamente
	err = database.AutoMigrate(&models.Note{}, &models.Category{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migrated!")

	// --- Inicialización de Repositorios ---
	categoryRepo := repositories.NewCategoryRepository(database)
	noteRepo := repositories.NewNoteRepository(database)

	// --- Inicialización de Servicios ---
	categoryService := services.NewCategoryService(categoryRepo)
	noteService := services.NewNoteService(noteRepo, categoryRepo)

	// --- Inicialización de Controladores ---
	categoryController := controllers.NewCategoryController(categoryService)
	noteController := controllers.NewNoteController(noteService)

	// Endpoint de prueba /health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	//////////////////ENDPOINTS NOTAS///////////////////

	// Crear nota
	r.POST("/notes", noteController.CreateNote)

	// Listar notas activas
	r.GET("/notes", noteController.ListActiveNotes)

	// Listar notas archivadas
	r.GET("/notes/archived", noteController.ListArchivedNotes)

	// Actualizar nota existente
	r.PUT("/notes/:id", noteController.UpdateNote)

	// Borrar nota
	r.DELETE("/notes/:id", noteController.DeleteNote)

	// Archivar/desarchivar nota
	r.PATCH("/notes/:id/archive", noteController.ToggleArchiveNote)

	// Asignar categoria
	r.PUT("/notes/:id/category/:categoryId", noteController.AddCategoryToNote)

	// Desasignar categoria
	r.DELETE("/notes/:id/category/:categoryId", noteController.RemoveCategoryFromNote)

	//////////////////ENDPOINTS CATEGORIAS///////////////////

	// Crear categoria
	r.POST("/categories", categoryController.CreateCategory)

	// Listar categorias
	r.GET("/categories", categoryController.ListCategories)

	// Listar notas por categoria
	r.GET("/categories/:id/notes", categoryController.ListNotesByCategory)

	// Eliminar categoria
	r.DELETE("/categories/:id", categoryController.DeleteCategory)

	// Levantar servidor en el puerto 8080
	r.Run(":8080")
}
