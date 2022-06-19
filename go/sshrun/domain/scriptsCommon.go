package domain

import (
	"io/fs"
)

func (Config) scriptsCommon(fsys fs.FS) (string, error) {
	files, err := fs.ReadDir(fsys, commonDir())
	if err != nil {
		return "", err
	}
	return Files{files}.names(), nil
}
