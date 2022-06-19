package domain

func (h Host) pathRemote(scriptName string) (string, error) {
	hostDir := hostDirWithHome(h.Name, "/home/"+h.User)
	command := "" +
		"cd " + hostDir + " &&" +
		"find . -type f -name '" + scriptName + "*.sh' | sed 's/\\.\\///'"
	hostScriptPathRemote, err := runCommandOn(h, command)
	if err != nil {
		return "", err
	}
	return hostDir + hostScriptPathRemote, nil
}
