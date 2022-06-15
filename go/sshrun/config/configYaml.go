package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"gopkg.in/yaml.v2"
	"io/fs"
)

type GetConfig = func() (shared.Config, error)

func GetConfigFromYaml(path string, fs fs.FS) (shared.Config, error) {
	file, err := fs.Open(path)
	if err != nil {
		return shared.Config{}, err
	}
	bytes, err := getBytes(file)
	if err != nil {
		return shared.Config{}, err
	}
	yaml, err := getConfigFromYamlBytes(bytes)
	if err != nil {
		return shared.Config{}, err
	}
	return yaml, nil
}

func getConfigFromYamlBytes(configBytes []byte) (shared.Config, error) {
	config := shared.Config{}
	err := yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
