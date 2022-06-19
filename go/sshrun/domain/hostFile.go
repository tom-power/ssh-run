package domain

import "io/fs"

func (host Host) Dir() string {
	return scriptsPath + "host/" + host.Name
}

func (host Host) Files(fsys fs.FS) ([]fs.DirEntry, error) {
	return fs.ReadDir(fsys, host.Dir())
}

func (host Host) FilePath(fsys fs.FS, scriptName string) (string, error) {
	return firstPathInDir(fsys, host.Dir()+"/", scriptName)
}

func (host Host) FilePathInSubDir(fsys fs.FS, dirName string, scriptName string) (string, error) {
	return firstPathInDir(fsys, host.Dir()+"/"+dirName+"/", scriptName)
}
