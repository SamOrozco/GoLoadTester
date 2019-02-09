#!/usr/bin/env bash
export dockerFileLocation=docker-files/Dockerfile
export imageName=local_load_tester
docker build -t $imageName -f $dockerFileLocation .
docker run $imageName