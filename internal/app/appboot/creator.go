package appboot

import (
	"errors"
	"fmt"
	"github.com/CatchZeng/gutils/file"
	gos "github.com/CatchZeng/gutils/os"
	"github.com/appboot/appboot/internal/pkg/logger"
	"os"
	"path"
	"strings"
)

// Callback app callback
type Callback func(app Application) error

// CreateCallback create callback
type CreateCallback struct {
	OnCreating Callback
	DidCreated Callback
}

// Create an application
func Create(app Application, force bool, skipPreSH bool, skipPostSH bool) error {
	return CreateWithCallback(app, force, skipPreSH, skipPostSH, nil)
}

// CreateWithCallback create an application with callback
func CreateWithCallback(app Application, force bool, skipPreSH bool, skipPostSH bool, callback *CreateCallback) error {
	if !force && file.Exists(app.Path) {
		return errors.New("the application already exists, you can force it to be created with the -f flag")
	}

	if !app.IsValid() {
		return errors.New("the application is invalid")
	}

	preScript := app.GetPreScript()
	if !skipPreSH && len(preScript) > 0 {
		logger.LogI("running script before the app is created")
		logger.LogW(preScript)
		if err := gos.RunBashCommand(preScript); err != nil {
			return err
		}
	}

	if callback != nil && callback.OnCreating != nil {
		if err := callback.OnCreating(app); err != nil {
			return err
		}
	}

	logger.LogI("creating all folders")
	if err := os.MkdirAll(app.Path, 0755); err != nil {
		return err
	}

	logger.LogI("creating all files")
	if err := createAllFiles(app); err != nil {
		return err
	}

	if callback != nil && callback.DidCreated != nil {
		if err := callback.DidCreated(app); err != nil {
			return err
		}
	}

	postScript := app.GetPostScript()
	if !skipPostSH && len(postScript) > 0 {
		logger.LogI("running script after the app is created")
		logger.LogW(postScript)
		if err := gos.RunBashCommand(postScript); err != nil {
			return err
		}
	}

	clean(app)

	logger.LogI("finish")
	return nil
}

func createAllFiles(app Application) error {
	templatePath := app.TemplatePath()

	files, err := GetFiles(templatePath)
	if err != nil {
		return err
	}

	params, err := app.GetParameters()
	if err != nil {
		return err
	}

	logger.LogI(fmt.Sprintf("Parameters:%v", params))

	for i := files.Front(); i != nil; i = i.Next() {
		templateFile := i.Value.(File)
		savePath := strings.Replace(templateFile.Path, templatePath, app.Path, -1)
		savePath = replaceWithParams(savePath, params)

		content := replaceWithParams(templateFile.Content, params)

		index := strings.LastIndex(savePath, "/")
		if index > 0 {
			dir := savePath[:index]
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		mode := file.Mode(templateFile.Path)
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

func clean(app Application) {
	configPath := path.Join(app.Path, ConfigFolder)
	if file.Exists(configPath) {
		_ = os.RemoveAll(configPath)
	}
}
