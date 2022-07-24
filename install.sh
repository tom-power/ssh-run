#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
$root/build.sh
if [[ -f $root/config.install.sh ]]; then
  source $root/config.install.sh
  chmod +x $root/sh/sshRun && cp -p $root/sh/sshRun $binPath
  chmod +x $root/go/build/_sshRun && cp -p $root/go/build/_sshRun $binPath
  if [[ -d $completionsDir ]]; then
    cp -p $root/config/_sshRunCompletion $completionsDir/_sshRun
  fi
fi
echo "finished"