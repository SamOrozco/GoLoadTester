#!/usr/bin/env bash


# gentle-wave-92777
# IF NO APP CREATED
#heroku create

export appName=gentle-wave-92777

# build web client changes
./web_build_and_prepare.sh
# build local image
./build_image_locally.sh

# DEPLOY TO HEROKU
# tag desired image
docker tag local_load_tester registry.heroku.com/$appName/web

#push tag..?
docker push registry.heroku.com/$appName/web

heroku container:release web --app $appName