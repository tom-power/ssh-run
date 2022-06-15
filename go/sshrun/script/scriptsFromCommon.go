package script

import (
	"io/fs"
)

func getScriptsFromCommon(fsys fs.FS) (string, error) {
	files, err := fs.ReadDir(fsys, commonDir())
	if err != nil {
		return "", err
	}
	return filesToFileNames(files), nil
}
