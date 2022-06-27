package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: 1, Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 10.99},
	{ID: 2, Title: "The Wall", Artist: "Pink Floyd", Price: 10.99},
	{ID: 3, Title: "The Division Bell", Artist: "Pink Floyd", Price: 10.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)
	router.Run("localhost:8080")
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum Album
	if len(albums) <= 0 {
		newAlbum.ID = 1
	}

	newAlbum.ID = albums[len(albums)-1].ID + 1

	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusOK, newAlbum)
}

func getAlbumByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbum(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedAlbum Album
	if err := ctx.BindJSON(&updatedAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, album := range albums {
		if album.ID == id {
			updatedAlbum.ID = album.ID
			albums[index] = updatedAlbum
			ctx.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
}

func deleteAlbum(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, album := range albums {
		if album.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
