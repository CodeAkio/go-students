package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("🚀 Running server at: http://localhost:8080")
	r.Run()
}
