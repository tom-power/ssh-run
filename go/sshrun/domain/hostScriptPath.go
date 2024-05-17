package domain

import (
	"fmt"
	"io/fs"

	"github.com/tom-power/ssh-run/sshrun/shared"
	"github.com/tom-power/ssh-run/sshrun/shared/generic"
)

func (h Host) Path(fsys fs.FS, scriptName string) (string, error) {
	commonPath, _ := pathCommon(fsys, scriptName)
	sharedPath, _ := h.pathShared(fsys, scriptName)
	hostPath, _ := h.pathLocal(fsys, scriptName)
	hostRemotePath := ""
	if h.CheckRemote {
		hostRemotePath, _ = h.pathRemote(scriptName)
	}
	path := generic.LastOr(generic.Filter([]string{commonPath, sharedPath, hostPath, hostRemotePath}, shared.IsNotEmpty), "")
	if path == "" {
		return "", fmt.Errorf("couldn't find path %s.sh for h %s", scriptName, h.Name)
	}
	return path, nil
}
