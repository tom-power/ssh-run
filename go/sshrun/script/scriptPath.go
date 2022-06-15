package script

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

type GetScriptPath = func(host shared.Host, scriptName string, config shared.Config) (string, error)

func GetScriptPathFromConf(fs fs.FS) GetScriptPath {
	return func(host shared.Host, scriptName string, config shared.Config) (string, error) {
		script := ""
		fileExistsFs := fileExists(fs)
		common, _ := scriptPathFromCommon(scriptName, fs)
		sharedScriptPath, _ := scriptPathFromShared(host, scriptName, fs)
		fromHost, _ := scriptPathFromHostLocal(host, scriptName, fs)
		fromHostRemote := ""
		if config.CheckRemoteForScripts {
			fromHostRemote, _ = scriptPathFromHostRemote(host, scriptName)
		}
		shared.UpdateIf(&script, common, fileExistsFs)
		shared.UpdateIf(&script, sharedScriptPath, fileExistsFs)
		shared.UpdateIf(&script, fromHost, fileExistsFs)
		shared.UpdateIf(&script, fromHostRemote, isNotEmpty)
		if script == "" {
			return "", errors.New("couldn't find script \"" + scriptName + ".sh\" for host \"" + host.Name + "\"")
		}
		return script, nil
	}
}
