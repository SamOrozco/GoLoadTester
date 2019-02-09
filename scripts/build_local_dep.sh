#!/usr/bin/env bash

# docker file location
export location=load_tester_backend/docker-files/DockerfileDep
# where we need to run the docker file from
export backend=..
# base image name
export name=base_tester
# go to location we need to run rom
cd $backend
docker build -t $name -f $location  .