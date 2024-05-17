package domain

import (
	"io/fs"

	"github.com/tom-power/ssh-run/sshrun/shared/generic"
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
			generic.ReplaceIf(&script, scriptPathSubDir, fileExists(fsys))
		}
	}
	scriptPath, _ := pathInDir(fsys, h.Dir()+"/", scriptName)
	generic.ReplaceIf(&script, scriptPath, fileExists(fsys))
	return script, nil
}
