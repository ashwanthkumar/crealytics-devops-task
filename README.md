[![Build Status](https://travis-ci.org/ashwanthkumar/crealytics-devops-task.svg?branch=master)](https://travis-ci.org/ashwanthkumar/crealytics-devops-task)
# crealytics-devops-task

## Solution Approach
The approach I'm taking for solving the problem is to have a startup-script which would create the required user and password with sudo privileges. The `startup-script.sh` is the file that's uploaded to `gs://crealytics-devops-task/startup-script.sh`. I then configure the instance with metadata (`startup-script-url`) pointing to that location on the cloud storage. Rest is taken care by Google Compute.

We store the username and password in plain text in the metadata as well while creating the instance. The startup script then deletes this sensitive information after configuring the required user.

## Usage
We use [`glide`](https://glide.sh/) for dependency management. To build and run the binary you should run the following commands

```
make setup
make build
```

The above 2 commands will install all the necessary dependencies and build a single binary `crealytics-devops-task`. You can then run the binary using

```
./crealytics-devops-task
```

This should launch the server on the port `8080` by default if no `PORT` environment variable is set, else it would run on that port.

## API
| Endpoint | Method | Description |
| :--- | :---: | :--- |
| /healthcheck | GET | Returns a 200 response code if the service is up |
| /v1/instances/create | POST | You need to pass "user" and "password" either as query parameter or part of the JSON payload in the body. You can also control the instance type, zone, instance image, etc. See **InstanceRequest Model** for the full specification. We would wait for the instance to start and return the username, password and all the ip addresses associated the instance. |
| /v1/instances/info/:name | GET | Returns the information about the instance identified by `:name` in the path parameter. This is mostly used for debugging the instance related information. It returns the [Instance](https://godoc.org/google.golang.org/api/compute/v1#Instance) object as the response. |

### InstanceRequest Model

- `user` - Login of the user to create. Once the instance is created you can become that user via `su - ${user}`. This user also has `sudo` rights on the system.
- `password` - Password of the user that we created.
- `instance-name` - Name of the instance (VM) that we create on Google cloud. Defaults to "crealytics-devops-task-demo".
- `instance-type` - [Machine Type](https://cloud.google.com/compute/docs/machine-types) of the VM to create. Defaults to "f1-micro".
- `project-id` - Project ID against which we should create the instance. Defaults to "crealytics-devops-task".
- `image-project-id` - Project ID of the image that we want to use. Defaults to "ubuntu-os-cloud".
- `image-name` - Name of the image from `image-project-id` to use. Defaults to "ubuntu-1604-xenial-v20170815a".
- `zone` - Google Compute zone in which we should create the instance. Defaults to "us-central1-a".

## Note on Google Cloud Credentials
The app configures itself with google cloud credentials via [Application Default Credentials](https://developers.google.com/identity/protocols/application-default-credentials). Please follow the instructions to configure the credentials.

## License
http://www.apache.org/licenses/LICENSE-2.0
