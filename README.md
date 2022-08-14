# Introduction
This repository is the continuous tutorial I am providing myself on deploying a go service into the clooud

# Completed(Reverse Chronological order):
* Add tests for endpoints and mock the database hits
* Create and Read with database
* Write Kubernetes configuration to run the static endpoint
* Push the image to docker hub 
* Test the endpoint is working in docker container
* Create a Service that has a single GET endpoint at port 8001

# TODO:
* Read Db connection string from the ConfigMap
* Add contract Tests for POST endpoint
* Add contract Tests for GET endpoint
* Update to Go standard template
* Deploy on AWS

# Learnings
* port number in the main file should not contain host like 127.0.0.1
* kubernetes load balancer with target port 80 does not work on local kubernetes cluster
* In DockerFile, the contents of folders are copied not the folder name, so specify same folder name to maintain structure
* The sqlite db requires ```RUN apk add build-base``` or it throws error
```"gcc" executable not found```

# Setup
* Clone the repo
* Build the image
```$ docker build -t api_blank:v1.2 .```
* Publish to Docker hub
   ```
   $ docker tag api_blank:v1.2 rjtmalik/docker_go_demo:v1.2
   $ docker push rjtmalik/docker_go_demo:v1.2
   ```

# Debug 
* To debug the built image you can enter the container with
```docker run -it api_blank:v1.1 /bin/sh```
* To run the docker container 
``` docker run -d -p 8001:8001 api_blank ```

# Health Check
```curl localhost:8001/products/1```
