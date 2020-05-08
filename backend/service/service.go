package service

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/appboot/appboot/git"

	"github.com/appboot/appboot/parameter"

	"github.com/appboot/appbctl/config"
	"github.com/appboot/appbctl/creator"
	"github.com/appboot/appboot/constant"
	"github.com/appboot/appboot/model"
	"github.com/appboot/appboot/utils"
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

// GetParams get params
func GetParams(template string) *parameter.Parameters {
	var params parameter.Parameters
	root, err := config.GetTemplateRoot()
	if err != nil {
		return &params
	}

	yamlPath := path.Join(root, template, "appboot", "parameters.yaml")

	result, err := parameter.GetParameters(yamlPath)
	if err != nil {
		return &params
	}
	return result
}

// CreateApp create app
func CreateApp(app model.Application) (constant.ErrCode, error) {
	application := app.Convert()

	if len(app.Name) < 1 || len(app.Template) < 1 {
		return constant.ErrEmpty, errors.New("application name and template can be empty")
	}

	if strings.Contains(app.Name, " ") {
		return constant.ErrContainBlanks, errors.New("application name can not contain blanks")
	}

	_ = os.RemoveAll(application.Path)

	if err := creator.Create(application, true, false); err != nil {
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
