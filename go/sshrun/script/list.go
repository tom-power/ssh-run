package script

import (
	"github.com/tom-power/ssh-run/sshrun/domain"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"strings"
)

func (fsys FileSys) List(host domain.Host) (string, error) {
	commonScripts, _ := fsys.listCommon()
	sharedScripts, _ := fsys.listShared(host)
	hostLocalScripts, _ := fsys.listHostLocal(host)
	hostRemoteScripts := ""
	if fsys.Config.CheckRemoteForScripts {
		hostRemoteScripts, _ = listHostRemote(host)
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
