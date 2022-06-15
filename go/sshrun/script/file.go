package script

import (
	"io/fs"
	"os"
)

func fileExistsFs(fs fs.FS) func(string) bool {
	return func(path string) bool {
		_, err := fs.Open(path)
		return err == nil
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isNotEmpty(path string) bool {
	return path != ""
}
