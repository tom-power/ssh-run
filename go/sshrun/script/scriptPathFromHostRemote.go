package script

import "github.com/tom-power/ssh-run/sshrun/shared"

func scriptPathFromHostRemote(host shared.Host, scriptName string) string {
	hostDir := hostDir(host.Name, "/home/"+host.User)
	command := "" +
		"cd " + hostDir + " &&" +
		"find . -type f -name '" + scriptName + "*.sh' | sed 's/\\.\\///'"
	hostScriptPathRemote, err := runCommandOn(host, command)
	if err != nil {
		return ""
	}
	return hostDir + hostScriptPathRemote
}
