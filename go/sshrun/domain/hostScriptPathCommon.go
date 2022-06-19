package domain

import "io/fs"

func pathCommon(fsys fs.FS, scriptName string) (string, error) {
	return pathInDir(fsys, commonDir()+"/", scriptName)
}
