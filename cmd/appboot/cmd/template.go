package cmd

import "github.com/spf13/cobra"

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "template manager",
	Long:  `template manager`,
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
