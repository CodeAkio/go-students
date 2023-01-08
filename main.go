package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "Pedro",
	})
}

func main() {
	r := gin.Default()

	r.GET("/api/students", getAllStudents)

	r.Run()
}
