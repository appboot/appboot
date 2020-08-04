package path

import (
	"strings"

	"github.com/mitchellh/go-homedir"
)

// HandleHomedir replace ~ with homeDir
func HandleHomedir(filePath string) string {
	const homeDir = "~"
	if strings.HasPrefix(filePath, homeDir) {
		home, err := homedir.Dir()
		if err != nil {
			return filePath
		}
		result := strings.Replace(filePath, homeDir, home, 1)
		return result
	}
	return filePath
}
