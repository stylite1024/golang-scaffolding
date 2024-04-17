package cmd

import (
	"go-app/pkg/common"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "print version info",
		Example: "go-app version",
		Run: func(cmd *cobra.Command, args []string) {
			common.Tip()
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
