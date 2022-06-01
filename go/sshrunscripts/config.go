package sshrunscripts

import (
	"io/ioutil"
	"os/user"
)

const configPathRelative = "/.config/ssh-run-scripts/config.yaml"

func ReadConfig() ([]byte, error) {
	usr, err := user.Current()
	if err != nil {
		return []byte{}, err
	}
	path := usr.HomeDir + configPathRelative
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return fileBytes, nil
}
