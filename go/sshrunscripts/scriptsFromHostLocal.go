package sshrunscripts

import (
	"os"
	"path/filepath"
)

func scriptsFromHostLocal(host Host) (string, error) {
	var files []os.FileInfo
	err := filepath.Walk(hostsDir(host.Name), appendFiles(&files))
	return filesToFileNames(filterKeep(files)), err
}

func filterKeep(files []os.FileInfo) []os.FileInfo {
	var noKeep = func(file os.FileInfo) bool { return file.Name() != ".keep" }
	return Filter(files, noKeep)
}
