package cmd

import (
	"github.com/aswgo/asw/service"
	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:   "init <path>",
	Short: "Create new project",
	RunE:  service.Init,
}

func configureInitCmd() {
}
