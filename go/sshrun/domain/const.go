package domain

const scriptsPath = ".config/ssh-run/scripts/"

func commonDir() string {
	return scriptsPath + "common"
}

func utilsDir() string {
	return scriptsPath + "utils"
}
