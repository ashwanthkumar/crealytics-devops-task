package handlers

import compute "google.golang.org/api/compute/v1"

// InstanceInfo is a model to bind the incoming request for /v1/instances/info/:name
type InstanceInfo struct {
	Name      string `form:"name"`
	ProjectID string `form:"project-id"`
	Zone      string `form:"zone"`
}

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
	Description    string `form:"description,omitempty" json:"description,omitempty"`
}

// Default Values for types
const (
	ProjectID      = "crealytics-devops-task"
	ImageProjectID = "ubuntu-os-cloud"
	ImageName      = "ubuntu-1604-xenial-v20170815a"
	Zone           = "us-central1-a"
	InstanceName   = "crealytics-devops-task-demo"
	InstanceType   = "f1-micro"
	Description    = "compute instance created via crealytics-devops-task"

	CustomUserPassword = "custom-user-passwd"
	CustomUsername     = "custom-user"

	CustomScriptKey = "custom-startup-script-status"
	ScriptStarted   = "started"
	ScriptPassed    = "OK"
	ScriptFailed    = "failed"
)

var scopes = []string{
	compute.DevstorageFullControlScope,
	compute.ComputeScope,
}
