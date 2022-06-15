package config

import (
	"io/fs"
)

type FileSys struct {
	Fsys       fs.FS
	ConfigPath string
	SshPath    string
}
