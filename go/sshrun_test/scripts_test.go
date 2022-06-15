package sshrun

import (
	"github.com/tom-power/ssh-run/sshrun/script"
	"testing"
)

func Test_scripts(t *testing.T) {
	t.Run("can get scripts", func(t *testing.T) {
		actual, err := script.GetScriptsFromConf(testFs)(scriptPathHost, scriptPathConf)
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "common sharedTest testSubDir test"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
