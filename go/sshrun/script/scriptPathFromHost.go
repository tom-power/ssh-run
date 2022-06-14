package script

import "github.com/tom-power/ssh-run/sshrun/shared"

func getScriptPathFromHost(host shared.Host, scriptName string, config shared.Config) string {
	scriptPath := ""
	shared.UpdateIf(&scriptPath, scriptPathFromHostLocal(host, scriptName), fileExists)
	if config.CheckRemoteForScripts {
		shared.UpdateIf(&scriptPath, scriptPathFromHostRemote(host, scriptName), isNotEmpty)
	}
	return scriptPath
}
