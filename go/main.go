package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

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

	hostName := shared.GetOr(os.Args, 1, "")
	scriptName := shared.GetOr(os.Args, 2, "")
	args := []string{""}
	if len(os.Args) > 3 {
		args = os.Args[3:]
	}

	sshRun, err := sshrun.Runner{Config: config, Fsys: homeDirFs}.Run(hostName, scriptName, args)
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
