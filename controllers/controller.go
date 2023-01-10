package controllers

import (
	"net/http"

	"github.com/CodeAkio/go-students/database"
	"github.com/CodeAkio/go-students/models"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{
		"API says:": "Hi, " + name + ", how are you?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}
