package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

func listHostRemote(host shared.Host) (string, error) {
	hostDir := hostDirWithHome(host.Name, "/home/"+host.User)
	command := "find " + hostDir + " -type f -printf '%f '"
	return runCommandOn(host, command)
}
