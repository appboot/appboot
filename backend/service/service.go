package service

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/appboot/appbctl/config"
	"github.com/appboot/appbctl/creator"
	"github.com/appboot/appbctl/template"
	"github.com/appboot/appboot/constant"
	"github.com/appboot/appboot/git"
	"github.com/appboot/appboot/model"
)

// InitAppbctlConfig init appbctl config
func InitAppbctlConfig() {
	config.InitConfig()
}

// GetTemplates get templates
func GetTemplates() []string {
	return template.GetTemplates()
}

// UpdateAllTemplates update all templates
func UpdateAllTemplates() []string {
	if err := template.UpdateAllTemplatesWithGit(); err != nil {
		log.Printf("update all templates: %v", err)
	}
	return GetTemplates()
}

// GetConfig get config
func GetConfig(t string) *template.Config {
	return template.GetTemplateConfig(t)
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

	if err := creator.CreateWithCallback(application, true, callback); err != nil {
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
