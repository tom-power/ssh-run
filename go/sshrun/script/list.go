package script

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"strings"
)

func (fsys FileSys) List(host shared.Host) (string, error) {
	commonScripts, _ := fsys.listCommon()
	sharedScripts, _ := fsys.listShared(host)
	hostLocalScripts, _ := fsys.listHostLocal(host)
	hostRemoteScripts := ""
	if fsys.Config.CheckRemoteForScripts {
		hostRemoteScripts, _ = listHostRemote(host)
		hostRemoteScripts = " " + hostRemoteScripts
	}
	var out = ""
	out += commonScripts
	out += " " + sharedScripts
	out += " " + hostLocalScripts
	out += hostRemoteScripts
	return Cleaner{out}.dropNotSh().nameOnly().Scripts, nil
}

type Cleaner struct {
	Scripts string
}

func (cleaner Cleaner) nameOnly() Cleaner {
	split := strings.Split(cleaner.Scripts, " ")
	mapped := shared.Map(split, nameOnly)
	joined := strings.Join(mapped, " ")
	return Cleaner{joined}
}

var nameOnly = func(script string) string { return strings.Split(script, ".")[0] }

func (cleaner Cleaner) dropNotSh() Cleaner {
	filter := shared.Filter(strings.Split(cleaner.Scripts, " "), noSh)
	return Cleaner{strings.Join(filter, " ")}
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
