package sshrun

import (
	"testing"
)

func Test_scripts(t *testing.T) {
	t.Run("can get scripts", func(t *testing.T) {
		actual, err := testConf.Scripts(testFs, testHost)

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "common sharedTest testSubDir test testType"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
