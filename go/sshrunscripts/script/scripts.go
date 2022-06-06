package script

import (
	"github.com/tom-power/ssh-run-scripts/sshrunscripts/shared"
	"strings"

	"io/ioutil"
	"os"
	"path/filepath"
)

type GetScripts = func(host shared.Host) (string, error)

var GetScriptsAll = func(host shared.Host) (string, error) {
	commonScripts, _ := getScriptsFromCommon()
	shared, _ := getScriptsFromShared(host)
	hostScripts, _ := getScriptsFromHost(host)
	var out = ""
	out += commonScripts
	out += " " + shared
	out += " " + hostScripts
	return removeSh(removeCommandTypes(out)), nil
}

func removeCommandTypes(scripts string) string {
	scripts = strings.ReplaceAll(scripts, ".local", "")
	scripts = strings.ReplaceAll(scripts, ".pty", "")
	scripts = strings.ReplaceAll(scripts, ".x11", "")
	scripts = strings.ReplaceAll(scripts, ".ssh", "")
	return scripts
}

func removeSh(scripts string) string {
	return strings.ReplaceAll(scripts, ".sh", "")
}

func getScriptsFromCommon() (string, error) {
	files, err := ioutil.ReadDir(commonDir())
	if err != nil {
		return "", err
	}
	return filesToFileNames(files), nil
}

func getScriptsFromShared(host shared.Host) (string, error) {
	var files []os.FileInfo
	hostFiles, _ := ioutil.ReadDir(hostDir(host.Name, homeDir()))
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			sharedDir := scriptsDir(homeDir()) + "shared/" + hostFile.Name()
			if fileExists(sharedDir) {
				err := filepath.Walk(sharedDir, appendFiles(&files))
				if err != nil {
					return "", err
				}
			}
		}
	}
	return filesToFileNames(files), nil
}

func filesToFileNames(files []os.FileInfo) string {
	filesToFileName := func(file os.FileInfo) string { return file.Name() }
	return strings.Join(shared.Map(files, filesToFileName), " ")
}

func appendFiles(files *[]os.FileInfo) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			*files = append(*files, info)
		}
		return err
	}
}
