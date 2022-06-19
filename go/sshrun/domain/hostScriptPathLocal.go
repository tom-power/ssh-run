package domain

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

func (h Host) pathLocal(fsys fs.FS, scriptName string) (string, error) {
	script := ""
	hostFiles, err := h.Files(fsys)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			scriptPathSubDir, _ := h.FilePathInSubDir(fsys, hostFile.Name(), scriptName)
			shared.ReplaceIf(&script, scriptPathSubDir, fileExists(fsys))
		}
	}
	scriptPath, err := h.FilePath(fsys, scriptName)
	shared.ReplaceIf(&script, scriptPath, fileExists(fsys))
	return script, nil
}
