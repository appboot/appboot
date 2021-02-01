package cmd

import (
	"github.com/go-ecosystem/utils/log"

	v "github.com/go-ecosystem/utils/version"
	"github.com/spf13/cobra"
)

const (
	version   = "0.1.0"
	buildTime = "2020/08/31"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "appboot version",
	Long:  `appboot version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	vs := v.Stringify(version, buildTime)
	log.H(vs)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
