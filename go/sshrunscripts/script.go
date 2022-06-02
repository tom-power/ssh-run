package sshrunscripts

import (
	"io/ioutil"
)

type GetScript = func(scriptPath string) (string, error)

var GetScriptFromHostLocal = func(scriptPath string) (string, error) {
	dat, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

var GetScriptFromHostRemote = func(host Host) GetScript {
	return func(scriptPath string) (string, error) {
		return runCommandOn(host, "cat "+scriptPath)
	}
}
