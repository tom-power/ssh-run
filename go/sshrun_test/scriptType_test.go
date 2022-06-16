package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/script"
	"testing"
)

func Test_scriptType(t *testing.T) {
	t.Run("can get script type from fs", func(t *testing.T) {
		sys := script.FileSys{Fsys: testFs, Config: testConf}

		actual, err := sys.Type(testHost, "testType")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "pty"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
