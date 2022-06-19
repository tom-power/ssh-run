package domain

import "io/fs"

func pathCommon(fsys fs.FS, scriptName string) (string, error) {
	return firstPathInDir(fsys, commonDir()+"/", scriptName)
}
