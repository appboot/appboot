package cmd

import (
	"fmt"

	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/appboot/appboot/internal/pkg/logger"
	"github.com/spf13/cobra"
)

var listTemplateCmd = &cobra.Command{
	Use:   "list",
	Short: "list templates",
	Long:  `list templates`,
	Run:   runListTemplate,
}

func runListTemplate(_ *cobra.Command, _ []string) {
	templates := appboot.GetTemplates()
	for i, t := range templates {
		logger.Log(logger.Blue, fmt.Sprintf("%d. %v", i+1, t))
	}
}

func init() {
	templateCmd.AddCommand(listTemplateCmd)
}
