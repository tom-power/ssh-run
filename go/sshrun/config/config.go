package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
)

type GetConfig = func() (shared.Config, error)

var GetConfigFromFileSystem = func() (shared.Config, error) {
	bytes, err := GetConfigYamlBytes()
	if err != nil {
		return shared.Config{}, err
	}
	return GetConfigFromYaml(bytes)
}
