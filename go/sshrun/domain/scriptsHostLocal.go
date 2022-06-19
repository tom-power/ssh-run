package domain

import (
	"io/fs"
)

func (host Host) scriptsLocal(fsys fs.FS) (string, error) {
	var files = Files{[]fs.DirEntry{}}
	err := fs.WalkDir(fsys, hostDir(host.Name), appendFiles(&files.Files))
	if err != nil {
		return "", err
	}
	return files.names(), nil
}
