package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/CatchZeng/gutils/file"
	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/logger"
	"github.com/appboot/appboot/internal/pkg/path"
	"github.com/spf13/cobra"
)

var createTemplateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a template",
	Long:  `Create a template from an existing project`,
	Args:  cobra.MinimumNArgs(0),
	Run:   createTemplate,
}

func createTemplate(_ *cobra.Command, _ []string) {
	// project path
	projectPath, err := prompt("existing project path", "existing project path cannot be empty")
	if err != nil {
		logger.LogE(err)
		return
	}
	if !file.Exists(projectPath) {
		logger.LogE("project path does not exist")
		return
	}
	logger.LogI(projectPath)

	// Path
	destinationPath, err := prompt("destination path", "destination path cannot be empty")
	if err != nil {
		logger.LogE(err)
		return
	}
	destinationPath = path.HandleHomedir(destinationPath)
	if file.Exists(destinationPath) {
		result, err := promptSelect(fmt.Sprintf("%s already exists, whether to overwrite?", destinationPath))
		if err != nil {
			logger.LogE(err)
			return
		}
		if result == selectNo {
			return
		}
	}

	// Parameters
	logger.LogI("extract the parameters.")

	parameters := make(map[string]string)

	for {
		var pk, pv string
		var err error
		pk, err = prompt("the name will be extracted", "the name cannot be empty")
		if err != nil {
			logger.LogE(err)
		}

		pv, err = prompt("parameter name", "the parameter name cannot be empty")
		if err != nil {
			logger.LogE(err)
		}

		parameters[pk] = pv

		r, err := promptSelect("finish extracting parameters?")
		if err != nil {
			logger.LogE(err)
			return
		}
		if r == selectYes {
			break
		}
	}

	logger.LogI(parameters)

	err = createTemplateFiles(projectPath, destinationPath, parameters)
	if err != nil {
		logger.LogE(err)
	}
}

func init() {
	templateCmd.AddCommand(createTemplateCmd)
}

func createTemplateFiles(projectPath, savePath string, params map[string]string) error {
	files, err := file.GetFiles(projectPath)
	if err != nil {
		return err
	}

	appbootPath := filepath.Join(projectPath, appboot.ConfigFolder)

	for _, f := range files {
		savePath := strings.Replace(f.Path, projectPath, savePath, -1)
		var content = f.Content

		if !strings.HasPrefix(f.Path, appbootPath) { // escape appboot folder
			savePath = replaceWithParams(savePath, params)
			content = replaceWithParams(f.Content, params)
		}

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
		result = strings.ReplaceAll(result, key, "{{."+value+"}}")
	}
	return result
}
