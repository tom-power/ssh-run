package domain

import (
	"io/fs"
	"strings"

	u "github.com/rjNemo/underscore"
)

func (h Host) Scripts(fsys fs.FS) (string, error) {
	commonScripts, _ := scriptsCommon(fsys)
	sharedScripts, _ := h.scriptsShared(fsys)
	hostLocalScripts, _ := h.scriptsLocal(fsys)
	hostRemoteScripts := ""
	if h.CheckRemote {
		hostRemoteScripts, _ = h.scriptsRemote()
	}
	out := strings.Join([]string{commonScripts, sharedScripts, hostLocalScripts, hostRemoteScripts}, " ")
	return Cleaner{out}.dropNotSh().nameOnly().Scripts, nil
}

type Cleaner struct {
	Scripts string
}

func (cleaner Cleaner) nameOnly() Cleaner {
	split := strings.Split(cleaner.Scripts, " ")
	mapped := u.Map(split, nameOnly)
	join := strings.Join(mapped, " ")
	return Cleaner{join}
}

var nameOnly = func(script string) string { return strings.Split(script, ".")[0] }

func (cleaner Cleaner) dropNotSh() Cleaner {
	split := strings.Split(cleaner.Scripts, " ")
	filter := u.Filter(split, noSh)
	join := strings.Join(filter, " ")
	return Cleaner{join}
}

var noSh = func(script string) bool { return strings.HasSuffix(script, ".sh") }

type Files struct {
	Files []fs.DirEntry
}

func (files Files) names() string {
	fileToFileName := func(dir fs.DirEntry) string { return dir.Name() }
	return strings.Join(u.Map(files.Files, fileToFileName), " ")
}
