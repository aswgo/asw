package service

import (
	"github.com/aswgo/asw/pkg"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	// err := os.Mkdir(args[0], 0755)
	// if err != nil {
	// 	return err
	// }

	// fModule := cmd.Flag("module")
	// err = pkg.ExecCommandInDir(args[0], "go", "mod", "init", cmp.Or(fModule.Value.String(), args[0]))
	// if err != nil {
	// 	return err
	// }
	//
	// err = pkg.FileCreateFromTemplate(path.Join(args[0], "main.go"), "main.go.tmpl")
	// if err != nil {
	// 	return err
	// }

	pathConfig, err := pkg.PathJoinCWD("asw.toml")
	if err != nil {
		return err
	}

	err = pkg.FileCreateFromTemplate(pathConfig, "asw.toml.tmpl")
	if err != nil {
		return err
	}

	// err = pkg.ExecCommandInDir(args[0], "go", "mod", "tidy")
	// if err != nil {
	// 	return err
	// }

	return nil
}
