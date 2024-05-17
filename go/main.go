package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/tom-power/ssh-run/sshrun"
	"github.com/tom-power/ssh-run/sshrun/config"
	"github.com/tom-power/ssh-run/sshrun/shared/generic"
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

	args := os.Args

	hostName := generic.GetOr(getCommands(args), 1, "")
	scriptName := generic.GetOr(getCommands(args), 2, "")

	flags := getFlags(os.Args)

	sshRunFlags := sshrun.RunFlags{
		Help:    generic.Any(flags, "--help"),
		Hosts:   generic.Any(flags, "--hosts"),
		Scripts: generic.Any(flags, "--scripts"),
		Explain: generic.Any(flags, "--explain"),
	}

	sshRun, err := runner.Run(hostName, scriptName, sshRunFlags, getScriptArgs(args))
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
	return generic.Filter(args, isNotFlag)
}

func getScriptArgs(args []string) []string {
	if len(args) > 2 {
		generic.Filter(args[2:], isNotFlag)
	}
	return []string{}
}

func getFlags(args []string) []string {
	return generic.Filter(args, isFlag)
}

var flags = []string{"--help", "--explain", "--hostName", "--scriptName"}

func isFlag(s string) bool {
	return generic.Any(flags, s)
}

func isNotFlag(s string) bool {
	return !isFlag(s)
}
