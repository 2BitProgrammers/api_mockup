# 2bitprogrammers/api_mockup
A simple tool to quickly mockup backend APIs.  

It is meant to be used for instructional and debugging purposes only.

## Configuration File
This is where you define each of the endpoints.

The configuration file is only read once, during application start.
If you modify your endpoint configuration(s) then you must restart your app to ensure the changes take effect.

Default Filename:  **config_api_mockup.json**

Things to note about the configuration file:
* The config file can have multiple unique endpoints in which the app can parse through
* Each endpoint can have multiple unique methods in which the app can parse through
* Each response can have multiple headers which the app dynamically appends to the response
* Each response can only have one payload
* This application doesn't currently support binary data for the response payload

The basic structure of the config file is:
* **URI - map[string]** - one or more endpoint routes the app listens for in order to server a response
  * **Method - map[string]** - each route will map to one or more methods (i.e. GET, POST, UPDATE, etc.)
    * **Endpoint Response** - each route method will have its own response definition.  This contains the following:
      * **header list** - an array of key/value pairs which will be added to the header list before you send the response
      * **payload** - a text-based body which is sent in the response  

### Default Response - Not Found (404)
If the app cannot find a valid definition for the specified URI/Method combination, then it returns the default response:
* Headers:
  * Content-Type = application/json
* Payload:  ```{ "errors": [{ "status": 404, "message": "Resource Not Found", "uri": "<requestURI>", "method": "<requestMethod>" }] }```

A sample error would be:
```bash
$ curl -v http://localhost:1234/uri/does/not/exist
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1234 (#0)
> GET /uri/does/not/exist HTTP/1.1
> Host: localhost:1234
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Fri, 11 Dec 2020 21:56:34 GMT
< Content-Length: 113
<
{ "errors": [{ "status": 404, "message": "Resource Not Found", "uri": "/uri/does/not/exist", "method": "GET" }] }
* Connection #0 to host localhost left intact
```

### Default Endpoint Headers (Content-Type: text/plain)
If the user defines an endpoint response, but it has no declared headers, then it will use the following definition:
* "Content-Type" = "plain/text"


### CORS
If you are mocking an API which requires CORS enabled, you must add the the appropriate headers.
* Access-Control-Allow-Origin: *
* Access-Control-Allow-Headers: *

Example config file which allows CORS from any origin:
```json
{
    "/data": { 
        "GET": {
            "headers": [
                { "key": "Content-Type", "value": "application/json" },
                { "key": "Access-Control-Allow-Origin", "value": "*" },
                { "key": "Access-Control-Allow-Headers", "value": "*" }               
            ],
            "payload": "{ \"status\": 200, \"statusText\": \"Ok\", \"data\": [ [\"r1\", 1.0, 53], [\"r2\", 1.4, 22], [\"r3\", 3.2, 81] ] }"
        }
    }
}
```

### Configuration Examples
Here we will show some simple examples. It is not an exhaustive list, but it should enough to get you going.

#### Config Example: Ping/Pong (GET)
Our 1st example will just return a simple string:
* URI: "/ping"
* Method: "GET"
* Headers:
  * "Content-Type" = "text/plain"
* Payload: "pong"

For this example, the config file would look like:
```json
{
    "/ping": {
        "GET": {
            "headers": [
                { "key": "Content-Type", "value": "text/plain" }
            ],
            "payload": "pong"
        }
    }
}
```

Run the server and then test the endpoint:
```bash
$ curl -X POST http://localhost:1234/users
{ "code": 200, "message": "OK - User added successfully.", "data": { "id": 1111 }, "error": null  }
C:\Users\rbala>curl -v http://localhost:1234/ping
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1234 (#0)
> GET /ping HTTP/1.1
> Host: localhost:1234
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain
< Date: Fri, 11 Dec 2020 21:39:28 GMT
< Content-Length: 4
<
pong
* Connection #0 to host localhost left intact
```

#### Config Example: App Status (GET)
In this example we will return a JSON string:
* URI: "/app/status"
* Method: "GET"
* Headers:
  * "Content-Type" = "application/json"
* Payload: ```{ "appStatus": "healthy" }```

For this example, the config file would look like:
```json
{
    "/app/status": {
        "GET": {
            "headers": [
                { "key": "Content-Type", "value": "application/json" }
            ],
            "payload": "{ \"appStatus\": \"healthy\" }"
        }
    }
}
```

Run the server and then test the endpoint:
```bash
$ curl -v http://localhost:1234/app/status
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1234 (#0)
> GET /app/status HTTP/1.1
> Host: localhost:1234
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 11 Dec 2020 21:40:21 GMT
< Content-Length: 26
<
{ "appStatus": "healthy" }
* Connection #0 to host localhost left intact
```

#### Config Example: Home Page (GET)
In this example we will return a static web page (index.html):
* URI: "/index.html"
* Method: "GET"
* Headers:
  * "Content-Type" = "text/html;charset=utf-8"
* Payload: ```<!doctype html>\n<html><head><meta charset=\"utf-8\"><title>2-Bit Mockup Index Page</title></head><body><b>/index.html</b> - This is a mockup HTML page :)</body></html>```

For this example, the config file would look like:
```json
{
    "/index.html": {
        "GET": {
            "headers": [
                { "key": "Content-Type", "value": "text/html;charset=utf-8" }
            ],
            "payload": "<!doctype html>\n<html><head><meta charset=\"utf-8\"><title>2-Bit Mockup Index Page</title></head><body><b>/index.html</b> - This is a mockup HTML page :)</body></html>"
        }
    }
}
```

Run the server and then test the endpoint:
```bash
$ curl -v http://localhost:1234/index.html
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1234 (#0)
> GET /index.html HTTP/1.1
> Host: localhost:1234
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/html;charset=utf-8
< Date: Fri, 11 Dec 2020 21:49:08 GMT
< Content-Length: 165
<
<!doctype html>
<html><head><meta charset="utf-8"><title>2-Bit Mockup Index Page</title></head><body><b>/index.html</b> - This is a mockup HTML page :)</body></html>
* Connection #0 to host localhost left intact
```

#### Config Example: Users (GET and POST)
In this example we will define two things:
1. Retrieve a single user (mock data)
   * URI: "/users"
   * Method: "GET"
   * Headers:
     * "Content-Type" = "application/json"
   * Payload: ```{ "code": 200, "message": "OK", "data": { "id": 1001, "first_name": "bubba", "last_name": "joe", "email": "bubba@joe.com" }, "error": null }```
1. Return a JSON string which will simulate the response after a new user is added.  The response will include the new userID.
   * URI: "/users"
   * Method: "POST"
   * Headers:
     * "Content-Type" = "application/json"
   * Payload: ```{ "code": 200, "message": "OK - User added successfully.", "data": { "id": 1111 }, "error": null  }```


For this example, the config file would look like:
```json
{
    "/users": {
        "GET": {
            "headers": [
                { "key": "Content-Type", "value": "application/json" }
            ],
            "payload": "{ \"code\": 200, \"message\": \"OK\", \"data\": { \"id\": 1001, \"first_name\": \"bubba\", \"last_name\": \"joe\", \"email\": \"bubba@joe.com\" }, \"error\": null }"
        },

        "POST": {
            "headers": [
                { "key": "Content-Type", "value": "application/json" }
            ],
            "payload": "{ \"code\": 200, \"message\": \"OK - User added successfully.\", \"data\": { \"id\": 1111 }, \"error\": null  }"
        }
    }
}
```

Run the server and then test the endpoints:
```bash
## Simulate getting a single user record
$ curl -v http://localhost:1234/users
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1234 (#0)
> GET /users HTTP/1.1
> Host: localhost:1234
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 11 Dec 2020 21:45:45 GMT
< Content-Length: 140
<
{ "code": 200, "message": "OK", "data": { "id": 1001, "first_name": "bubba", "last_name": "joe", "email": "bubba@joe.com" }, "error": null }
* Connection #0 to host localhost left intact


## Simulate adding a new user
$ curl -v -X POST http://localhost:1234/users
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 1234 (#0)
> POST /users HTTP/1.1
> Host: localhost:1234
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 11 Dec 2020 21:43:26 GMT
< Content-Length: 99
<
{ "code": 200, "message": "OK - User added successfully.", "data": { "id": 1111 }, "error": null  }
* Connection #0 to host localhost left intact
```

## Running the App as Standalone GoLang App
To run directly from the source code:
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

1. Modify the "appPort" variable in main.go before you run the app.  For many reasons, this isn't an ideal work around.  But if you wish to do it anyways, look for the constant "appPort" in main.go.
   * change the default bind port ("1234") to the port you wish the app to listen on localhost
1. A better approach would be to run multiple docker containers and bind the container ports to your localhost (on different ports).
   * *See Below (in docker section):*  **Example: Running Multiple Containers (listening on different ports)**


## Run within Docker
This will run the components on your local system without using minikube or kubernetes.


### Building the Docker Image
If you wish to pull the latest docker image from DockerHub:
```bash
$ docker pull 2bitprogrammers/api_mockup

Using default tag: latest
latest: Pulling from 2bitprogrammers/api_mockup
Digest: sha256:d201cb21fc46b719562d9f4e7435239ca7c138672d524ca30a6ea1a2f95cfb9d
Status: Image is up to date for 2bitprogrammers/api_mockup:latest
docker.io/2bitprogrammers/api_mockup:latest
```

### Building the Docker Image
For most, you don't need to build the container.  The instructions are here for "documentation completeness" only.
```bash
$ docker build . -t 2bitprogrammers/api_mockup
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
$ docker stop api_mu_2 

## Verify that the containers are no longer running
$ docker ps
CONTAINER ID   IMAGE                        COMMAND                 CREATED         STATUS         PORTS          NAMES

## To remove the image from your local system
$ docker rmi 2bitprogrammers/api_mockup

```
