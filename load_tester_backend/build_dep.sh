#!/usr/bin/env bash
export location=./docker-files/DockerfileDep
export name=base_tester
docker build -t $name -f $location  .