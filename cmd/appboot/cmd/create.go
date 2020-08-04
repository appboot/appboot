package cmd

import (
	"errors"
	"fmt"
	"github.com/CatchZeng/gutils/convert"
	"github.com/CatchZeng/gutils/file"
	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/logger"
	"github.com/appboot/appboot/internal/pkg/path"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"strconv"
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
			logger.LogI("updating templates...")
			if err := appboot.UpdateAllTemplates(); err != nil {
				logger.LogE(fmt.Sprintf("update templates error: %v", err))
				return
			}
			templates = appboot.GetTemplates()
			if len(templates) < 1 { // check again
				logger.LogE("Without any template, the application cannot be created. Please check the configuration item of templatesSource")
				return
			}
		}

		selectedTemplate, err := promptTemplates(templates)
		if err != nil {
			logger.LogE(err)
			return
		}
		app.Template = selectedTemplate

		// Name
		name, err := prompt("name", "application name cannot be empty")
		if err != nil {
			logger.LogE(err)
			return
		}
		app.Name = name

		// Path
		savePath, err := prompt("path", "application path cannot be empty")
		if err != nil {
			logger.LogE(err)
			return
		}
		app.Path = path.HandleHomedir(savePath)

		if file.Exists(app.Path) {
			result, err := promptSelect(fmt.Sprintf("%s already exists, whether to overwrite?", app.Path))
			if err != nil {
				logger.LogE(err)
				return
			}
			if result == selectNo {
				return
			}
		}

		// Params
		cnf, err := appboot.GetTemplateConfig(selectedTemplate)
		if err != nil {
			logger.LogE(err)
			return
		}
		params := handleParams(cnf.Parameters)
		log.Print(params)

		valueString, err := convert.MapToJSON(params)
		if err != nil {
			logger.LogE(err)
			return
		}
		app.Parameters = valueString

		skipPreSH, err := promptSelect("skip pre script?")
		if err != nil {
			logger.LogE(err)
			return
		}
		skipPostSH, err := promptSelect("skip post script?")
		if err != nil {
			logger.LogE(err)
			return
		}

		// Create
		if err := appboot.Create(app, true, skipPreSH == selectYes, skipPostSH == selectYes); err != nil {
			logger.LogE(err)
			return
		}
	},
}

const (
	selectYes string = "YES"
	selectNo  string = "NO"
)

func promptSelect(label string) (string, error) {
	items := []string{selectYes, selectNo}
	return promptSelectWithItems(label, items)
}

func promptSelectWithItems(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return result, err
	}
	return result, nil
}

func promptTemplates(template []string) (string, error) {
	prompt := promptui.Select{
		Label: "select template",
		Items: template,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return result, err
	}
	return result, nil
}

func prompt(label string, alert string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: emptyValidate(alert),
	}
	return prompt.Run()
}

func promptStringParam(param appboot.StringParameter) (string, error) {
	prompt := promptui.Prompt{
		Label:    param.Key,
		Validate: emptyValidate(param.Key + " can not be empty"),
		Default:  param.Default,
	}
	return prompt.Run()
}

func promptIntParam(param appboot.IntParameter) (string, error) {
	prompt := promptui.Prompt{
		Label:    param.Key,
		Validate: intParamValidate(param, param.Key+" can not be empty"),
		Default:  strconv.FormatInt(param.Default, 10),
	}
	return prompt.Run()
}

func promptFloatParam(param appboot.FloatParameter) (string, error) {
	prompt := promptui.Prompt{
		Label:    param.Key,
		Validate: floatParamValidate(param, param.Key+" can not be empty"),
		Default:  fmt.Sprintf("%f", param.Default),
	}
	return prompt.Run()
}

func emptyValidate(alert string) promptui.ValidateFunc {
	return func(input string) error {
		if len(input) < 1 {
			return errors.New(alert)
		}
		return nil
	}
}

func intParamValidate(param appboot.IntParameter, alert string) promptui.ValidateFunc {
	return func(input string) error {
		if len(input) < 1 {
			return errors.New(alert)
		}

		intValue, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return err
		}

		if intValue < param.Min || intValue > param.Max {
			return fmt.Errorf("the value of %v must be in the range of %d to %d", param.Key, param.Min, param.Max)
		}
		return nil
	}
}

func floatParamValidate(param appboot.FloatParameter, alert string) promptui.ValidateFunc {
	return func(input string) error {
		if len(input) < 1 {
			return errors.New(alert)
		}

		floatValue, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return err
		}

		if floatValue < param.Min || floatValue > param.Max {
			return fmt.Errorf("the value of %v must be in the range of %f to %f", param.Key, param.Min, param.Max)
		}
		return nil
	}
}

func handleParams(params appboot.Parameters) map[string]string {
	result := make(map[string]string)
	logger.LogI("enter the parameters, if you need to use the default value, just press Enter.")

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
