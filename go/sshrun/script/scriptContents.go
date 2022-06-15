package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
)

type GetScriptContents = func(host shared.Host, scriptPath string, config shared.Config) (string, error)

func GetScriptContentsFromHost(fs fs.FS) GetScriptContents {
	return func(host shared.Host, scriptPath string, config shared.Config) (string, error) {
		script, err := getScriptContentsFromHostLocal(scriptPath, fs)
		if config.CheckRemoteForScripts {
			script, err = getScriptContentsFromHostRemote(host, scriptPath)
		}
		if err != nil {
			return "", err
		}
		return script, nil
	}
}

func getScriptContentsFromHostLocal(scriptPath string, fsys fs.FS) (string, error) {
	dat, err := fs.ReadFile(fsys, scriptPath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func getScriptContentsFromHostRemote(host shared.Host, scriptPath string) (string, error) {
	return runCommandOn(host, "cat "+scriptPath)
}
