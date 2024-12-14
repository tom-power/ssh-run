#!/bin/bash
sh/build.sh &&
if [[ -f sh/.install.env ]]; then
  source sh/.install.env
  if [ -n "${binPath+set}" ]; then
    echo "copying scripts"
    cp -p release/sshRun $binPath/sshRun && chmod +x $binPath/sshRun &&
    cp -p go/build/_sshRun $binPath/_sshRun && chmod +x $binPath/_sshRun
  fi
  if [ -n "${completionsPath+set}" ]; then
    echo "copying completions"
    cp -p release/completions/_sshRun $completionsPath/_sshRun
  fi
fi
echo "finished"
