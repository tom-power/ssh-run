package script

import "github.com/tom-power/ssh-run/sshrun/shared"

func getScriptPathFromHost(host shared.Host, scriptName string) string {
	scriptPath := ""
	shared.UpdateIf(&scriptPath, scriptPathFromHostLocal(host, scriptName), fileExists)
	if host.CheckForScripts {
		shared.UpdateIf(&scriptPath, scriptPathFromHostRemote(host, scriptName), isNotEmpty)
	}
	return scriptPath
}
