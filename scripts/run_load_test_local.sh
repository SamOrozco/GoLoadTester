#!/usr/bin/env bash
export dockerFileLocation=../load_tester_backend/docker-files/Dockerfile
export imageName=local_load_tester
cp -r load_tester_frontend/build load_tester_backend/built-web
docker build -t $imageName -f $dockerFileLocation .
docker run -p 9000:9000 $imageName