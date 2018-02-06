#!/bin/bash
go build
./pubg_ello_data
DATE=$('date')
git add .
git commit -m "$DATE"