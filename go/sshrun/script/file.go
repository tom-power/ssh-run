package script

import (
	"errors"
	"io/fs"
)

func fileExists(fs fs.FS) func(string) bool {
	return func(path string) bool {
		_, err := fs.Open(path)
		return err == nil
	}
}

func firstFileInDir(dir string, name string, fsys fs.FS) (string, error) {
	matches, err := fs.Glob(fsys, dir+name+".*")
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", errors.New("no match")
	}
	return matches[0], nil
}
func isNotEmpty(path string) bool {
	return path != ""
}
