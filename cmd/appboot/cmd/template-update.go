package cmd

import (
	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/go-ecosystem/utils/v2/log"
	"github.com/spf13/cobra"
)

var updateTemplateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update template",
	Long:  `Update template`,
	Run:   runUpdateTemplate,
}

func runUpdateTemplate(_ *cobra.Command, _ []string) {
	templates := appboot.GetTemplateNames()
	if len(templates) < 1 {
		log.I("Updating templates...")
		if err := appboot.UpdateAllTemplates(); err != nil {
			log.E("Update templates error: %v", err)
		}
		return
	}

	const All = "All"
	templates = append(templates, All)
	selectedTemplate, err := promptSelectWithItems("Select template", templates)
	if err != nil {
		log.E(err.Error())
		return
	}

	if selectedTemplate == All {
		if err := appboot.UpdateAllTemplates(); err != nil {
			log.E(err.Error())
			return
		}
	} else {
		if err := appboot.UpdateTemplate(selectedTemplate); err != nil {
			log.E(err.Error())
			return
		}
	}
}

func init() {
	templateCmd.AddCommand(updateTemplateCmd)
}
