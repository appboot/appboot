package appboot

import (
	"errors"
	"os"

	"github.com/go-ecosystem/utils/file"
	"github.com/go-ecosystem/utils/log"
	gos "github.com/go-ecosystem/utils/os"
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

	params, _ := app.GetParameters()
	log.I("Parameters:%v", params)

	preScript := app.GetPreScript()
	if !skipPreSH && len(preScript) > 0 {
		log.I("running script before the app is created")
		log.W(preScript)
		if err := gos.RunBashCommand(preScript); err != nil {
			return err
		}
	}

	if callback != nil && callback.OnCreating != nil {
		if err := callback.OnCreating(app); err != nil {
			return err
		}
	}

	log.I("creating folders")
	if err := os.MkdirAll(app.Path, 0755); err != nil {
		return err
	}

	log.I("creating files")
	if err := app.CreateFiles(); err != nil {
		return err
	}

	if callback != nil && callback.DidCreated != nil {
		if err := callback.DidCreated(app); err != nil {
			return err
		}
	}

	postScript := app.GetPostScript()
	if !skipPostSH && len(postScript) > 0 {
		log.I("running script after the app is created")
		log.W(postScript)
		if err := gos.RunBashCommand(postScript); err != nil {
			return err
		}
	}

	app.Clean()

	log.I("finish")
	return nil
}
