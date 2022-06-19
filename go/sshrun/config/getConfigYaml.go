package config

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"gopkg.in/yaml.v2"
	"io/fs"
)

func (c ConfigFs) getConfigFromYaml() (domain.Config, error) {
	file, err := fs.ReadFile(c.Fsys, c.ConfigPath)
	if err != nil {
		return domain.Config{}, err
	}
	yaml, err := getConfigFromYamlBytes(file)
	if err != nil {
		return domain.Config{}, err
	}
	return yaml, nil
}

func getConfigFromYamlBytes(configBytes []byte) (domain.Config, error) {
	config := domain.Config{}
	err := yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
