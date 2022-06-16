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

func (fsys FileSys) fileExists() func(string) bool {
	return func(path string) bool {
		_, err := fsys.Fsys.Open(path)
		return err == nil
	}
}

func (fsys FileSys) firstFileInDir(dir string, name string) (string, error) {
	matches, err := fs.Glob(fsys.Fsys, dir+name+".*")
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
