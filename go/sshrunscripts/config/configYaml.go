package config

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
)

func GetConfigFromYaml(configBytes []byte) (shared.Config, error) {
	config := shared.Config{}
	err := yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

const configPathRelative = "/.config/ssh-run-scripts/config.yaml"

func GetConfigYamlBytes() ([]byte, error) {
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
