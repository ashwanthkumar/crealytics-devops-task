package main

import (
	"github.com/ashwanthkumar/crealytics-devops-task/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	r.POST("/v1/instances/create", handlers.CreateInstanceHandler)
	r.GET("/v1/instances/info/:name", handlers.GetInstanceInfoHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
