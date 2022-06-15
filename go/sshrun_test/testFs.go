package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"testing/fstest"
)

var scriptsDir = ".config/ssh-run/scripts/"

var testFs = fstest.MapFS{
	scriptsDir + "common/common.sh":                   {},
	scriptsDir + "host/testHost":                      {Mode: fs.ModeDir},
	scriptsDir + "host/testHost/test.sh":              {},
	scriptsDir + "host/testHost/subDir":               {Mode: fs.ModeDir},
	scriptsDir + "host/testHost/subDir/testSubDir.sh": {},
	scriptsDir + "host/testHost/stuff":                {Mode: fs.ModeDir},
	scriptsDir + "shared/stuff":                       {Mode: fs.ModeDir},
	scriptsDir + "shared/stuff/sharedTest.sh":         {},
}

var scriptPathHost = shared.Host{Name: "testHost"}

var scriptPathConf = shared.Config{CheckRemoteForScripts: false}
