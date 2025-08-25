package main

import (
	// "fmt"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	api := r.Group("/api/v0")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, "Hello world!")
		})
	}

	r.Run()
}
