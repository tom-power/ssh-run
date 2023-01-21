#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
sh/build.sh &&
if [[ -f sh/config.install.sh ]]; then
  source sh/config.install.sh
  chmod +x config/sshRun && cp -p config/sshRun $binPath
  chmod +x go/build/_sshRun && cp -p go/build/_sshRun $binPath
  if [[ -d $completionsDir ]]; then
    cp -p config/_sshRun $completionsDir/_sshRun
  fi
fi
echo "finished"
