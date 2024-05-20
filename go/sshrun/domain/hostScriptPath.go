package domain

import (
	"fmt"
	"io/fs"

	"github.com/tom-power/ssh-run/sshrun/fp"
	"github.com/tom-power/ssh-run/sshrun/fp/pred"
)

func (h Host) Path(fsys fs.FS, scriptName string) (string, error) {
	commonPath, _ := pathCommon(fsys, scriptName)
	utilsPath, _ := h.pathShared(fsys, scriptName)
	hostPath, _ := h.pathLocal(fsys, scriptName)
	hostRemotePath := ""
	if h.CheckRemote {
		hostRemotePath, _ = h.pathRemote(scriptName)
	}
	paths := []string{commonPath, utilsPath, hostPath, hostRemotePath}
	path := fp.Stream(paths).Filter(pred.IsNotEmpty).Last("")
	if path == "" {
		return "", fmt.Errorf("couldn't find path for script %s.sh on host %s", scriptName, h.Name)
	}
	return path, nil
}
