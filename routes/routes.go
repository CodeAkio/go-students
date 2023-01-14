package routes

import (
	"github.com/CodeAkio/go-students/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/api/students", controllers.GetAllStudents)
	r.GET("/api/students/:id", controllers.GetStudentById)
	r.GET("/api/students/cpf/:cpf", controllers.GetStudentByCpf)
	r.POST("/api/students", controllers.CreateStudent)
	r.PATCH("/api/students/:id", controllers.UpdateStudent)
	r.DELETE("/api/students/:id", controllers.DeleteStudent)

	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.NotFoundPage)

	r.Run()
}
