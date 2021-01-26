package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/CatchZeng/gutils/file"
	"github.com/appboot/appboot/internal/pkg/logger"
	"github.com/appboot/appboot/internal/pkg/path"
	"github.com/spf13/cobra"
)

var createTemplateCmd = &cobra.Command{
	Use:   "create",
	Short: "create template",
	Long:  `create template from existed project`,
	Args:  cobra.MinimumNArgs(0),
	Run:   createTemplate,
}

func createTemplate(_ *cobra.Command, _ []string) {
	// project path
	projectPath, err := prompt("project path", "existed project path cannot be empty")
	if err != nil {
		logger.LogE(err)
		return
	}
	if !file.Exists(projectPath) {
		logger.LogE("project path is not existed")
		return
	}
	logger.LogI(projectPath)

	// Path
	savePath, err := prompt("save path", "save path cannot be empty")
	if err != nil {
		logger.LogE(err)
		return
	}
	savePath = path.HandleHomedir(savePath)
	if file.Exists(savePath) {
		result, err := promptSelect(fmt.Sprintf("%s already exists, whether to overwrite?", savePath))
		if err != nil {
			logger.LogE(err)
			return
		}
		if result == selectNo {
			return
		}
	}

	// Parameters
	logger.LogI("extract the parameters. Input wq! to end")

	parameters := make(map[string]string)

	for {
		var pk, pv string
		var err error
		pk, err = prompt("name", "the name will be extracted")
		if err != nil {
			logger.LogE(err)
		}
		if pk == "wq!" {
			break
		}
		pv, err = prompt("parameter name", "the parameter name")
		if err != nil {
			logger.LogE(err)
		}

		parameters[pk] = pv
	}

	logger.LogI(parameters)

	createTemplateFiles(projectPath, savePath, parameters)
}

func init() {
	templateCmd.AddCommand(createTemplateCmd)
}

func createTemplateFiles(projectPath, savePath string, params map[string]string) error {
	files, err := file.GetFiles(projectPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		savePath := strings.Replace(f.Path, projectPath, savePath, -1)
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
		result = strings.ReplaceAll(result, key, "{{."+value+"}}")
	}
	return result
}
