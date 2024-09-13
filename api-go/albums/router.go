package albums

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// albums slice to seed record album data.
// declaration of array is [x?]T where T is type of array where x? is optional and it represents the length of array
var albums = []album{
	{"1", "Blue Train", "John Coltrane"},
	{"2", "Jeru", "Gerry Mulligan"},
	{"3", "Sarah Vaughan and Clifford Brown", "Sarah Vaughan"},
}

func storeAlbum(c *gin.Context) {
	var newAlbum album
	// c.BindJSON is a method that binds the request body to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// gin.Context is a struct that holds the context of the current request and
// c.indentedJSON is a method that serializes the given struct as JSON
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func pingAlbum(c *gin.Context) {
	response := map[string]string{
		"message": "pong",
	}
	c.JSON(http.StatusOK, response)
}

func RegisterAlbumRouters(router *gin.Engine) {
	router.GET("/albums", getAlbums)
	router.POST("/albums", storeAlbum)
	router.GET("/health", pingAlbum)
}
