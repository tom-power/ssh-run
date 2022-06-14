package config

import (
	"os"
	"os/user"
)

func getConfigReader(path string) (*os.File, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	reader, err := os.Open(user.HomeDir + "/" + path)
	if err != nil {
		return nil, err
	}
	return reader, nil
}
