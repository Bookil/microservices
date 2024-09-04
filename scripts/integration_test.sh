#!/bin/bash

for file in $(ls); do
    if [[ $file =~ ^(user|auth) ]]; 
    then
        cd $file
        make db_integration_test
        cd ../        
    fi
done