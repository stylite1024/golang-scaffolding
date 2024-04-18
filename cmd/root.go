package cmd

import (
	"go-app/pkg/common"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "go-app",
		Short:        "go-app",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			common.Tip()
			common.HelpTip()
		},
	}
)

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func init() {
	// 关闭官方completion命令
	completion := completionCommand()
	completion.Hidden = true
	rootCmd.AddCommand(completion)
	// 智能提示最小位数
	rootCmd.SuggestionsMinimumDistance = 1
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
