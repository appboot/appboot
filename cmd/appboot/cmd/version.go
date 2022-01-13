package cmd

import (
	"github.com/go-ecosystem/utils/log"

	v "github.com/go-ecosystem/utils/version"
	"github.com/spf13/cobra"
)

const (
	version   = "0.3.0"
	buildTime = "2022/01/13"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Appboot version",
	Long:  `Appboot version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	vs := v.Stringify(version, buildTime)
	log.I(vs)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
