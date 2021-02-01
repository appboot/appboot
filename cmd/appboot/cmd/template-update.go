package cmd

import (
	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/logger"
	"github.com/spf13/cobra"
)

var updateTemplateCmd = &cobra.Command{
	Use:   "update",
	Short: "update template",
	Long:  `update template`,
	Run:   runUpdateTemplate,
}

func runUpdateTemplate(_ *cobra.Command, _ []string) {
	templates := appboot.GetTemplates()
	if len(templates) < 1 {
		logger.LogI("updating templates...")
		if err := appboot.UpdateAllTemplates(); err != nil {
			logger.LogE("update templates error: %v", err)
		}
		return
	}

	const All = "All"
	templates = append(templates, All)
	selectedTemplate, err := promptSelectWithItems("select template", templates)
	if err != nil {
		logger.LogE(err.Error())
		return
	}

	if selectedTemplate == All {
		if err := appboot.UpdateAllTemplates(); err != nil {
			logger.LogE(err.Error())
			return
		}
	} else {
		if err := appboot.UpdateTemplate(selectedTemplate); err != nil {
			logger.LogE(err.Error())
			return
		}
	}
}

func init() {
	templateCmd.AddCommand(updateTemplateCmd)
}
