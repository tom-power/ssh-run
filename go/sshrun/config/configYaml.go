package config

import (
	"bytes"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"gopkg.in/yaml.v2"
	"io"
)

type GetConfigFrom = func() (shared.Config, error)

func GetConfigFromYaml() (shared.Config, error) {
	reader, err := getConfigReader("/.config/ssh-run/config.yaml")
	if err != nil {
		return shared.Config{}, err
	}
	return GetConfigFromYamlReader(reader)
}

func GetConfigFromYamlReader(reader io.Reader) (shared.Config, error) {
	yaml, err := getConfigFromYamlBytes(getBytes(reader))
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

func getBytes(reader io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.Bytes()
}
