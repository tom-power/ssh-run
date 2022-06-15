package script

import (
	"io/fs"
)

func scriptPathFromCommon(scriptName string, fs fs.FS) (string, error) {
	return firstFileInDir(commonDir()+"/", scriptName, fs)
}
