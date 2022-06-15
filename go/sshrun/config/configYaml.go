package config

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"gopkg.in/yaml.v2"
	"io/fs"
)

func (fsys FileSys) getConfigFromYaml() (shared.Config, error) {
	file, err := fs.ReadFile(fsys.Fsys, fsys.ConfigPath)
	if err != nil {
		return shared.Config{}, err
	}
	yaml, err := getConfigFromYamlBytes(file)
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
