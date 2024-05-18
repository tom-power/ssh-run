package sshrun

import (
	"testing"
)

func Test_scripts(t *testing.T) {
	t.Run("can get scripts", func(t *testing.T) {
		actual, err := testHost.Scripts(testFs)

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "common utilsTest testSubDir test testType"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
