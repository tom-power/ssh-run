package domain

import (
	"fmt"
	"io/fs"

	"github.com/tom-power/ssh-run/sshrun/utils"
	"github.com/tom-power/ssh-run/sshrun/utils/fp"
)

func (h Host) Path(fsys fs.FS, scriptName string) (string, error) {
	commonPath, _ := pathCommon(fsys, scriptName)
	utilsPath, _ := h.pathShared(fsys, scriptName)
	hostPath, _ := h.pathLocal(fsys, scriptName)
	hostRemotePath := ""
	if h.CheckRemote {
		hostRemotePath, _ = h.pathRemote(scriptName)
	}
	path := fp.LastOr(fp.Filter([]string{commonPath, utilsPath, hostPath, hostRemotePath}, utils.IsNotEmpty), "")
	if path == "" {
		return "", fmt.Errorf("couldn't find path %s.sh for h %s", scriptName, h.Name)
	}
	return path, nil
}
