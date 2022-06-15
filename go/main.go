package main

import (
	"fmt"
	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/script"
	"github.com/tom-power/ssh-run/sshrun/shared"
	"io/fs"
	"log"
	"os"
)

func main() {
	hostName := shared.SafeSlice(os.Args, 1, "")
	scriptName := shared.SafeSlice(os.Args, 2, "")
	args := []string{""}
	if len(os.Args) > 3 {
		args = os.Args[3:]
	}
	dir, err := homeDirFs()
	if err != nil {
		log.Fatal(err)
	}
	config, err := getConfig(dir)
	if err != nil {
		log.Fatal(err)
	}
	sshRun, err := getRun(config)(hostName, scriptName, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sshRun)
}

func getRun(config shared.Config) sshrun.Run {
	return sshrun.GetRun(
		sshrun.GetHostFromConfig(config),
		script.GetScriptPathFromConf,
		script.GetScriptContentsFromHost,
		sshrun.GetCommandSsh,
		script.GetScriptsAll,
		sshrun.GetHostsFromConfig(config),
		config,
	)
}

func homeDirFs() (fs.FS, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return os.DirFS(homeDir), nil
}

func getConfig(homeDirFs fs.FS) (shared.Config, error) {
	var getHostsFromSshConfig = func() ([]shared.Host, error) {
		return config.GetHostsFromSshConfig(".ssh/config", homeDirFs)
	}
	var getConfigFromYaml = func() (shared.Config, error) {
		return config.GetConfigFromYaml(".config/ssh-run/config.yaml", homeDirFs)
	}
	return config.GetConfigUsing(
		getHostsFromSshConfig,
		getConfigFromYaml,
	)
}
