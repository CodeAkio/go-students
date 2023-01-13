package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CodeAkio/go-students/controllers"
	"github.com/CodeAkio/go-students/database"
	"github.com/CodeAkio/go-students/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var StudentId int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "John Doo", CPF: "11111111111", RG: "111111111"}

	database.DB.Create(&student)

	StudentId = int(student.ID)
}

func DeleteStudentMock() {
	database.DB.Delete(&models.Student{}, StudentId)
}

func TestGetAllStudentsHandler(t *testing.T) {
	database.ConnectDb()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/api/students", controllers.GetAllStudents)

	req, _ := http.NewRequest("GET", "/api/students", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}

func TestGetStudentByCpfHandler(t *testing.T) {
	database.ConnectDb()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/api/students/cpf/:cpf", controllers.GetAllStudents)

	req, _ := http.NewRequest("GET", "/api/students/cpf/11111111111", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}
