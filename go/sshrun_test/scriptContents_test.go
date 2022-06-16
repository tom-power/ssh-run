package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/script"
	"testing"
)

func Test_scriptContents(t *testing.T) {
	t.Run("can get script contents 2", func(t *testing.T) {
		sys := script.FileSys{Fsys: testFs}

		actual, err := sys.Contents(testHost, "test")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "hello"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
