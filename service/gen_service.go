package service

import (
	"os"
	"path"
	"strings"

	"github.com/aswgo/asw/pkg"
	"github.com/ngamux/ngamux"
	"github.com/spf13/cobra"
)

func GenService(cmd *cobra.Command, args []string) error {
	project := cmd.Flag("project")

	projectPath := ""
	if f := project.Value.String(); f != "" {
		projectPath = f
	} else {
		projectPath = "."
	}

	servicePath := path.Join(projectPath, "service")
	_ = os.Mkdir(servicePath, 0755)

	err := pkg.FileCreateFromTemplate(path.Join(servicePath, args[0]+".service.go"), "name_service.go.tmpl", ngamux.Map{
		"name": strings.ToUpper(string(args[0][0])) + string(args[0][1:]),
	})
	if err != nil {
		return err
	}

	return nil
}
