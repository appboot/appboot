package path

import (
	"os"
	"path"
	"strings"
)

// HandleDir HandleHomedir & HandlePWD
func HandleDir(filePath string) string {
	dir := HandleHomedir(filePath)
	dir = HandlePWD(dir)
	return dir
}

// HandleHomedir replace ~，$HOME，$home with homeDir
func HandleHomedir(filePath string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return filePath
	}

	const homeDir = "~"
	if strings.HasPrefix(filePath, homeDir) {
		result := strings.Replace(filePath, homeDir, home, 1)
		return result
	}

	const homeEnv = "$HOME"
	const startIndex = len(homeEnv)
	if strings.HasPrefix(filePath, homeEnv) || strings.HasPrefix(filePath, strings.ToLower(homeEnv)) {
		subPath := filePath[startIndex:]
		result := path.Join(home, subPath)
		return result
	}

	return filePath
}

// HandlePWD replace $PWD with Getwd
func HandlePWD(filePath string) string {
	wd, err := os.Getwd()
	if err != nil {
		return filePath
	}

	const pwd = "$PWD"
	if strings.HasPrefix(filePath, pwd) {
		result := strings.Replace(filePath, pwd, wd, 1)
		return result
	}

	return filePath
}
