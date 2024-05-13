package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared"
)

func main() {
	homeDirFs, err := getHomeDirFs()
	if err != nil {
		log.Fatal(err)
	}
	config, err := getConfigFs(homeDirFs).GetConfig()
	if err != nil {
		log.Fatal("config error:", err)
	}
	runner := sshrun.Runner{
		Config: config,
		Fsys:   homeDirFs,
	}

	hostName := shared.GetOr(getCommands(os.Args), 1, "")
	scriptName := shared.GetOr(getCommands(os.Args), 2, "")
	args := getArgs(os.Args)

	sshRun, err := runner.Run(hostName, scriptName, args, []string{""})
	if err != nil {
		log.Fatal("runner error:", err)
	}
	fmt.Println(sshRun)
}

func getConfigFs(homeDirFs fs.FS) config.ConfigFs {
	return config.ConfigFs{
		Fsys:      homeDirFs,
		ConfigDir: ".config/ssh-run",
		SshPath:   ".ssh/config",
	}
}

func getHomeDirFs() (fs.FS, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return os.DirFS(homeDir), nil
}

func getCommands(args []string) []string {
	return shared.Filter(args, isCommand)
}

func getArgs(args []string) []string {
	return shared.Filter(args, isArgs)
}

func isCommand(s string) bool {
	return !isArgs(s)
}

func isArgs(s string) bool {
	return strings.HasPrefix(s, "--") || strings.HasPrefix(s, "-")
}
