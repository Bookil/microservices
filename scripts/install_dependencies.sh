#!/bin/bash

for file in $(ls); do
    if [[ $file =~ ^(user|auth) ]]; 
    then
        cd $file
        go mod tidy
        cd ../        
    fi
done