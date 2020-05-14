package service

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/appboot/appboot/git"

	cnf "github.com/appboot/appboot/config"

	"github.com/appboot/appbctl/config"
	"github.com/appboot/appbctl/creator"
	"github.com/appboot/appboot/constant"
	"github.com/appboot/appboot/model"
	"github.com/appboot/appboot/utils"
)

const (
	appboot    = "appboot"
	configYaml = "appboot.yaml"
)

// GetTemplates get templates
func GetTemplates() []string {
	var templates []string

	root, err := config.GetTemplateRoot()
	if err != nil {
		return templates
	}

	templates, _ = utils.GetDirList(root)
	return templates
}

// GetConfig get config
func GetConfig(template string) *cnf.Config {
	var result *cnf.Config
	root, err := config.GetTemplateRoot()
	if err != nil {
		return result
	}

	yamlPath := path.Join(root, template, appboot, configYaml)

	result, _ = cnf.GetConfig(yamlPath)
	return result
}

// CreateApp create app
func CreateApp(app model.Application, callback *creator.CreateCallback) (constant.ErrCode, error) {
	application := app.Convert()

	if len(app.Name) < 1 || len(app.Template) < 1 {
		return constant.ErrEmpty, errors.New("application name and template can be empty")
	}

	if strings.Contains(app.Name, " ") {
		return constant.ErrContainBlanks, errors.New("application name can not contain blanks")
	}

	_ = os.RemoveAll(application.Path)

	if err := creator.CreateWithCallback(application, true, false, callback); err != nil {
		return constant.ErrCreate, err
	}

	return constant.OK, nil
}

// PushCode push code
func PushCode(app model.Application) error {
	if len(app.Git) < 1 {
		return nil
	}

	codeFolder := app.Convert().Path
	return git.Push(app.Git, codeFolder)
}
