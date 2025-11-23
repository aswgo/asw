package cmd

import (
	"github.com/aswgo/asw/service"
	"github.com/spf13/cobra"
)

var Gen = &cobra.Command{
	Use:     "generate",
	Short:   "Generate any kind of file with given name",
	Aliases: []string{"g", "gen"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var genControllerCmd = &cobra.Command{
	Use:     "controller <name>",
	Short:   "Generate controller from name to name.controller.go and func Name()",
	Aliases: []string{"c"},
	Args:    cobra.MinimumNArgs(1),
	RunE:    service.GenController,
}

var genServiceCmd = &cobra.Command{
	Use:     "service <name>",
	Short:   "Generate service from name to name.service.go and func Name()",
	Aliases: []string{"s"},
	Args:    cobra.MinimumNArgs(1),
	RunE:    service.GenService,
}

var genRepositoryCmd = &cobra.Command{
	Use:     "repository <name>",
	Short:   "Generate repository from name to name.repository.go and func Name()",
	Aliases: []string{"r"},
	Args:    cobra.MinimumNArgs(1),
	RunE:    service.GenRepository,
}

func configureGenCmd() {
	Gen.Flags().StringP("project", "p", "", "project path")
	genControllerCmd.Flags().AddFlagSet(Gen.Flags())
	genServiceCmd.Flags().AddFlagSet(Gen.Flags())
	genRepositoryCmd.Flags().AddFlagSet(Gen.Flags())
	Gen.AddCommand(
		genControllerCmd,
		genServiceCmd,
		genRepositoryCmd,
	)
}
