package script

import (
	"errors"
	"github.com/tom-power/ssh-run/sshrun/shared"
)

func (fsys FileSys) Path(host shared.Host, scriptName string) (string, error) {
	script := ""
	fileExistsFs := fsys.fileExists()
	fromCommon, _ := fsys.pathCommon(scriptName)
	fromShared, _ := fsys.pathShared(host, scriptName)
	fromHost, _ := fsys.pathHostLocal(host, scriptName)
	fromHostRemote := ""
	if fsys.Config.CheckRemoteForScripts {
		fromHostRemote, _ = pathHostRemote(host, scriptName)
	}
	shared.UpdateIf(&script, fromCommon, fileExistsFs)
	shared.UpdateIf(&script, fromShared, fileExistsFs)
	shared.UpdateIf(&script, fromHost, fileExistsFs)
	shared.UpdateIf(&script, fromHostRemote, isNotEmpty)
	if script == "" {
		return "", errors.New("couldn't find script \"" + scriptName + ".sh\" for host \"" + host.Name + "\"")
	}
	return script, nil
}
