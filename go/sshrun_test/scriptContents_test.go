package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/script"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"testing"
)

func Test_scriptContents(t *testing.T) {
	t.Run("can get script contents", func(t *testing.T) {
		scriptPath := scriptsDir + "host/testHost/test.sh"
		actual, err := script.GetScriptContentsFromHost(testFs)(shared.Host{}, scriptPath, shared.Config{})
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "hello"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
