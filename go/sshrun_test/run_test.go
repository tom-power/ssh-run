package sshrun_test

import (
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/domain"
)

var scriptsDir = ".config/ssh-run/scripts/"

var testFs = fstest.MapFS{
	scriptsDir + "host/test":                {Mode: fs.ModeDir},
	scriptsDir + "host/test/local.local.sh": {Data: []byte("command")},
	scriptsDir + "host/test/remote.sh":      {Data: []byte("command")},
	scriptsDir + "host/test/pty.pty.sh":     {Data: []byte("pty command")},
	scriptsDir + "host/test/x11.x11.sh":     {Data: []byte("x11 command")},
}

var testHosts = []domain.Host{
	{
		Name:       "test",
		User:       "user",
		Ip:         "192.0.2.1",
		Port:       "22",
		PortTunnel: "1081",
	},
	{
		Name: "test1",
		User: "user1",
		Ip:   "192.0.2.1",
		Port: "22",
	},
}

var testConfig = domain.Config{
	Hosts: testHosts,
}

var runner = sshrun.Runner{Config: testConfig, Fsys: testFs}

func Test_runFs(t *testing.T) {
	t.Run("can list hosts", func(t *testing.T) {
		actual, _ := runner.Run("", "", sshrun.RunFlags{Hosts: true})
		expected := "localhost test test1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can explain host", func(t *testing.T) {
		actual, _ := runner.Run("test", "", sshrun.RunFlags{Explain: true})
		expected := `{"Ip":"192.0.2.1","User":"user","Name":"test","Port":"22","PortTunnel":"1081","CheckRemote":false}`
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can ssh", func(t *testing.T) {
		actual, _ := runner.Run("test", "", sshrun.RunFlags{Ssh: true})
		expected := "ssh -p 22 user@192.0.2.1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can ssh on host when no script or flags", func(t *testing.T) {
		actual, _ := runner.Run("test", "", sshrun.RunFlags{})
		expected := "ssh -p 22 user@192.0.2.1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can list scripts", func(t *testing.T) {
		actual, _ := runner.Run("test", "", sshrun.RunFlags{Scripts: true})
		expected := "local pty remote x11"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run script", func(t *testing.T) {
		actual, _ := runner.Run("test", "remote", sshrun.RunFlags{})
		expected := "ssh -p 22 user@192.0.2.1 \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can explain script", func(t *testing.T) {
		actual, _ := runner.Run("test", "remote", sshrun.RunFlags{Explain: true})
		expected := "ssh -p 22 user@192.0.2.1 \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run script with pty", func(t *testing.T) {
		actual, _ := runner.Run("test", "pty", sshrun.RunFlags{})
		expected := "ssh -p 22 user@192.0.2.1 -t \"pty command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run script x11", func(t *testing.T) {
		actual, _ := runner.Run("test", "x11", sshrun.RunFlags{})

		expected := "ssh -p 22 user@192.0.2.1 -Y \"x11 command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run script local", func(t *testing.T) {
		actual, _ := runner.Run("test", "local", sshrun.RunFlags{})
		if actual != "command" {
			t.Errorf("'%v' should equal '%v'", actual, "command")
		}
	})

	t.Run("can get help on empty", func(t *testing.T) {
		expectedContains := "Usage"
		actual, _ := runner.Run("", "", sshrun.RunFlags{})
		if !strings.Contains(actual, expectedContains) {
			t.Errorf("'%v' should start with '%v'", actual, expectedContains)
		}
	})

	t.Run("can get help with option", func(t *testing.T) {
		expectedContains := "Usage"
		actual, _ := runner.Run("", "", sshrun.RunFlags{Help: true})
		if !strings.Contains(actual, expectedContains) {
			t.Errorf("'%v' should start with '%v'", actual, expectedContains)
		}
	})

	t.Run("localhost should be configured or delegated", func(t *testing.T) {
		_, err := runner.Run("localhost", "remote", sshrun.RunFlags{})

		if !strings.Contains(err.Error(), "localhost should be configured or delegated") {
			t.Errorf("error should contain 'localhost should be configured or delegated'")
		}
	})

	t.Run("can delegate using LocalhostIs", func(t *testing.T) {
		testConfig.LocalhostIs = "test"
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("localhost", "remote", sshrun.RunFlags{})
		expected := "ssh -p 22 user@192.0.2.1 \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

}
