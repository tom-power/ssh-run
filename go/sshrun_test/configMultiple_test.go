package sshrun

import (
	"testing"
	"testing/fstest"

	"github.com/tom-power/ssh-run/sshrun/config"
)

func Test_fileSysConfigMultiple(t *testing.T) {

	t.Run("get config multiple fs", func(t *testing.T) {
		var testFsConfig = fstest.MapFS{
			"config/ssh-run.config.yml":       {Data: []byte(configYamlGeneralText)},
			"config/ssh-run.config.local.yml": {Data: []byte(configYamlLocalText)},
		}

		sys := config.ConfigFs{
			Fsys:      testFsConfig,
			ConfigDir: "config",
		}

		actual, _ := sys.GetConfig()

		assertConfigEqual(t, actual, expectedConfigFromYaml)
	})
}

var configYamlGeneralText = `
hosts:
  - ip: 192.0.2.2
    user: testUser2
    name: testName2
    port: 23
    portTunnel: 1081
  - ip: 192.0.2.3
    user: testUser3
    name: testName3
    port: 24`

var configYamlLocalText = `
includeSshConfigHosts: true
localhostIs: testName1
hosts:
  - ip: 192.0.2.1
    user: testUser1
    name: testName1
    port: 22
`
