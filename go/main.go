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
	homeDirFs, err := getHomeDirFs()
	if err != nil {
		log.Fatal(err)
	}
	configFs := config.FileSys{
		Fsys:       homeDirFs,
		ConfigPath: ".config/ssh-run/config.yaml",
		SshPath:    ".ssh/config",
	}
	config, err := configFs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	scriptFs := script.FileSys{
		Fsys:   homeDirFs,
		Config: config,
	}
	sshRun, err := getRun(config, configFs, scriptFs)(hostName, scriptName, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sshRun)
}

func getRun(
	config shared.Config,
	configFs config.FileSys,
	scriptFs script.FileSys,
) sshrun.Run {
	return sshrun.GetRun(sshrun.GetCommandSsh, config, scriptFs)
}

func getHomeDirFs() (fs.FS, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return os.DirFS(homeDir), nil
}
