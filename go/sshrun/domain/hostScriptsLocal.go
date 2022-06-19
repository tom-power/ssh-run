package domain

import (
	"io/fs"
)

func (h Host) scriptsLocal(fsys fs.FS) (string, error) {
	var files = Files{[]fs.DirEntry{}}
	err := fs.WalkDir(fsys, h.Dir(), appendFiles(&files.Files))
	if err != nil {
		return "", err
	}
	return files.names(), nil
}
