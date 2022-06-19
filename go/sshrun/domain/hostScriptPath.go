package domain

import (
	"errors"
	"fmt"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (h Host) Path(fsys fs.FS, scriptName string) (string, error) {
	commonPath, _ := pathCommon(fsys, scriptName)
	sharedPath, _ := h.pathShared(fsys, scriptName)
	hostPath, _ := h.pathLocal(fsys, scriptName)
	hostRemotePath := ""
	if h.CheckRemote {
		hostRemotePath, _ = h.pathRemote(scriptName)
	}
	path := shared.LastOr(shared.Filter([]string{commonPath, sharedPath, hostPath, hostRemotePath}, shared.IsNotEmpty), "")
	if path == "" {
		return "", errors.New(fmt.Sprintf("couldn't find path %s.sh for h %s", scriptName, h.Name))
	}
	return path, nil
}
