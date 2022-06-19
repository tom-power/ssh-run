package domain

import "io/fs"

func (h Host) Dir() string {
	return scriptsPath + "host/" + h.Name
}

func (h Host) Files(fsys fs.FS) ([]fs.DirEntry, error) {
	return fs.ReadDir(fsys, h.Dir())
}
