package cmd

import (
	"github.com/aswgo/asw/service"
	"github.com/spf13/cobra"
)

var Run = &cobra.Command{
	Use:    "run",
	Short:  "Start project",
	PreRun: func(cmd *cobra.Command, args []string) { setup() },
	RunE:   service.Run,
}

func configureRunCmd() {
	Run.Flags().BoolP("watch", "w", false, "reload every code changed")
}
