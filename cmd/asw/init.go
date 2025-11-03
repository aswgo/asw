package main

import (
	"github.com/aswgo/asw/service"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <path>",
	Short: "Create new project",
	Args:  cobra.MinimumNArgs(1),
	RunE:  service.Init,
}

func configureInitCmd() {
	initCmd.Flags().StringP("module", "m", "", "module name in go.mod")
}
