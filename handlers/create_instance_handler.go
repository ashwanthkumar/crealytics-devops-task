package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateInstanceHandler is the entry point for /v1/instances/create API
func CreateInstanceHandler(c *gin.Context) {
	// Start off with defaults for certain properties, override them during BindJSON
	request := &InstanceRequest{
		InstanceName:   InstanceName,
		InstanceType:   InstanceType,
		ProjectID:      ProjectID,
		ImageProjectID: ImageProjectID,
		ImageName:      ImageName,
		Zone:           Zone,
	}
	err := c.Bind(request)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
			"msg":    fmt.Sprintf("%v", err),
		})
		return
	}

	ipAddresses, err := createInstanceAndGetIPAddresses(request)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "failed",
			"msg":    fmt.Sprintf("%s", err),
		})
	} else {
		c.JSON(200, gin.H{
			"status":   "OK",
			"username": request.Username,
			"password": request.Password,
			"ips":      ipAddresses,
		})
	}
}
