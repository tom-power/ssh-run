package domain

import "io/fs"

func pathCommon(fsys fs.FS, scriptName string) (string, error) {
	return firstFileInDir(fsys, commonDir()+"/", scriptName)
}
