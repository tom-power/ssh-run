package domain

import (
	"io/fs"
)

type FileSys struct {
	Fsys   fs.FS
	Config Config
}
