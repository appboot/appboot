package cmd

import "github.com/spf13/cobra"

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Template manager",
	Long:  `Template manager`,
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
