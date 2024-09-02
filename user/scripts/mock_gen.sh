#!/bin/bash

mkdir -p mocks

for file in `ls ./internal/ports/*.go`; do
  filename=$(basename "$file" .go)

  if [[ ! $filename =~ ^(api|validation) ]]; then
    mockgen -package=mocks -source="$file" -destination=mocks/"$filename"_mock.go
  fi
done