package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"strings"
)

type GetScripts = func(host shared.Host, config shared.Config) (string, error)

func GetScriptsFromConf(fs fs.FS) GetScripts {
	return func(host shared.Host, config shared.Config) (string, error) {
		commonScripts, _ := getScriptsFromCommon(fs)
		sharedScripts, _ := getScriptsFromShared(host, fs)
		hostScripts, _ := getScriptsFromHost(host, config, fs)
		var out = ""
		out += commonScripts
		out += " " + sharedScripts
		out += " " + hostScripts
		return removeSh(removeCommandTypes(out)), nil
	}
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

func getScriptsFromCommon(fsys fs.FS) (string, error) {
	files, err := fs.ReadDir(fsys, commonDir())
	if err != nil {
		return "", err
	}
	return filesToFileNames(files), nil
}

func getScriptsFromShared(host shared.Host, fsys fs.FS) (string, error) {
	var files []fs.DirEntry
	hostFiles, _ := fs.ReadDir(fsys, hostDir(host.Name))
	for _, hostFile := range hostFiles {
		if hostFile.IsDir() {
			sharedDir := scriptsPath + "shared/" + hostFile.Name()
			if fileExistsFs(fsys)(sharedDir) {
				err := fs.WalkDir(fsys, sharedDir, appendFiles(&files))
				if err != nil {
					return "", err
				}
			}
		}
	}
	return filesToFileNames(files), nil
}

func filesToFileNames(files []fs.DirEntry) string {
	filesToFileName := func(dir fs.DirEntry) string { return dir.Name() }
	return strings.Join(shared.Map(files, filesToFileName), " ")
}

func appendFiles(files *[]fs.DirEntry) func(string, fs.DirEntry, error) error {
	return func(path string, info fs.DirEntry, err error) error {
		if info != nil && !info.IsDir() {
			*files = append(*files, info)
		}
		return err
	}
}
