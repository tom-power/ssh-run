package domain

import (
	"errors"
	"fmt"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (config Config) Path(fsys fs.FS, host Host, scriptName string) (string, error) {
	commonPath, _ := config.pathCommon(fsys, scriptName)
	sharedPath, _ := config.pathShared(fsys, host, scriptName)
	hostPath, _ := config.pathHostLocal(fsys, host, scriptName)
	hostRemotePath := ""
	if config.CheckRemoteForScripts {
		hostRemotePath, _ = config.pathHostRemote(host, scriptName)
	}

	path := shared.LastOr(shared.Filter([]string{commonPath, sharedPath, hostPath, hostRemotePath}, isNotEmpty), "")
	if path == "" {
		return "", errors.New(fmt.Sprintf("couldn't find path %s.sh for host %s", scriptName, host.Name))
	}
	return path, nil
}
