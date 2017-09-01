package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// InstanceInfo is a model to bind the incoming request for /v1/instances/info/:name
type InstanceInfo struct {
	Name      string `form:"name"`
	ProjectID string `form:"project-id"`
	Zone      string `form:"zone"`
}

// GetInstanceInfoHandler is the entry point for /v1/instances/info/:name
func GetInstanceInfoHandler(c *gin.Context) {
	name := c.Param("name")
	request := &InstanceInfo{
		Name: name,
	}
	err := c.Bind(request)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "failed",
			"msg":    fmt.Sprintf("%v", err),
		})
		return
	}

	instance, err := getInstance(request)
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
