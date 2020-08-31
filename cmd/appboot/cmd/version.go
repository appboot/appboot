package cmd

import (
	v "github.com/CatchZeng/gutils/version"
	"github.com/spf13/cobra"
	"log"
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
	log.Println(vs)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
