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
	scriptsDir + "host/test/withArgs.sh":    {Data: []byte("command $1 $2")},
	scriptsDir + "host/test/pty.pty.sh":     {Data: []byte("pty command")},
	scriptsDir + "host/test/x11.x11.sh":     {Data: []byte("x11 command")},
}

var testHosts = []domain.Host{
	{
		Name:       "test",
		User:       "user",
		Host:       "192.0.2.1",
		Port:       "22",
		PortTunnel: "1081",
	},
	{
		Name: "test1",
		User: "user1",
		Host: "192.0.2.1",
		Port: "22",
	},
}

var testConfig = domain.Config{
	Hosts: testHosts,
}

func Test_runFs(t *testing.T) {
	t.Run("can ssh", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "ssh", []string{})
		expected := "ssh -p 22 user@192.0.2.1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can ssh on no command", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "", []string{})
		expected := "ssh -p 22 user@192.0.2.1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run local", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "local", []string{})
		if actual != "command" {
			t.Errorf("'%v' should equal '%v'", actual, "command")
		}
	})

	t.Run("can run ssh", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "remote", []string{})
		expected := "ssh -p 22 user@192.0.2.1 \"command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh with args", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "withArgs", []string{"arg1", "arg2"})
		expected := "ssh -p 22 user@192.0.2.1 \"command arg1 arg2\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh with pty", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "pty", []string{})
		expected := "ssh -p 22 user@192.0.2.1 -t \"pty command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can run ssh x11", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "x11", []string{})

		expected := "ssh -p 22 user@192.0.2.1 -Y \"x11 command\""
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can list scripts", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "scripts", []string{})
		expected := "echo local pty remote withArgs x11"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can list hosts", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("hosts", "", []string{})
		expected := "echo test test1"
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can explain host", func(t *testing.T) {
		actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run("test", "explain", []string{})
		expected := `Host{Host:"192.0.2.1", User:"user", Name:"test", Port:"22", PortTunnel:"1081", CheckRemote:false}`
		if actual != expected {
			t.Errorf("'%v' should equal '%v'", actual, expected)
		}
	})

	t.Run("can get help", func(t *testing.T) {
		helps := []string{"", "--help", "-h"}
		for _, help := range helps {
			expectedContains := "Usage"
			actual, _ := sshrun.Runner{Config: testConfig, Fsys: testFs}.Run(help, "", []string{})
			if !strings.Contains(actual, expectedContains) {
				t.Errorf("'%v' should start with '%v'", actual, expectedContains)
			}
		}
	})

}
