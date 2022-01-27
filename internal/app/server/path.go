package server

import (
	"log"
	"os"
	"path"

	"github.com/go-ecosystem/utils/file"
	"github.com/mitchellh/go-homedir"
)

func getSavePath(appName string) string {
	p := path.Join(getWorkspacePath(), appName)
	createIfNoExist(p)
	return p
}

func getStaticPath() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	p := path.Join(home, ".appboot", "static")
	createIfNoExist(p)
	return p
}

func getWorkspacePath() string {
	home, err := homedir.Dir()
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
		os.MkdirAll(path, 0755)
	}
}
