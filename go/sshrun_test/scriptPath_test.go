package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/script"
	"testing"
)

func Test_scriptPath(t *testing.T) {
	t.Run("can get common script path from fs", func(t *testing.T) {
		sys := script.FileSys{Fsys: testFs, Config: testConf}

		actual, err := sys.Path(testHost, "common")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "common/common.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get host script path from fs", func(t *testing.T) {
		sys := script.FileSys{Fsys: testFs, Config: testConf}

		actual, err := sys.Path(testHost, "test")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "host/testHost/test.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get subDir host script path from fs", func(t *testing.T) {
		sys := script.FileSys{Fsys: testFs, Config: testConf}

		actual, err := sys.Path(testHost, "testSubDir")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "host/testHost/subDir/testSubDir.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get shared script path from fs", func(t *testing.T) {
		sys := script.FileSys{Fsys: testFs, Config: testConf}

		actual, err := sys.Path(testHost, "sharedTest")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "shared/stuff/sharedTest.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
