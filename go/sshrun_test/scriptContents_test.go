package sshrun

import (
	"testing"
)

func Test_scriptContents(t *testing.T) {
	t.Run("can get script contents 2", func(t *testing.T) {
		actual, err := testHost.Contents(testFs, "test")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "hello"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
