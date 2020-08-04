package cmd

import (
	"fmt"
	"github.com/appboot/appboot/configs"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "appboot",
	Short: "Appboot is an application creation platform that helps you quickly create applications with templates.",
	Long:  "Appboot is an application creation platform that helps you quickly create applications with templates.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(configs.InitConfig)
}
