#!/usr/bin/env bash

# where do we want to execute this script from and go to it
export location=..
cd $location

# <location> relative docker file location
export dockerFileLocation=load_tester_backend/docker-files/Dockerfile
#desire image name
export imageName=local_load_tester
# copy built web relative <location>
cp -r load_tester_frontend/build load_tester_backend/built-web

docker build -t $imageName -f $dockerFileLocation .

docker run -p 9000:9000 $imageName