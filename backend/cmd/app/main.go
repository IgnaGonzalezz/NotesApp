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
	err = database.AutoMigrate(&models.Note{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migrated!")

	// Crear controlador
	noteController := controllers.NewNoteController(database)

	// Endpoint de prueba /health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Endpoint para crear nota
	r.POST("/notes", noteController.CreateNote)

	// Endpoint para listar notas activas
	r.GET("/notes", noteController.ListActiveNotes)

	//Endpoint para listar notas archivadas
	r.GET("/notes/archived", noteController.ListArchivedNotes)

	//Endpoint para actualizar nota existente
	r.PUT("/notes/:id", noteController.UpdateNote)

	//Endpoint para borrar nota
	r.DELETE("/notes/:id", noteController.DeleteNote)

	//Endpoint para archivar nota
	r.PATCH("/notes/:id/archive", noteController.ToggleArchiveNote)

	// Levantamos el servidor en el puerto 8080
	r.Run(":8080")
}
