package script

func (fsys FileSys) pathCommon(scriptName string) (string, error) {
	return fsys.firstFileInDir(commonDir()+"/", scriptName)
}
