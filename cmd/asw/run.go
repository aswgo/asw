package main

import (
	"github.com/aswgo/asw/service"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:    "run",
	Short:  "Start project",
	Args:   cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) { setup() },
	RunE:   service.Run,
}

func configureRunCmd() {
}
