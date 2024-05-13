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
	configFs := config.ConfigFs{
		Fsys:      homeDirFs,
		ConfigDir: ".config/ssh-run",
		SshPath:   ".ssh/config",
	}
	config, err := configFs.GetConfig()
	if err != nil {
		log.Fatal("config error:", err)
	}

	hostName := shared.GetOr(getCommands(os.Args), 1, "")
	scriptName := shared.GetOr(getCommands(os.Args), 2, "")
	args := getArgs(os.Args)

	sshRun, err := sshrun.Runner{Config: config, Fsys: homeDirFs}.Run(hostName, scriptName, args, []string{""})
	if err != nil {
		log.Fatal("runner error:", err)
	}
	fmt.Printf(sshRun)
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
