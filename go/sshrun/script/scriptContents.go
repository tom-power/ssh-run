package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/ioutil"
)

type GetScriptContents = func(host shared.Host, scriptPath string) (string, error)

var GetScriptContentsFromHost = func(host shared.Host, scriptPath string) (string, error) {
	script, err := getScriptContentsFromHostLocal(scriptPath)
	if host.CheckForScripts {
		script, err = getScriptContentsFromHostRemote(host, scriptPath)
	}
	if err != nil {
		return "", err
	}
	return script, nil
}

func getScriptContentsFromHostLocal(scriptPath string) (string, error) {
	dat, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func getScriptContentsFromHostRemote(host shared.Host, scriptPath string) (string, error) {
	return runCommandOn(host, "cat "+scriptPath)
}
