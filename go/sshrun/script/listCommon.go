package script

import (
	"io/fs"
)

func (fsys FileSys) listCommon() (string, error) {
	files, err := fs.ReadDir(fsys.Fsys, commonDir())
	if err != nil {
		return "", err
	}
	return Files{files}.names(), nil
}
