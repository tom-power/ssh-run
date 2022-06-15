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
	sshRun, err := getRun(configFs, homeDirFs)(hostName, scriptName, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sshRun)
}

func getRun(configFs config.FileSys, fsys fs.FS) sshrun.Run {
	config, err := configFs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	return sshrun.GetRun(
		script.GetScriptPathFromConf(fsys),
		script.GetScriptContentsFromHost(fsys),
		sshrun.GetCommandSsh,
		script.GetScriptsFromConf(fsys),
		config,
	)
}

func getHomeDirFs() (fs.FS, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return os.DirFS(homeDir), nil
}
