package domain

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"strings"
)

func (host Host) Scripts(fsys fs.FS) (string, error) {
	commonScripts, _ := scriptsCommon(fsys)
	sharedScripts, _ := host.scriptsShared(fsys)
	hostLocalScripts, _ := host.scriptsLocal(fsys)
	hostRemoteScripts := ""
	if host.CheckRemote {
		hostRemoteScripts, _ = host.scriptsRemote()
	}
	out := strings.Join([]string{commonScripts, sharedScripts, hostLocalScripts, hostRemoteScripts}, " ")
	return Cleaner{out}.dropNotSh().nameOnly().Scripts, nil
}

type Cleaner struct {
	Scripts string
}

func (cleaner Cleaner) nameOnly() Cleaner {
	split := strings.Split(cleaner.Scripts, " ")
	mapped := shared.Map(split, nameOnly)
	join := strings.Join(mapped, " ")
	return Cleaner{join}
}

var nameOnly = func(script string) string { return strings.Split(script, ".")[0] }

func (cleaner Cleaner) dropNotSh() Cleaner {
	split := strings.Split(cleaner.Scripts, " ")
	filter := shared.Filter(split, noSh)
	join := strings.Join(filter, " ")
	return Cleaner{join}
}

var noSh = func(script string) bool { return strings.HasSuffix(script, ".sh") }

type Files struct {
	Files []fs.DirEntry
}

func (files Files) filter(predicate func(fs.DirEntry) bool) Files {
	return Files{shared.Filter(files.Files, predicate)}
}

func (files Files) names() string {
	fileToFileName := func(dir fs.DirEntry) string { return dir.Name() }
	return strings.Join(shared.Map(files.Files, fileToFileName), " ")
}
