package cmd

import (
	"fmt"

	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/path"
	"github.com/go-ecosystem/utils/convert"
	"github.com/go-ecosystem/utils/file"
	"github.com/go-ecosystem/utils/log"
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
		templates := appboot.GetTemplates()
		if len(templates) < 1 {
			log.I("Updating templates...")
			if err := appboot.UpdateAllTemplates(); err != nil {
				log.E("Update templates error: %v", err)
				return
			}
			templates = appboot.GetTemplates()
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
		app.Path = path.HandleHomedir(savePath)

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

		skipPreSH, err := promptSelect("Skip pre script")
		if err != nil {
			log.E(err.Error())
			return
		}
		skipPostSH, err := promptSelect("Skip post script")
		if err != nil {
			log.E(err.Error())
			return
		}

		// Create
		if err := appboot.Create(app, true, skipPreSH == selectYes, skipPostSH == selectYes); err != nil {
			log.E(err.Error())
			return
		}
	},
}

func handleParams(params appboot.Parameters) map[string]string {
	result := make(map[string]string)
	log.H("Enter the parameters, if you need to use the default value, just press Enter.")

	stringParams := params.StringParameters
	if len(stringParams) > 0 {
		for _, param := range stringParams {
			value, err := promptStringParam(param)
			if err != nil {
				result = make(map[string]string)
				return result
			}
			result[param.Key] = value
		}
	}

	intParams := params.IntParameters
	if len(intParams) > 0 {
		for _, param := range intParams {
			value, err := promptIntParam(param)
			if err != nil {
				result = make(map[string]string)
				return result
			}
			result[param.Key] = value
		}
	}

	floatParams := params.FloatParameters
	if len(floatParams) > 0 {
		for _, param := range floatParams {
			value, err := promptFloatParam(param)
			if err != nil {
				result = make(map[string]string)
				return result
			}
			result[param.Key] = value
		}
	}

	selectParams := params.SelectParameters
	if len(selectParams) > 0 {
		for _, param := range selectParams {
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
