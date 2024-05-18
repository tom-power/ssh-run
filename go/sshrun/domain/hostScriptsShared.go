package domain

import (
	"io/fs"
)

func (h Host) scriptsShared(fsys fs.FS) (string, error) {
	var files = Files{[]fs.DirEntry{}}
	hostFiles, _ := fs.ReadDir(fsys, h.Dir())
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			utilsDir := scriptsPath + "utils/" + hostFile.Name()
			if fileExists(fsys)(utilsDir) {
				err := fs.WalkDir(fsys, utilsDir, appendFiles(&files.Files))
				if err != nil {
					return "", err
				}
			}
		}
	}
	return files.names(), nil
}

func appendFiles(files *[]fs.DirEntry) func(string, fs.DirEntry, error) error {
	return func(path string, info fs.DirEntry, err error) error {
		if info != nil && !info.IsDir() {
			*files = append(*files, info)
		}
		return err
	}
}
