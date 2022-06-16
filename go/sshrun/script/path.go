package script

import (
	"errors"
	"fmt"
	"github.com/tom-power/ssh-run/sshrun/shared"
)

func (fsys FileSys) Path(host shared.Host, scriptName string) (string, error) {
	commonPath, _ := fsys.pathCommon(scriptName)
	sharedPath, _ := fsys.pathShared(host, scriptName)
	hostPath, _ := fsys.pathHostLocal(host, scriptName)
	hostRemotePath := ""
	if fsys.Config.CheckRemoteForScripts {
		hostRemotePath, _ = pathHostRemote(host, scriptName)
	}

	path := shared.LastOr(shared.Filter([]string{commonPath, sharedPath, hostPath, hostRemotePath}, isNotEmpty), "")
	if path == "" {
		return "", errors.New(fmt.Sprintf("couldn't find path %s.sh for host %s", scriptName, host.Name))
	}
	return path, nil
}
