#!/bin/bash
sh/build.sh &&
tar -cvzf release.tar.gz ./go/build/_sshRun ./release/*
