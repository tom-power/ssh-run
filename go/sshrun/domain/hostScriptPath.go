package domain

import (
	"fmt"
	"io/fs"

	"github.com/samber/lo"
	"github.com/tom-power/ssh-run/sshrun/utils"
)

func (h Host) Path(fsys fs.FS, scriptName string) (string, error) {
	commonPath, _ := pathCommon(fsys, scriptName)
	utilsPath, _ := h.pathShared(fsys, scriptName)
	hostPath, _ := h.pathLocal(fsys, scriptName)
	hostRemotePath := ""
	if h.CheckRemote {
		hostRemotePath, _ = h.pathRemote(scriptName)
	}
	paths := lo.Filter([]string{commonPath, utilsPath, hostPath, hostRemotePath}, utils.IsNotEmptyWithIndex)
	path, err := lo.Last(paths)
	if err != nil {
		return "", fmt.Errorf("couldn't find path for script %s.sh on host %s", scriptName, h.Name)
	}
	return path, nil
}
