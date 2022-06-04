package script

import "github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"

func getScriptPathFromHost(host shared.Host, scriptName string) string {
	hostScriptPath := ""
	hostScriptPathLocal := scriptPathFromHostLocal(host, scriptName)
	if fileExists(hostScriptPathLocal) {
		hostScriptPath = hostScriptPathLocal
	}
	if host.CheckForScripts {
		hostScriptPathRemote := scriptPathFromHostRemote(host, scriptName)
		if hostScriptPathRemote != "" {
			hostScriptPath = hostScriptPathRemote
		}
	}
	return hostScriptPath
}
