package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func getScriptPathFromHost(host shared.Host, scriptName string, config shared.Config, fs fs.FS) (string, error) {
	scriptPath := ""
	local, err := scriptPathFromHostLocal(host, scriptName, fs)
	if err != nil {
		return "", err
	}
	shared.UpdateIf(&scriptPath, local, fileExistsFs(fs))
	if config.CheckRemoteForScripts {
		shared.UpdateIf(&scriptPath, scriptPathFromHostRemote(host, scriptName), isNotEmpty)
	}
	return scriptPath, nil
}
