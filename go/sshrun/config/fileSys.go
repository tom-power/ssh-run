package config

import (
	"io/fs"
)

type ConfigFs struct {
	Fsys      fs.FS
	ConfigDir string
	SshPath   string
}
