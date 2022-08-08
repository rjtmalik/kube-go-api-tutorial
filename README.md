# Introduction
This repository is the continuous tutorial I am providing myself on deploying a go service into the clooud

# Completed(Reverse Chronological order):
* Pushed the image to docker hub 
* Tested the endpoint is working in docker container
* Create a Service that has a single GET endpoint at port 8000

# TODO:
* Write Kubernetes configuration to run the static endpoint
* Deploy on AWS


# Setup
* Clone the repo
* Build the image
```$ docker build -t api_blank:v1.1```
* Publish to Docker hub
   ```
   $ docker tag api_blank:v1 rjtmalik/docker_go_demo:v1
   $ docker push rjtmalik/docker_go_demo:v1
   ```
   
