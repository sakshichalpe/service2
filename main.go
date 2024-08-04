package main

import (
	"service2/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/getData", handler.GetData)
	r.Run("localhost:8080")
}
