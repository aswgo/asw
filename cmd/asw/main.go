package main

import (
	"log/slog"
	"os"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/logger"
	"github.com/spf13/cobra"
)

func setup() {
	gowok.Configures(
		logger.Configure(slog.NewJSONHandler(os.Stdout, nil).
			WithAttrs([]slog.Attr{
				slog.Any("app", "asw"),
			}),
		),
	)
}

var rootCmd = &cobra.Command{
	Use:   "asw",
	Short: "Make your holiday happen today!",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	gowok.CMD(rootCmd)

	configureInitCmd()
	configureRunCmd()
	configureGenCmd()

	gowok.AddCMD(
		initCmd,
		runCmd,
		genCmd,
	)
	gowok.CMD().Execute()
}
