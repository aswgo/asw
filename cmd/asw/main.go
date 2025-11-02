package main

import (
	"log/slog"
	"os"

	"github.com/aswgo/asw/service"
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

var initCmd = &cobra.Command{
	Use:   "init <path>",
	Short: "Create new project",
	Args:  cobra.MinimumNArgs(1),
	RunE:  service.Init,
}

var runCmd = &cobra.Command{
	Use:    "run",
	Short:  "Start project",
	Args:   cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) { setup() },
	RunE:   service.Run,
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate any kind of file with given name",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var genControllerCmd = &cobra.Command{
	Use:   "controller <name>",
	Short: "Generate controller from name to name.controller.go and func Name()",
	Args:  cobra.MinimumNArgs(1),
	RunE:  service.GenController,
}

var genServiceCmd = &cobra.Command{
	Use:   "service <name>",
	Short: "Generate service from name to name.service.go and func Name()",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var genRepositoryCmd = &cobra.Command{
	Use:   "repository <name>",
	Short: "Generate repository from name to name.repository.go and func Name()",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	gowok.CMD(rootCmd)

	initCmd.Flags().StringP("module", "m", "", "module name in go.mod")

	genCmd.Flags().StringP("project", "p", "", "project path")
	genControllerCmd.Flags().AddFlagSet(genCmd.Flags())
	genServiceCmd.Flags().AddFlagSet(genCmd.Flags())
	genRepositoryCmd.Flags().AddFlagSet(genCmd.Flags())
	genCmd.AddCommand(
		genControllerCmd,
		genServiceCmd,
		genRepositoryCmd,
	)

	gowok.AddCMD(
		initCmd,
		runCmd,
		genCmd,
	)
	gowok.CMD().Execute()
}
