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
	scriptsDir + "host/testHost/rubbish":              {Mode: fs.ModeDir},
	scriptsDir + "host/testHost/test.sh":              {Data: []byte("hello")},
	scriptsDir + "host/testHost/subDir":               {Mode: fs.ModeDir},
	scriptsDir + "host/testHost/subDir/testSubDir.sh": {},
	scriptsDir + "host/testHost/stuff":                {Mode: fs.ModeDir},
	scriptsDir + "host/testHost/stuff/.keep":          {Mode: fs.ModeDir},
	scriptsDir + "shared/stuff":                       {Mode: fs.ModeDir},
	scriptsDir + "shared/stuff/sharedTest.sh":         {},
	scriptsDir + "shared/stuff/rubbish":               {},
}

var scriptPathHost = shared.Host{Name: "testHost"}

var scriptPathConf = shared.Config{CheckRemoteForScripts: false}
