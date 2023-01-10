package routes

import (
	"github.com/CodeAkio/go-students/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:id", controllers.GetStudentById)
	r.POST("/api/students", controllers.CreateStudent)

	r.Run()
}
