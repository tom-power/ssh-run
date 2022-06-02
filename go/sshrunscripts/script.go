package sshrunscripts

import (
	"io/ioutil"
)

type GetScript = func(scriptPath string) (string, error)

var GetScriptFromConf = func(scriptPath string) (string, error) {
	dat, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
