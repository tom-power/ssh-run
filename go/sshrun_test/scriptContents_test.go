package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/script"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
)

func Test_scriptContents(t *testing.T) {
	t.Run("can get script contents", func(t *testing.T) {
		sys := script.FileSys{Fsys: testFs}
		scriptPath := scriptsDir + "host/testHost/test.sh"

		actual, err := sys.Contents(shared.Host{}, scriptPath)

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "hello"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
