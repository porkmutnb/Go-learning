package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct
type Toy struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// in-memory data array struct
var toys = []Toy{
	{
		ID:          "1",
		Name:        "Teddy Bear",
		Description: "A soft toy bear",
		Price:       19.99,
	},
	{
		ID:          "2",
		Name:        "Toy Car",
		Description: "A fast toy car",
		Price:       9.99,
	},
}

func main() {
	r := gin.Default()
	fmt.Printf("Gin starting.... on port:8080 \n")
	// GET /
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "GO get ready!!!"})
	})
	// GET /toys
	r.GET("/toys", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": toys})
	})
	// GET /toys/:id
	r.GET("/toys/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, toy := range toys {
			if toy.ID == id {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": toy})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Toy not found"})
	})
	// POST /toys
	r.POST("/toys", func(c *gin.Context) {
		var toy Toy
		if err := c.BindJSON(&toy); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid request"})
			return
		}
		toys = append(toys, toy)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": toy})
	})
	// PUT /toys/:id
	r.PUT("/toys/:id", func(c *gin.Context) {
		id := c.Param("id")
		var toy Toy
		if err := c.BindJSON(&toy); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid request"})
			return
		}
		for i, t := range toys {
			if t.ID == id {
				toys[i] = toy
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": toy})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Toy not found"})
	})
	// DELETE /toys/:id
	r.DELETE("/toys/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, t := range toys {
			if t.ID == id {
				toys = append(toys[:i], toys[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Toy not found"})
	})
	r.Run(":8080")
}
