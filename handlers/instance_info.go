package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetInstanceInfoHandler is the entry point for /v1/instances/info/:name
func GetInstanceInfoHandler(c *gin.Context) {
	name := c.Param("name")
	instance, err := getInstance(name)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "failed",
			"msg":    fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "OK",
			"result": instance,
		})
	}
}
