package script

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
)

func scriptsFromHostRemote(host shared.Host) (string, error) {
	hostDir := hostDir(host.Name, "/home/"+host.User)
	command := "find " + hostDir + " -type f -printf '%f '"
	return runCommandOn(host, command)
}
