#!/bin/bash
cd ~/go/src/pubg_ello_data/
ssh git@github.com 
git pull
go build
./pubg_ello_data
DATE=$('date')
cd ../../../pubg_project_site
ssh git@github.com
git pull
git add .
git commit -m "$DATE"
git push
