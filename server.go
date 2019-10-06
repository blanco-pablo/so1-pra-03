package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/processes", func(c *gin.Context) {
		processes := getLinuxProcesses()
		c.JSON(http.StatusOK, processes);
	})


	_ = r.Run(":8080") // listen and serve on 0.0.0.0:8080
}