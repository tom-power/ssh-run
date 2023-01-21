#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
$root/build.sh
tar -cvzf release.tar.gz \
./config/sshRun ./go/build/_sshRun ./config/.bash_aliases ./config/_sshRunCompletion
