package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
)

type GetConfigFrom = func() (shared.Config, error)

func GetConfigFromYaml() (shared.Config, error) {
	bytes, err := getConfigYamlBytes()
	if err != nil {
		return shared.Config{}, err
	}
	yaml, err := GetConfigFromYamlBytes(bytes)
	if err != nil {
		return shared.Config{}, err
	}
	return yaml, nil
}

func GetConfigFromYamlBytes(configBytes []byte) (shared.Config, error) {
	config := shared.Config{}
	err := yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

const configPathRelative = "/.config/ssh-run/config.yaml"

func getConfigYamlBytes() ([]byte, error) {
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
