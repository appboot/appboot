package server

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

func getSavePath(appName string) string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	savePath := path.Join(home, ".appboot", ".workspace", appName)
	return savePath
}
