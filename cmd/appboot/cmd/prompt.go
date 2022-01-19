package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/manifoldco/promptui"
)

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
		Default:  strconv.Itoa(param.Default),
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

		intValue, err := strconv.Atoi(input)
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
