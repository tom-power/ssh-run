package domain

import (
	"errors"
	"io/fs"
)

func (h Host) pathShared(fsys fs.FS, scriptName string) (string, error) {
	hostFiles, err := h.Files(fsys)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			scriptPathShared, _ := pathInDir(fsys, sharedDir()+"/"+hostFile.Name()+"/", scriptName)
			if fileExistsIn(fsys)(scriptPathShared) {
				return scriptPathShared, nil
			}

		}
	}
	return "", errors.New("nothing in shared")
}
