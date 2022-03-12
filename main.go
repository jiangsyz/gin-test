package main

import (
	"gin-test/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.PrintResponse)

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": c.Params.ByName("name")})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}