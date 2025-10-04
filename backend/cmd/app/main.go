package main

import (
	"fmt"
	"net/http"
	"notesapp/internal/controllers"
	"notesapp/internal/db"
	"notesapp/pkg/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creamos el router de Gin
	r := gin.Default()

	// Conectar a la DB
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	// Migrar tabla Note automáticamente
	err = database.AutoMigrate(&models.Note{}, &models.Category{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migrated!")

	// Crear controladores
	noteController := controllers.NewNoteController(database)
	categoryController := controllers.NewCategoryController(database)

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

	// Archivar nota
	r.PATCH("/notes/:id/archive", noteController.ToggleArchiveNote)

	// Asignar categoria
	r.PUT("/notes/:id/category/:categoryId", noteController.AssignCategory)

	// Desasignar categoria
	r.DELETE("/notes/:id/category", noteController.RemoveCategory)

	//////////////////ENDPOINTS CATEGORIAS///////////////////

	// Crear categoria
	r.POST("/categories", categoryController.CreateCategory)

	// Listar categorias
	r.GET("/categories", categoryController.ListCategories)

	//Listar notas por categoria
	r.GET("/categories/:id/notes", categoryController.ListNotesByCategory)

	// Levantar servidor en el puerto 8080
	r.Run(":8080")
}
