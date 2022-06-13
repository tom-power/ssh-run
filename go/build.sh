#!/bin/bash
go test -count=1 ./sshrun/ &&
go test -count=1 ./sshrun_test/ &&
if [[ ! -d ./build ]]; then
  mkdir -p ./build
fi
go build -o ./build/_sshRun main.go