#!/usr/bin/env bash
export imageName=local_load_tester
docker build -t $imageName .
docker r tun $imageName