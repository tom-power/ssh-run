package domain

import (
	"io/fs"
)

func (h Host) pathLocal(fsys fs.FS, scriptName string) (string, error) {
	script := ""
	hostFiles, err := h.Files(fsys)
	if err != nil {
		return "", err
	}
	fileExists := fileExistsIn(fsys)
	for _, file := range hostFiles {
		if file.IsDir() {
			scriptPathSubDir, _ := pathInDir(fsys, h.Dir()+"/"+file.Name()+"/", scriptName)
			if fileExists(scriptPathSubDir) {
				script = scriptPathSubDir
			}
		}
	}
	scriptPath, _ := pathInDir(fsys, h.Dir()+"/", scriptName)
	if fileExists(scriptPath) {
		script = scriptPath
	}
	return script, nil
}
