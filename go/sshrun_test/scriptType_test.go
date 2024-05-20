package sshrun

import (
	"testing"

	"github.com/tom-power/ssh-run/sshrun/domain"
)

func Test_scriptType(t *testing.T) {
	t.Run("can get script type from fs", func(t *testing.T) {
		actual, err := testHost.ScriptTypeFor(testFs, "testType")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := domain.Pty
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
