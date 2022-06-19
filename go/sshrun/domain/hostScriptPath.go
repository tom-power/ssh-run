package domain

import (
	"errors"
	"fmt"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (host Host) Path(fsys fs.FS, scriptName string) (string, error) {
	commonPath, _ := pathCommon(fsys, scriptName)
	sharedPath, _ := host.pathShared(fsys, scriptName)
	hostPath, _ := host.pathLocal(fsys, scriptName)
	hostRemotePath := ""
	if host.CheckRemote {
		hostRemotePath, _ = host.pathRemote(scriptName)
	}
	path := shared.LastOr(shared.Filter([]string{commonPath, sharedPath, hostPath, hostRemotePath}, isNotEmpty), "")
	if path == "" {
		return "", errors.New(fmt.Sprintf("couldn't find path %s.sh for host %s", scriptName, host.Name))
	}
	return path, nil
}
