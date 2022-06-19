package domain

import (
	"io/fs"
)

func (h Host) Contents(fsys fs.FS, scriptName string) (string, error) {
	path, err := h.Path(fsys, scriptName)
	if err != nil {
		return "", err
	}
	script, err := contentsFromHostLocal(fsys, path)
	if h.CheckRemote {
		script, err = contentsFromHostRemote(h, path)
	}
	if err != nil {
		return "", err
	}
	return script, nil
}

func contentsFromHostLocal(fsys fs.FS, scriptPath string) (string, error) {
	dat, err := fs.ReadFile(fsys, scriptPath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func contentsFromHostRemote(host Host, scriptPath string) (string, error) {
	return runCommandOn(host, "cat "+scriptPath)
}
