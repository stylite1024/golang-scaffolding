package cmd

import (
	"go-app/pkg/common"
	"go-app/pkg/config"
	"go-app/pkg/logger"
	"go-app/pkg/tools"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	rootCmd = &cobra.Command{
		Use:          "go-app",
		Short:        "go-app",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			common.Tip()
			common.HelpTip()
		},
	}
)

func setup() {
	applicationConfig := config.ApplicationConfig{
		Mode: "dev",
	}
	logConfig := config.LogConfig{
		LogLevel:    "info",
		LogFilename: "logs/app.log",
		MaxSize:     200,
		MaxAge:      30,
		MaxBackups:  7,
		Compress:    true,
	}
	// 初始化日志
	logger.Setup(&logConfig, applicationConfig.Mode)
	zap.L().Error(tools.Red("aaa"))
}

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
