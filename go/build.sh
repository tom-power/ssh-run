#!/bin/bash
go test -count=1 ./sshrunscripts/ &&
go test -count=1 ./sshrunscripts_test/ &&
if [[ ! -d ./build ]]; then
  mkdir -p ./build
fi
go build -o ./build/_sshRun main.go