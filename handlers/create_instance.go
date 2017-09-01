package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// InstanceRequest is a model to bind the incoming request to a struct
type InstanceRequest struct {
	Username       string `json:"user" binding:"required"`
	Password       string `json:"password" binding:"required"`
	InstanceName   string `json:"name,omitempty"`
	InstanceType   string `json:"type,omitempty"`
	ProjectID      string `json:"projectId,omitempty"`
	ImageProjectID string `json:"imageProjectId,omitempty"`
	ImageName      string `json:"imageName,omitempty"`
	Zone           string `json:"zone,omitempty"`
}

// Default Values for InstanceRequest
const (
	ProjectID      = "crealytics-devops-task"
	ImageProjectID = "ubuntu-os-cloud"
	ImageName      = "ubuntu-1604-xenial-v20170815a"
	Zone           = "us-central1-a"
	InstanceName   = "test-1"
	InstanceType   = "f1-micro"
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
	err := c.BindJSON(request)
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
