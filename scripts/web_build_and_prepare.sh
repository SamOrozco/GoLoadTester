#!/usr/bin/env bash
export webLoc=../load_tester_frontend
cd $webLoc
webdev build
cp -r build/ ../load_tester_backend/built-web