#!/bin/bash
cd ~/go/src/pubg_ello_data/
git pull
go build
./pubg_ello_data
DATE=$('date')
git add .
git commit -m "$DATE"
git push