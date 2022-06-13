package script

import (
	"os"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isNotEmpty(path string) bool {
	return path != ""
}
