package domain

import (
	"errors"
	"io/fs"
)

func fileExists(fsys fs.FS) func(string) bool {
	return func(path string) bool {
		_, err := fsys.Open(path)
		return err == nil
	}
}

func firstPathInDir(fsys fs.FS, dir string, name string) (string, error) {
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
