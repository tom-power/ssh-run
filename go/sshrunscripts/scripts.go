package sshrunscripts

import (
	"strings"

	"io/ioutil"
	"os"
	"path/filepath"
)

type GetScripts = func(host Host) (string, error)

var getHostScripts = func(host Host) (string, error) {
	return scriptsFromHostLocal(host)
	// return scriptsFromHostRemote(host)
}

var GetScriptsAll = func(host Host) (string, error) {
	commonScripts, err := getScriptsFromCommon()
	if err != nil {
		return "", err
	}
	shared, err := getScriptsFromShared(host)
	if err != nil {
		return "", err
	}
	hostScripts, err := getHostScripts(host)
	if err != nil {
		return "", err
	}
	var out = ""
	out += commonScripts
	out += " " + shared
	out += " " + hostScripts
	return removeCommandTypes(out), nil
}

func removeCommandTypes(scripts string) string {
	scripts = strings.ReplaceAll(scripts, ".local", "")
	scripts = strings.ReplaceAll(scripts, ".sudo", "")
	scripts = strings.ReplaceAll(scripts, ".x11", "")
	scripts = strings.ReplaceAll(scripts, ".ssh", "")
	return scripts
}

func getScriptsFromCommon() (string, error) {
	files, err := ioutil.ReadDir(commonDir())
	if err != nil {
		return "", err
	}
	return filesToFileNames(files), nil
}

func getScriptsFromShared(host Host) (string, error) {
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
	filesToFileName := func(file os.FileInfo) string { return fileName(file) }
	return strings.Join(Map(files, filesToFileName), " ")
}

func fileName(file os.FileInfo) string {
	return strings.ReplaceAll(file.Name(), ".sh", "")
}

func appendFiles(files *[]os.FileInfo) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			*files = append(*files, info)
		}
		return err
	}
}
