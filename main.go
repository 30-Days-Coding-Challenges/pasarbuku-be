package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":      "Muhammad Wage Juli Saputra",
			"job_title": "Software Engineer",
			"greet":     "Welcome aboard! Enjoy your adventure",
		})
	})

	router.Run()

}
