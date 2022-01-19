package appboot

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/appboot/appboot/configs"
	"github.com/go-ecosystem/utils/convert"
	"github.com/go-ecosystem/utils/file"
)

const (
	// Name template Name
	Name = "Name"
	// Path template Path
	Path = "Path"

	// ConfigFolder appboot configuration items
	ConfigFolder = "appboot"
)

// Application data struct
type Application struct {
	Name       string
	Path       string
	Template   string
	Parameters string
}

// Description application description
func (app *Application) Description() string {
	return fmt.Sprintf("Name:%s \nPath:%s \nTemplate:%s \nParameters:%s\n", app.Name, app.Path, app.Template, app.Parameters)
}

// GetTemplatePath application template path
func (app *Application) GetTemplatePath() string {
	root, err := configs.GetTemplateRoot()
	if err != nil {
		return ""
	}
	templatePath := path.Join(root, app.Template)
	return templatePath
}

// GetParameters get application parameters
func (app *Application) GetParameters() (map[string]string, error) {
	parameters, err := convert.JSONToMap(app.Parameters)
	if err != nil {
		return nil, err
	}

	if _, ok := parameters[Name]; !ok {
		parameters[Name] = app.Name
	}
	if _, ok := parameters[Path]; !ok {
		parameters[Path] = app.Path
	}
	return parameters, nil
}

// IsValid is it valid
func (app *Application) IsValid() bool {
	return len(app.Name) > 0 &&
		len(app.Path) > 0 &&
		len(app.Template) > 0
}

// Clean clean config folder
func (app *Application) Clean() {
	configPath := path.Join(app.Path, ConfigFolder)
	if file.Exists(configPath) {
		_ = os.RemoveAll(configPath)
	}
}

// CreateFiles create files
func (app Application) CreateFiles() error {
	templatePath := app.GetTemplatePath()

	files, err := file.GetFiles(templatePath)
	if err != nil {
		return err
	}

	params, err := app.GetParameters()
	if err != nil {
		return err
	}

	for _, f := range files {
		savePath := strings.Replace(f.Path, templatePath, app.Path, -1)
		savePath = replaceWithParams(savePath, params)

		content := replaceWithParams(f.Content, params)

		index := strings.LastIndex(savePath, "/")
		if index > 0 {
			dir := savePath[:index]
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		mode := file.Mode(f.Path)
		if err := file.WriteStringToFile(content, savePath, mode); err != nil {
			return err
		}
	}

	return nil
}

func replaceWithParams(source string, params map[string]string) string {
	var result = source
	for key, value := range params {
		result = strings.ReplaceAll(result, "{{."+key+"}}", value)
	}
	return result
}
