#!/bin/bash

for file in $(ls); do
    if [[ $file =~ ^(user|auth|email) ]]; 
    then
        cd $file
        make unit_test
        cd ../        
    fi
done