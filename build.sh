#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
cd $root/go && ./build.sh
