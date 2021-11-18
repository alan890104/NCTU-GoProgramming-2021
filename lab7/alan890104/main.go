package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	// write your own struct
	ID    string `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	{
		ID:    "1",
		Name:  "Blue Bird",
		Pages: "500",
	},
}

func remove(slice []Book, s int) []Book {
	return append(slice[:s], slice[s+1:]...)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(200, bookshelf)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	for i := range bookshelf {
		if bookshelf[i].ID == id {
			c.IndentedJSON(200, bookshelf[i])
			return
		}
	}
	c.IndentedJSON(200, gin.H{
		"message": "book not found",
	})
}

func addBook(c *gin.Context) {
	var b Book
	c.BindJSON(&b)
	for i := range bookshelf {
		if bookshelf[i].ID == b.ID {
			c.IndentedJSON(200, gin.H{
				"message": "duplicate book id",
			})
			return
		}
	}
	bookshelf = append(bookshelf, b)
	c.IndentedJSON(200, b)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i := range bookshelf {
		if bookshelf[i].ID == id {
			c.IndentedJSON(200, bookshelf[i])
			bookshelf = remove(bookshelf, i)
			return
		}
	}
	c.IndentedJSON(200, gin.H{
		"message": "book not found",
	})
}

func updateBook(c *gin.Context) {
	var b Book
	c.BindJSON(&b)
	for i := range bookshelf {
		if bookshelf[i].ID == b.ID {
			bookshelf[i] = b
		}
	}
	c.IndentedJSON(200, b)
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/*id", updateBook)
	r.POST("/bookshelf", addBook)

	port := "8081"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
