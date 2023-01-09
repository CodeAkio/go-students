package controllers

import (
	"net/http"

	"github.com/CodeAkio/go-students/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{
		"API says:": "Hi, " + name + ", how are you?",
	})
}
