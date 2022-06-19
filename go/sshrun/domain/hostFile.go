package domain

import "io/fs"

func (h Host) Dir() string {
	return scriptsPath + "host/" + h.Name
}

func (h Host) Files(fsys fs.FS) ([]fs.DirEntry, error) {
	return fs.ReadDir(fsys, h.Dir())
}

func (h Host) FilePath(fsys fs.FS, scriptName string) (string, error) {
	return firstPathInDir(fsys, h.Dir()+"/", scriptName)
}

func (h Host) FilePathInSubDir(fsys fs.FS, dirName string, scriptName string) (string, error) {
	return firstPathInDir(fsys, h.Dir()+"/"+dirName+"/", scriptName)
}
