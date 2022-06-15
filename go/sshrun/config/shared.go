package config

import "io/fs"

func getBytes(file fs.File) ([]byte, error) {
	stat, _ := file.Stat()
	bytes := make([]byte, stat.Size())
	_, err := file.Read(bytes)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}
