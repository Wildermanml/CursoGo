package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Tittle string  `json:"tittle"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Datos simulando una bd
var albums = []album{
	{ID: "1", Tittle: "Blue Train", Artist: "John Coltrane", Price: 56.9},
	{ID: "2", Tittle: "Jeru", Artist: "John Coltrane", Price: 20.99},
	{ID: "3", Tittle: "Sarah Vaughan and clifford", Artist: "Sarah Vaughan", Price: 47.99},
}

func main() {
	fmt.Println("Bienvenido a vinyl-api")
	// Se crea una instancia de gin, Default() devuelve un router con configuración por defecto
	router := gin.Default()
	// Se define una ruta GET, el primer parámetro es la ruta y el segundo es un handler, el handler es una función que recibe un contexto como parámetro y devuelve un JSON
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Bienvenido a vinyl-api"}) })

	//Consultar albums
	router.GET("/albums", getAlbums)

	//Consultar album por id
	router.GET("/album/:id", getAlbumById)

	//Guardar Album
	router.POST("/postAlbum", postAlbum)

	//Actualizar
	router.PUT("album/:id", updateAlbum)

	//Eliminar
	router.DELETE("album/:id", deleteAlbum)
	// Se inicia el servidor en el puerto 8080 en el localhost
	router.Run(":8080")
}

// funciones controller

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id: ", id)
	for _, album := range albums {
		if id == album.ID {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	// Bind the JSON data from the request to updatedAlbum
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	for i, album := range albums {
		if id == album.ID {
			albums[i].Tittle = updatedAlbum.Tittle
			albums[i].Artist = updatedAlbum.Artist
			albums[i].Price = updatedAlbum.Price
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if id == album.ID {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "El album: " + id + " no fue encontrado"})
}
