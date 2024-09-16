#!/bin/bash

service=$1

echo service:$service

cd $service

mkdir -p mocks

for file in $(ls ./internal/ports/*.go); do
  filename=$(basename "$file" .go)
  if [[ ! $filename =~ ^(api|validation) ]]; then
    echo generate $filename mocks....
    mockgen -package=mocks -source="$file" -destination="./mocks/${filename}_mock.go"
  fi
done