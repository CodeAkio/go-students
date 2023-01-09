package routes

import (
	"github.com/CodeAkio/go-students/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/api/students", controllers.GetAllStudents)

	r.Run()
}
