package domain

import (
	"io/fs"

	"github.com/tom-power/ssh-run/sshrun/utils/fp"
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
			fp.ReplaceIf(&script, scriptPathSubDir, fileExists(fsys))
		}
	}
	scriptPath, _ := pathInDir(fsys, h.Dir()+"/", scriptName)
	fp.ReplaceIf(&script, scriptPath, fileExists(fsys))
	return script, nil
}
