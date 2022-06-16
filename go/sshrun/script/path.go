package script

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
)

func (fsys FileSys) Path(host shared.Host, scriptName string) (string, error) {
	script := ""
	fileExistsFs := fsys.fileExists()
	common, _ := fsys.pathFromCommon(scriptName)
	sharedScriptPath, _ := fsys.pathFromShared(host, scriptName)
	fromHost, _ := fsys.pathFromHostLocal(host, scriptName)
	fromHostRemote := ""
	if fsys.Config.CheckRemoteForScripts {
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
