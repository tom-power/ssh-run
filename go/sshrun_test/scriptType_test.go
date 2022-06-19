package sshrun

import (
	"testing"
)

func Test_scriptType(t *testing.T) {
	t.Run("can get script type from fs", func(t *testing.T) {
		actual, err := testHost.Type(testFs, "testType")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "pty"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
