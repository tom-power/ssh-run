package config

import (
	"io/fs"
	"strings"

	"dario.cat/mergo"
	"github.com/tom-power/ssh-run/sshrun/domain"
	"gopkg.in/yaml.v2"
)

func (c ConfigFs) getConfigFromYaml() (domain.Config, error) {
	dirEntries, err := fs.ReadDir(c.Fsys, c.ConfigDir)
	if err != nil {
		return domain.Config{}, err
	}

	config := domain.Config{}
	for _, dirEntry := range dirEntries {
		if dirEntry.Type().IsRegular() && strings.HasSuffix(dirEntry.Name(), ".yml") {
			err := c.updateConfigFromFile(&config, dirEntry)
			if err != nil {
				return domain.Config{}, err
			}
		}
	}

	return config, nil
}

func (c ConfigFs) updateConfigFromFile(config *domain.Config, dirEntry fs.DirEntry) error {
	file, err := fs.ReadFile(c.Fsys, c.ConfigDir+"/"+dirEntry.Name())
	if err != nil {
		return err
	}
	thisConfig, err := getConfigFromYamlBytes(file)
	if err != nil {
		return err
	}
	if err := mergo.Merge(config, thisConfig); err != nil {
		return err
	}
	return nil
}

func getConfigFromYamlBytes(configBytes []byte) (domain.Config, error) {
	config := domain.Config{}
	err := yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
