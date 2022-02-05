package files

import (
	"os"
	"path/filepath"
)

func FindGluetunDir() (path string, err error) {
	path, err = filepath.Abs("gluetun")
	if err == nil && isDir(path) {
		return path, nil
	}

	path = "/gluetun"
	if isDir(path) {
		return path, nil
	}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		homeSubPaths := []string{"gluetun", ".config.gluetun"}
		for _, subPath := range homeSubPaths {
			path = filepath.Join(homeDir, subPath)
			if isDir(path) {
				return path, nil
			}
		}
	}

	path, err = filepath.Abs("gluetun")
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(path, os.ModeDir)
	if err != nil {
		return "", err
	}

	return path, nil
}

func isDir(path string) (valid bool) {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
