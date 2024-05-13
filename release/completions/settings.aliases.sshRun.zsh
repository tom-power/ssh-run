#!/bin/zsh
if [ -f "/path/to/sshRun" ]; then
  for host in $(/path/to/sshRun --hosts); do
    alias "$host=sshRun $host "
  done
fi
