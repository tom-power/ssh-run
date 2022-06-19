package domain

import (
	"io/fs"
)

func (Config) scriptsHostLocal(fsys fs.FS, host Host) (string, error) {
	var files = Files{[]fs.DirEntry{}}
	err := fs.WalkDir(fsys, hostDir(host.Name), appendFiles(&files.Files))
	if err != nil {
		return "", err
	}
	return files.names(), nil
}
