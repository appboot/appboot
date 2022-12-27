package server

import (
	"log"
	"os"
	"path"

	"github.com/appboot/appboot/internal/pkg/common"
	"github.com/go-ecosystem/utils/v2/file"
)

func getSavePath(appName string) string {
	p := path.Join(getWorkspacePath(), appName)
	createIfNoExist(p)
	return p
}

func getStaticPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	p := path.Join(home, ".appboot", "static")
	createIfNoExist(p)
	return p
}

func getWorkspacePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	p := path.Join(home, ".appboot", "workspace")
	createIfNoExist(p)
	return p
}

func createIfNoExist(path string) {
	if !file.Exists(path) {
		_ = os.MkdirAll(path, common.DefaultFileMode)
	}
}
