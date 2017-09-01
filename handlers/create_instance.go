package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// InstanceRequest is a model to bind the incoming request for /v1/create/instances
type InstanceRequest struct {
	Username       string `form:"username" json:"username" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	InstanceName   string `form:"instance-name,omitempty" json:"instance-name,omitempty"`
	InstanceType   string `form:"instance-type,omitempty" json:"instance-type,omitempty"`
	ProjectID      string `form:"project-id,omitempty" json:"project-id,omitempty"`
	ImageProjectID string `form:"image-project-id,omitempty" json:"image-project-id,omitempty"`
	ImageName      string `form:"image-name,omitempty" json:"image-name,omitempty"`
	Zone           string `form:"zone,omitempty" json:"zone,omitempty"`
}

// Default Values for InstanceRequest
const (
	ProjectID      = "crealytics-devops-task"
	ImageProjectID = "ubuntu-os-cloud"
	ImageName      = "ubuntu-1604-xenial-v20170815a"
	Zone           = "us-central1-a"
	InstanceName   = "crealytics-devops-task-demo"
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
