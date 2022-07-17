package domain

const scriptsPath = ".config/ssh-run/scripts/"

func commonDir() string {
	return scriptsPath + "common"
}

func sharedDir() string {
	return scriptsPath + "shared"
}

func noScriptCommands() []string {
	return []string{"explain"}
}
