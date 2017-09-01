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

## Note on Google Cloud Credentials
The app configures itself with google cloud credentials via [Application Default Credentials](https://developers.google.com/identity/protocols/application-default-credentials). Please follow the instructions to configure the credentials.

## License
http://www.apache.org/licenses/LICENSE-2.0
