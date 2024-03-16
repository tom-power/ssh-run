#!/bin/bash
cd ./go && 
go test -count=1 ./... &&
if [[ ! -d ./build ]]; then
  mkdir -p ./build
fi
go build -o ./build/_sshRun main.go
