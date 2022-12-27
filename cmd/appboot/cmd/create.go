package cmd

import (
	"fmt"

	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/path"
	"github.com/go-ecosystem/utils/v2/convert"
	"github.com/go-ecosystem/utils/v2/file"
	"github.com/go-ecosystem/utils/v2/log"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Create an application",
	Long:  `Create an application`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		app := appboot.Application{}

		// Template
		templates := appboot.GetTemplateNames()
		if len(templates) < 1 {
			log.I("Updating templates...")
			if err := appboot.UpdateAllTemplates(); err != nil {
				log.E("Update templates error: %v", err)
				return
			}
			templates = appboot.GetTemplateNames()
			if len(templates) < 1 { // check again
				log.E("Without any template, the application cannot be created. Please check the configuration item of templatesSource")
				return
			}
		}

		selectedTemplate, err := promptSelectWithItems("Select template", templates)
		if err != nil {
			log.E(err.Error())
			return
		}
		app.Template = selectedTemplate

		// Name
		name, err := prompt("Name", "Application name cannot be empty.")
		if err != nil {
			log.E(err.Error())
			return
		}
		app.Name = name

		// Path
		savePath, err := prompt("Path", "Application path cannot be empty.")
		if err != nil {
			log.E(err.Error())
			return
		}
		app.Path = path.HandleDir(savePath)

		if file.Exists(app.Path) {
			result, err := promptSelect(fmt.Sprintf("%s already exists, whether to overwrite?", app.Path))
			if err != nil {
				log.E(err.Error())
				return
			}
			if result == selectNo {
				return
			}
		}

		// Params
		cnf, err := appboot.GetTemplateConfig(selectedTemplate)
		if err != nil {
			log.E(err.Error())
			return
		}
		params := handleParams(cnf.Parameters)
		log.I("Parameters: %v", params)

		valueString, err := convert.MapToJSON(params)
		if err != nil {
			log.E(err.Error())
			return
		}
		app.Parameters = valueString

		skipBeforeScripts, err := promptSelect("Skip executing before scripts")
		if err != nil {
			log.E(err.Error())
			return
		}
		skipAfterScripts, err := promptSelect("Skip executing after scripts")
		if err != nil {
			log.E(err.Error())
			return
		}

		// Create
		if err := appboot.Create(app,
			true,
			cnf.Scripts.Before,
			cnf.Scripts.After,
			skipBeforeScripts == selectYes,
			skipAfterScripts == selectYes); err != nil {
			log.E(err.Error())
			return
		}
	},
}

func handleParams(params []interface{}) map[string]string {
	result := make(map[string]string)
	log.H("Enter the parameters, if you need to use the default value, just press Enter.")

	for _, v := range params {
		switch param := v.(type) {
		case appboot.StringParameter:
			value, err := promptStringParam(param)
			if err != nil {
				result = make(map[string]string)
				return result
			}
			result[param.Key] = value

		case appboot.IntParameter:
			value, err := promptIntParam(param)
			if err != nil {
				result = make(map[string]string)
				return result
			}
			result[param.Key] = value

		case appboot.FloatParameter:
			value, err := promptFloatParam(param)
			if err != nil {
				result = make(map[string]string)
				return result
			}
			result[param.Key] = value

		case appboot.SelectParameter:
			value, err := promptSelectWithItems(param.Key, param.Options)
			if err != nil {
				result = make(map[string]string)
				return result
			}
			result[param.Key] = value
		}
	}
	return result
}

func init() {
	rootCmd.AddCommand(create)
}
