package cmd

import (
	"github.com/appboot/appboot/internal/app/appboot"
	"github.com/spf13/cobra"
	"log"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "appboot version",
	Long:  `appboot version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	versionString := appboot.GetVersion()
	log.Println(versionString)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
