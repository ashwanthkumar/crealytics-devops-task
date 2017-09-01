[![Build Status](https://travis-ci.org/ashwanthkumar/crealytics-devops-task.svg?branch=master)](https://travis-ci.org/ashwanthkumar/crealytics-devops-task)
# crealytics-devops-task

Solution to the [devops task](https://docs.google.com/document/d/18zk1WbVBPuooO_sCPwA1Y8KkHubQSPJHfBZZ7eHFgKs/edit?ts=59a81932#) given as part of interview process at [Crealytics](https://crealytics.com/career/), Devops Role.

## Solution Approach
The approach I took for solving the problem is to inject a startup-script to the VM that would create the required user with password and sudo privileges. It would pick the username and password from the instance metadata. The startup-script is configured via the metadata (`startup-script-url`) pointing it to the [file on Github](https://raw.githubusercontent.com/ashwanthkumar/crealytics-devops-task/master/startup-script.sh). Rest is taken care by Google Compute.

Given that we store the username and password in plain text in the metadata while creating the instance, the script deletes the sensitive information at the end. We also set a key on the metadata at the end of the script to make sure the script has passed fine.

While there various ways to solving this problem, I found this to be the simplest of them. Would love to learn if there are more simpler ways to do this. Some approaches I thought and rejected were

1. Once the instance is up, SSH into the machine programmatically to run the required commands. This would require a lot of complexity in terms of managing SSH connections from Go.
2. Use ansible or any other tool to configure the instance once it is up.  This would require complexity in terms of managing process lifecycle from Go. Also needs other dependencies in order to run the service.

## Usage
We use [`glide`](https://glide.sh/) for dependency management. Please make sure you've glide installed on your machine before attempting to build the project. To build and run the service you should run the following command

```
make
```

The above command will install all the necessary dependencies, build a single binary `crealytics-devops-task` and start the service. You can also then manually start the service using

```
./crealytics-devops-task
```

This should launch the server on the port `8080` by default if no `PORT` environment variable is set, else it would run on that port.

## API
| Endpoint | Method | Description |
| :--- | :---: | :--- |
| /healthcheck | GET | Returns a 200 response code if the service is up |
| /v1/instances/create | POST | You need to pass "user" and "password" either as query parameter or part of the JSON payload in the body. You can also control the instance type, zone, instance image, etc. See [InstanceRequest Model](#instancerequest-model) for the full specification. We would wait for the instance to start and return the username, password and all the ip addresses associated the instance. |
| /v1/instances/info/:name | GET | Returns the information about the instance identified by `:name` in the path parameter. It assumes the default project-id and zone as defined in [InstanceRequest Model](#instancerequest-model). You can override them via `project-id` and `zone` query parameters. This is mostly used for debugging the instance related information. It returns the [Instance](https://godoc.org/google.golang.org/api/compute/v1#Instance) object as the response. |

### InstanceRequest Model

| Name | Description | Default |
| --- | --- | --- |
| username | Login of the user to create. Once the instance is created you can become that user via `su - ${user}`. This user also has `sudo` rights on the system. | None (Required) |
| password | Password of the user that we created. | None (Required) |
| instance-name | Name of the instance (VM) that we create on Google cloud. | crealytics-devops-task-demo |
| instance-type | [Machine Type](https://cloud.google.com/compute/docs/machine-types) of the VM to create. | f1-micro |
| project-id | Project ID against which we should create the instance. | crealytics-devops-task |
| image-project-id | Project ID of the image that we want to use. | ubuntu-os-cloud |
| image-name | Name of the image from `image-project-id` to use. | ubuntu-1604-xenial-v20170815a |
| zone | Google Compute zone in which we should create the instance. | us-central1-a |
| description | Description of the instance that we create. | compute instance created via crealytics-devops-task |

## Note on Google Cloud Credentials
The app configures itself with google cloud credentials via [Application Default Credentials](https://developers.google.com/identity/protocols/application-default-credentials). Please follow the instructions to configure the credentials.

## License
http://www.apache.org/licenses/LICENSE-2.0
