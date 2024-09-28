#!/bin/bash

for file in $(ls); do
    if [[ $file =~ ^(user|auth|product) ]]; 
    then
        cd $file
        make db_integration_test
        cd ../        
    fi
done