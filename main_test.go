package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	r.GET("/api/students/cpf/:cpf", controllers.GetStudentByCpf)

	req, _ := http.NewRequest("GET", "/api/students/cpf/11111111111", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}

func TestGetStudentByIdHandler(t *testing.T) {
	database.ConnectDb()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/api/students/:id", controllers.GetStudentById)

	path := "/api/students/" + strconv.Itoa(StudentId)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var studentMock models.Student
	json.Unmarshal(res.Body.Bytes(), &studentMock)

	assert.Equal(t, "John Doo", studentMock.Name, "Deveriam ser iguais")
	assert.Equal(t, "11111111111", studentMock.CPF, "Deveriam ser iguais")
	assert.Equal(t, "111111111", studentMock.RG, "Deveriam ser iguais")
	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDb()

	CreateStudentMock()

	r := SetupTestRoutes()
	r.DELETE("/api/students/:id", controllers.DeleteStudent)

	path := "/api/students/" + strconv.Itoa(StudentId)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}
