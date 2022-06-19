package domain

import (
	"io/fs"
)

func (config Config) Contents(fsys fs.FS, host Host, scriptName string) (string, error) {
	path, err := config.Path(fsys, host, scriptName)
	if err != nil {
		return "", err
	}
	script, err := contentsFromHostLocal(fsys, path)
	if config.CheckRemoteForScripts {
		script, err = contentsFromHostRemote(host, path)
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
