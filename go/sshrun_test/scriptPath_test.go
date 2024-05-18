package sshrun

import (
	"testing"
)

func Test_scriptPath(t *testing.T) {
	t.Run("can get common script path from fs", func(t *testing.T) {
		actual, err := testHost.Path(testFs, "common")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "common/common.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get host script path from fs", func(t *testing.T) {
		actual, err := testHost.Path(testFs, "test")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "host/testHost/test.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get subDir host script path from fs", func(t *testing.T) {
		actual, err := testHost.Path(testFs, "testSubDir")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "host/testHost/subDir/testSubDir.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
	t.Run("can get utils script path from fs", func(t *testing.T) {
		actual, err := testHost.Path(testFs, "utilsTest")

		if err != nil {
			t.Errorf(err.Error())
		}
		expected := scriptsDir + "utils/stuff/utilsTest.sh"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})
}
