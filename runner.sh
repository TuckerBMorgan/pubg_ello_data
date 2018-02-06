#!/bin/bash
go build
./pubg_ello_data
DATE=(date +%Y-%m-%d)
git add .
git commit -m "$DATE"