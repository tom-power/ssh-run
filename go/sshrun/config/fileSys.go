package config

import (
	"io/fs"
)

type ConfigFs struct {
	Fsys       fs.FS
	ConfigPath string
	SshPath    string
}
