# crealytics-devops-task

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

## License
http://www.apache.org/licenses/LICENSE-2.0

