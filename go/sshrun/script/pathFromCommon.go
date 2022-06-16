package script

func (fsys FileSys) pathFromCommon(scriptName string) (string, error) {
	return fsys.firstFileInDir(commonDir()+"/", scriptName)
}
