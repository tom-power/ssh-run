package domain

import (
	"errors"
	"io/fs"
)

func (host Host) pathShared(fsys fs.FS, scriptName string) (string, error) {
	hostDir := hostDir(host.Name)
	hostFiles, err := fs.ReadDir(fsys, hostDir)
	if err != nil {
		return "", err
	}
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			fromShared, _ := firstFileInDir(fsys, scriptsPath+"shared/"+hostFile.Name()+"/", scriptName)
			if fileExists(fsys)(fromShared) {
				return fromShared, nil
			}

		}
	}
	return "", errors.New("nothing in shared")
}