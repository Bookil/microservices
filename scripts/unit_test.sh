#!/bin/bash

for file in $(ls); do
    if [[ $file =~ ^(user|auth|email|product) ]]; 
    then
        cd $file
        make unit_test
        cd ../        
    fi
done