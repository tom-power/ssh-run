package config

import (
	"bytes"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"os/user"
)

type GetConfigFrom = func() (shared.Config, error)

func GetConfigFromYaml() (shared.Config, error) {
	reader, err := getReader()
	if err != nil {
		return shared.Config{}, err
	}
	defer reader.Close()
	return GetConfigFromYamlReader(reader)
}

const configPathRelative = "/.config/ssh-run/config.yaml"

func getReader() (*os.File, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	return os.Open(user.HomeDir + configPathRelative)
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
