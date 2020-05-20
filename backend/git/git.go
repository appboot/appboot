package git

import (
	"errors"
	"log"

	"github.com/CatchZeng/gutils/os"
)

var runBashCommand = os.RunBashCommand

// Push push code to git
func Push(gitURL string, codeFolder string) error {
	if len(codeFolder) < 1 {
		return errors.New("code floder is empty")
	}
	if len(gitURL) < 1 {
		return errors.New("gitURL is empty")
	}

	cmd := "cd " + codeFolder + " && " + "git init" + " && " + "git remote add origin " + gitURL + " && " +
		"git add ." + " && " + "git reset appboot" + " && " + "git commit -m \"Initial commit\"" + " && " + "git push -u origin master"
	log.Printf("%s", cmd)
	return runBashCommand(cmd)
}
