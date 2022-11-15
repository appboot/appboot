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
func Create(app Application,
	force bool,
	beforeScripts []string,
	afterScripts []string,
	skipBeforeScripts bool,
	skipAfterScripts bool) error {
	return CreateWithCallback(app, force, beforeScripts, afterScripts, skipBeforeScripts, skipAfterScripts, nil)
}

// CreateWithCallback create an application with callback
func CreateWithCallback(app Application,
	force bool,
	beforeScripts []string,
	afterScripts []string,
	skipBeforeScripts bool,
	skipAfterScripts bool,
	callback *CreateCallback) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// Avoid `shell-init: error retrieving current directory: getcwd` error when the generated code is deleted due to calling os.Chdir when executing afterScripts
	defer os.Chdir(home)

	if !force && file.Exists(app.Path) {
		return errors.New("the application already exists, you can force it to be created with the -f flag")
	}

	if !app.IsValid() {
		return errors.New("the application is invalid")
	}

	if !skipBeforeScripts && len(beforeScripts) > 0 {
		log.H("Running script before the app is created")
		for _, script := range beforeScripts {
			log.W(script)
			if err := gos.RunBashCommand(script); err != nil {
				return err
			}
		}
	}

	if callback != nil && callback.OnCreating != nil {
		if err := callback.OnCreating(app); err != nil {
			return err
		}
	}

	log.H("Creating folders")
	if err := os.MkdirAll(app.Path, 0755); err != nil {
		return err
	}

	log.H("Creating files")
	if err := app.CreateFiles(); err != nil {
		return err
	}

	if callback != nil && callback.DidCreated != nil {
		if err := callback.DidCreated(app); err != nil {
			return err
		}
	}

	if !skipAfterScripts && len(afterScripts) > 0 {
		log.H("Running script after the app is created")

		// changes the current working directory to the app's directory
		if file.Exists(app.Path) {
			os.Chdir(app.Path)
		}

		for _, script := range afterScripts {
			log.W(script)
			if err := gos.RunBashCommand(script); err != nil {
				return err
			}
		}
	}

	app.Clean()

	log.H("Finish")
	return nil
}
