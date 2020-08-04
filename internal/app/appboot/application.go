package appboot

import (
	"fmt"
	"github.com/CatchZeng/gutils/convert"
	"github.com/CatchZeng/gutils/file"
	"github.com/appboot/appboot/configs"
	"path"
)

const (
	// Name template Name
	Name = "Name"
	// Path template Path
	Path = "Path"

	// ConfigFolder appboot configuration items
	ConfigFolder = "appboot"
	// PreSH appboot pre script
	PreSH = "pre.sh"
	// PostSH appboot post script
	PostSH = "post.sh"
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

// TemplatePath application template path
func (app *Application) TemplatePath() string {
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

// GetPreScript  pre script
func (app *Application) GetPreScript() string {
	templatePath := app.TemplatePath()
	prePath := path.Join(templatePath, ConfigFolder, PreSH)
	if file.Exists(prePath) {
		cmd := "sh " + prePath
		return cmd
	}
	return ""
}

// GetPostScript post script
func (app *Application) GetPostScript() string {
	postPath := path.Join(app.Path, ConfigFolder, PostSH)
	if file.Exists(postPath) {
		cmd := "sh " + postPath
		return cmd
	}
	return ""
}

// IsValid is it valid
func (app *Application) IsValid() bool {
	return len(app.Name) > 0 &&
		len(app.Path) > 0 &&
		len(app.Template) > 0
}
