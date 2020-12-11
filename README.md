# 2bitprogrammers/api_mockup
A simple tool to quickly mockup backend APIs.  
It is meant to be used for instructional and debugging purposes only.

## Configuration File
This is where you define each of the endpoints. 

The configuration file is only read once, during application start.
If you modify your endpoint configuration(s) then you must restart your app to ensure the changes take effect.

Default Filename:  config_api_mockup.json

The basic structure of the config file is:
* **URI - map[string]** - the endpoint route which the app listens for in order to server a response
  * **Method - map[string]** - each route will map to various methods (i.e. GET, POST, UPDATE, etc.)
    * **Endpoint Response** - each route method will have its own response definition.  This contains the following:
      * **header list** - an array of key/value pairs which will be added before you send the response
      * **payload** - the text-based body which is sent in the response  


### Default Response - Not Found (404)
If the app cannot find a valid definition for the specifiec URI/Method combo, then it returns the default response:
* Headers:
  * Content-Type = application/json
* Payload:  ```{ "errors": [{ "status": 404, "message": "Resource Not Found", "uri": "<requestURI>", "method": "<requestMethod>" }] }```


### Default Endpoint Headers (Content-Type: text/plain)
If the user defines an endpoint response, but it has no declared headers, then it will use the following:
* Content-Type: plain/text

### Endpoint Examples
This will retu


## Running as Standalone GoLang App
To run directly from the source code
```bash
$ cd src/
$ go run main.go 

Loading config file:  config_api_mockup.json
2bitprogrammers/api_mockup v2018.31a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2020/12/11 12:04:39 Starting App on Port 1234
2020/12/11 12:05:00 200 GET     /ping
2020/12/11 12:05:22 200 GET     /app/status
2020/12/11 12:05:42 200 GET     /users
2020/12/11 12:06:42 200 POST    /users
2020/12/11 12:07:47 404 GET     /uri/does/not/exist

CTRL+C
```

## Running Multiple Mock APIs on Different Ports 
The original app wasn't designed to handle this use case. 
With that in mind, there are a couple different ways to achieve this.

1. Modify the "appPort" variable in main.go and then run a 2nd instance of the application
   * const appPort = "1234"
     * change "1234" to the port you wish the app to listen on
1. Run multiple docker containers and bind the container ports to your localhost (on different ports). 
   * See Below (in docker section):  Example: Running Multiple Containers (listening on different ports)


## Run within Docker 
This will run the components on your local system without using minikube or kubernetes.

### Building the Docker Image
```bash
docker build . -t 2bitprogrammers/api_mockup
Sending build context to Docker daemon  122.4kB
Step 1/12 : FROM golang:alpine AS builder
 ---> b3bc898ad092
Step 2/12 : ENV GO111MODULE=on     CGO_ENABLED=0     GOOS=linux     GOARCH=amd64
 ---> Running in be436e4f992a
Removing intermediate container be436e4f992a
 ---> 02a30406951d
Step 3/12 : WORKDIR /build
 ---> Running in fe1b45a5b7cf
Removing intermediate container fe1b45a5b7cf
 ---> ec68a845a448
Step 4/12 : COPY $PWD/src/go.mod .
 ---> 74c92b30b6cb
Step 5/12 : COPY $PWD/src/main.go .
 ---> e1b4c32b5d25
Step 6/12 : RUN go mod download
 ---> Running in 1d94af5a7c7b
Removing intermediate container 1d94af5a7c7b
 ---> 50879a21c7d2
Step 7/12 : RUN go build -o api_mockup .
 ---> Running in cc736f568f8b
Removing intermediate container cc736f568f8b
 ---> 896dd6c9ce7e
Step 8/12 : FROM scratch
 --->
Step 9/12 : WORKDIR /
 ---> Running in 43fa2f2d5143
Removing intermediate container 43fa2f2d5143
 ---> fd0717420536
Step 10/12 : COPY --from=builder /build/api_mockup .
 ---> e07b37003689
Step 11/12 : COPY $PWD/src/config_api_mockup.json .
 ---> 470d251d8112
Step 12/12 : ENTRYPOINT [ "/api_mockup" ]
 ---> Running in 0461d121c87e
Removing intermediate container 0461d121c87e
 ---> 55a266454390
Successfully built 55a266454390
Successfully tagged 2bitprogrammers/api_mockup:latest
SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is recommended to double check and reset permissions for sensitive files and directories.
```



### Image Status
```bash
$ docker images

REPOSITORY                     TAG        IMAGE ID       CREATED          SIZE
2bitprogrammers/api_mockup     latest     55a266454390   16 minutes ago   6.67MB
```

### Running the Container
```bash
$ docker run --rm --name "api_mockup" -p 1234:1234 2bitprogrammers/api_mockup

Loading config file:  config_api_mockup.json

2bitprogrammers/api_mockup v2018.31a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2020/12/11 21:13:22 Starting App on Port 1234

CTRL+C
```

### Check the Container Status (docker)
```bash
$ docker ps

CONTAINER ID   IMAGE                         COMMAND             CREATED          STATUS          PORTS                     NAMES
63af12438b67   2bitprogrammers/api_mockup    "/api_mockup"       35 seconds ago   Up 34 seconds   0.0.0.0:1234->1234/tcp    api_mockup
```

### Watch Container Logs
```bash
$ docker logs -f 2bitprogrammers/api_mockup

Loading config file:  config_api_mockup.json

2bitprogrammers/api_mockup v2018.31a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

2020/12/11 21:13:22 Starting App on Port 1234
2020/12/11 21:15:51 404 GET     /url/does/not/exist
2020/12/11 21:15:58 200 GET     /ping
2020/12/11 21:16:03 200 GET     /app/status
2020/12/11 21:16:18 200 GET     /users
2020/12/11 21:16:18 200 POST    /users
```

### Stopping the Container
```bash
$ docker stop api_mockup
```


### Example: Running Multiple Containers (listening on different ports)
You will need to either build or pull the container image from docker hub.  Once the image is pulled, you need to launch each container with the appropriate mapping of config file and port(s).

In this example, we wil use two separate config files to show that each container is distinct with their endpoint response.  Both will return a "text/plain" response on the URI "/app".  Response text will be:
* Container 1:  "app_1 is the Best!!!"
* Container 2:  "app_2 makes all others drool :)"


Here are the commands to complete this demo:
```bash
## The examples/ folder contains the required files for this demo
$ cd examples/

## Show the contents for app1 config
$ cat config_api_mu_mc_1.json
{
    "/app": { 
        "GET": {
            "headers": [
                { "key": "Content-Type", "value": "text/plain" }
            ],
            "payload": "app_1 is the Best!!!"
        }
    }
}


## Show the contents for app2 config
$ cat config_api_mu_mc_2.json
{
    "/app": { 
        "GET": {
            "headers": [
                { "key": "Content-Type", "value": "text/plain" }
            ],
            "payload": "app_2 makes all others drool :)"
        }
    }
}

## Pull image from docker hub registry
$ docker pull 2bitprogrammers/api_mockup

## Launch the 1s5 container (on port 1111)
## Docker image name:  2bitprogrammers/api_mockup
## Custom container name:  api_mu_1
## Config file (on container):  /config_api_mockup.json
## Config file (local filesystem):  config_api_mu_mc_1.json
## Listen port (on the container):  1234
## Listen por (localhost):  1111
$ docker run -d --rm --name api_mu_1 -v "$PWD/config_api_mu_mc_1.json:/config_api_mockup.json" -p "1111:1234" 2bitprogrammers/api_mockup

## Launch the 2nd container (on port 2222)
## Docker image name:  2bitprogrammers/api_mockup
## Custom container name:  api_mu_2
## Config file (on container):  /config_api_mockup.json
## Config file (local filesystem):  config_api_mu_mc_2.json
## Listen port (on the container):  1234
## Listen por (localhost):  2222
$ docker run -d --rm --name api_mu_2 -v "$PWD/config_api_mu_mc_2.json:/config_api_mockup.json" -p "2222:1234" 2bitprogrammers/api_mockup


## Make sure both containers are running
$ docker ps
CONTAINER ID   IMAGE                        COMMAND                 CREATED          STATUS          PORTS                     NAMES
63af12438b67   2bitprogrammers/api_mockup   "/api_mockup"           35 seconds ago   Up 34 seconds   0.0.0.0:2222->1234/tcp    api_mu_2
1a01eef229cd   2bitprogrammers/api_mockup   "/api_mockup"           49 seconds ago   Up 47 seconds   0.0.0.0:1111->1234/tcp    api_mu_1

## Test endppoint for 1st container
$ curl -v http://localhost:1111/app
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1111 (#0)
> GET /app HTTP/1.1
> Host: localhost:1111
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain
< Date: Fri, 11 Dec 2020 21:05:40 GMT
< Content-Length: 20
<
app_1 is the Best!!!
* Connection #0 to host localhost left intact

## Test endppoint for 2nd container
$ curl -v http://localhost:2222/app
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 2222 (#0)
> GET /app HTTP/1.1
> Host: localhost:2222
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain
< Date: Fri, 11 Dec 2020 21:06:03 GMT
< Content-Length: 31
<
app_2 makes all others drool :)
* Connection #0 to host localhost left intact


## When done, stop the containers
$ docker stop api_mu_1 
$ docker stop api_mu_1 

## Verify that the containers are no longer running
$ docker ps
CONTAINER ID   IMAGE                        COMMAND                 CREATED         STATUS         PORTS          NAMES

## To remove the image from your local system
$ docker rmi 2bitprogrammers/api_mockup

```
