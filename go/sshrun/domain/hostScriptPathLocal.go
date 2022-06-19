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
	for _, file := range hostFiles {
		if file.IsDir() {
			scriptPathSubDir, _ := pathInDir(fsys, h.Dir()+"/"+file.Name()+"/", scriptName)
			shared.ReplaceIf(&script, scriptPathSubDir, fileExists(fsys))
		}
	}
	scriptPath, err := pathInDir(fsys, h.Dir()+"/", scriptName)
	shared.ReplaceIf(&script, scriptPath, fileExists(fsys))
	return script, nil
}
