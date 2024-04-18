#!/bin/bash
sh/build.sh &&
if [[ -f sh/.install.env ]]; then
  source sh/.install.env
  chmod +x release/sshRun && cp -p release/sshRun $binPath
  chmod +x go/build/_sshRun && cp -p go/build/_sshRun $binPath
fi
echo "finished"
