package handlers

import (
	"fmt"
	"time"

	"github.com/ashwanthkumar/crealytics-devops-task/gclient"
	compute "google.golang.org/api/compute/v1"
)

var scopes = []string{
	compute.DevstorageFullControlScope,
	compute.ComputeScope,
}

func createInstanceAndGetIPAddresses(request *InstanceRequest) ([]string, error) {
	client, err := gclient.CreateOrGetClient(scopes)
	if err != nil {
		fmt.Printf("%v\n", err)
		return []string{}, err
	}
	service, err := compute.New(client)
	if err != nil {
		fmt.Printf("%v\n", err)
		return []string{}, err
	}

	prefix := "https://www.googleapis.com/compute/v1/projects/" + request.ProjectID
	imageURL := "https://www.googleapis.com/compute/v1/projects/" + request.ImageProjectID + "/global/images/" + request.ImageName

	instance := &compute.Instance{
		Name:        request.InstanceName,
		Description: "compute sample instance",
		MachineType: prefix + "/zones/" + request.Zone + "/machineTypes/" + request.InstanceType,
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    request.InstanceName + "-root-pd",
					SourceImage: imageURL,
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				AccessConfigs: []*compute.AccessConfig{
					&compute.AccessConfig{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
				Network: prefix + "/global/networks/default",
			},
		},
		ServiceAccounts: []*compute.ServiceAccount{
			{
				Email:  "default",
				Scopes: scopes,
			},
		},
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				item("startup-script-url", "gs://crealytics-devops-task/startup-script.sh"),
				item("custom-user", request.Username),
				item("custom-user-passwd", request.Password),
			},
		},
	}

	_, err = service.Instances.Insert(request.ProjectID, request.Zone, instance).Do()
	if err != nil {
		return []string{}, err
	}

	inst := waitForInstanceToStart(request, service)
	var allIPAddress []string
	for _, netInterface := range inst.NetworkInterfaces {
		allIPAddress = append(allIPAddress, netInterface.NetworkIP)
	}

	return allIPAddress, err
}

func waitForInstanceToStart(request *InstanceRequest, service *compute.Service) *compute.Instance {
	stillStarting := true
	initialWaitDuration := time.Duration(0)
	var inst *compute.Instance
	for stillStarting {
		inst, _ = service.Instances.Get(request.ProjectID, request.Zone, request.InstanceName).Do()
		if inst != nil && inst.Status == "RUNNING" {
			stillStarting = false
		} else {
			initialWaitDuration = initialWaitDuration + 1000
			fmt.Printf("%s is still starting up, waiting for %d milliseconds\n", inst.Name, initialWaitDuration)
			time.Sleep(initialWaitDuration * time.Millisecond)
		}
	}

	return inst
}

func getInstance(instanceName string) (*compute.Instance, error) {
	client, err := gclient.CreateOrGetClient(scopes)
	if err != nil {
		return nil, err
	}

	service, err := compute.New(client)
	if err != nil {
		return nil, err
	}

	inst, err := service.Instances.Get(ProjectID, Zone, instanceName).Do()
	if err != nil {
		return nil, err
	}

	return inst, nil
}

func item(key, value string) *compute.MetadataItems {
	return &compute.MetadataItems{
		Key:   key,
		Value: &value,
	}
}
