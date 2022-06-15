package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/script"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"testing"
	"testing/fstest"
)

func Test_scriptPath(t *testing.T) {
	t.Run("can get common script path from fs", func(t *testing.T) {
		actual, err := script.GetScriptPathFromConf(testFs)(scriptPathHost, "common", scriptPathConf)
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "common/common.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get host script path from fs", func(t *testing.T) {
		actual, err := script.GetScriptPathFromConf(testFs)(scriptPathHost, "test", scriptPathConf)
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "host/testHost/test.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get subdir host script path from fs", func(t *testing.T) {
		actual, err := script.GetScriptPathFromConf(testFs)(scriptPathHost, "subdirTest", scriptPathConf)
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "host/testHost/subdir/subdirTest.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get shared script path from fs", func(t *testing.T) {
		actual, err := script.GetScriptPathFromConf(testFs)(scriptPathHost, "testShared", scriptPathConf)
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "shared/stuff/testShared.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}

var scriptsDir = ".config/ssh-run/scripts/"

var testFs = fstest.MapFS{
	scriptsDir + "common/common.sh":                   {},
	scriptsDir + "host/testHost":                      {Mode: fs.ModeDir},
	scriptsDir + "host/testHost/test.sh":              {},
	scriptsDir + "host/testHost/subdir":               {Mode: fs.ModeDir},
	scriptsDir + "host/testHost/subdir/subdirTest.sh": {},
	scriptsDir + "host/testHost/stuff":                {Mode: fs.ModeDir},
	scriptsDir + "shared/stuff":                       {Mode: fs.ModeDir},
	scriptsDir + "shared/stuff/testShared.sh":         {},
}

var scriptPathHost = shared.Host{Name: "testHost"}

var scriptPathConf = shared.Config{CheckRemoteForScripts: false}
