package service

import (
	"os"
	"path"
	"strings"

	"github.com/aswgo/asw/pkg"
	"github.com/ngamux/ngamux"
	"github.com/spf13/cobra"
)

func GenRepository(cmd *cobra.Command, args []string) error {
	project := cmd.Flag("project")

	projectPath := ""
	if f := project.Value.String(); f != "" {
		projectPath = f
	} else {
		projectPath = "."
	}

	repositoryPath := path.Join(projectPath, "repository")
	_ = os.Mkdir(repositoryPath, 0755)

	err := pkg.FileCreateFromTemplate(path.Join(repositoryPath, args[0]+".repository.go"), "name_repository.go.tmpl", ngamux.Map{
		"name": strings.ToUpper(string(args[0][0])) + string(args[0][1:]),
	})
	if err != nil {
		return err
	}

	return nil
}
