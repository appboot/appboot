package git

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/CatchZeng/gutils/os"
	"gopkg.in/yaml.v2"
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
		"git add ." + " && " + "git commit -m \"Initial commit\"" + " && " + "git push -u origin master"
	log.Printf("%s", cmd)
	return runBashCommand(cmd)
}

// Parameters git parameters
type Parameters struct {
	Prefix string `yaml:"prefix" json:"prefix"`
}

// GetParameters get parameters from yaml path
func GetParameters(yamlPath string) (git *Parameters, err error) {
	git = new(Parameters)
	var yamlFile []byte
	if yamlFile, err = ioutil.ReadFile(yamlPath); err != nil {
		return
	}
	if err = yaml.Unmarshal(yamlFile, git); err != nil {
		return
	}
	return
}
