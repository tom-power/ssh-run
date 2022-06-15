package script

const scriptsPath = ".config/ssh-run/scripts/"

func hostDir(hostsName string) string {
	return scriptsPath + "host/" + hostsName
}

func commonDir() string {
	return scriptsPath + "common"
}
