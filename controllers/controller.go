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

func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func GetStudentByCpf(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})

		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno deletado com sucesso",
	})
}

func ShowIndexPage(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}
