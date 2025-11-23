package service

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/aswgo/asw/pkg"
	"github.com/ngamux/ngamux"
	"github.com/spf13/cobra"
)

func GenController(cmd *cobra.Command, args []string) error {
	project := cmd.Flag("project")

	projectPath := ""
	if f := project.Value.String(); f != "" {
		projectPath = f
	} else {
		projectPath = "."
	}

	controllerPath := path.Join(projectPath, "controller")
	_ = os.Mkdir(controllerPath, 0755)

	_, err := os.Stat(path.Join(controllerPath, "controller.go"))
	if err != nil {
		err := pkg.FileCreateFromTemplate(path.Join(controllerPath, "controller.go"), "controller.go.tmpl")
		if err != nil {
			return err
		}

		gomod, err := pkg.FileGomodRead(path.Join(projectPath, "go.mod"))
		if err != nil {
			return err
		}

		err = pkg.FileWriteAfterMarker(
			path.Join(projectPath, "main.go"),
			"import (",
			fmt.Sprintf(`"%s/controller"`, gomod.Mod.String()),
		)
		if err != nil {
			return err
		}

		err = pkg.FileWriteAfterMarker(
			path.Join(projectPath, "main.go"),
			"import (",
			`"github.com/gowok/gowok"`,
		)
		if err != nil {
			return err
		}

		err = pkg.FileWriteAfterMarker(
			path.Join(projectPath, "main.go"),
			"func main()",
			"gowok.Configures(controller.Configure)",
		)
		if err != nil {
			return err
		}
	}

	err = pkg.FileCreateFromTemplate(path.Join(controllerPath, args[0]+".controller.go"), "name_controller.go.tmpl", ngamux.Map{
		"name": strings.ToUpper(string(args[0][0])) + string(args[0][1:]),
	})
	if err != nil {
		return err
	}

	return nil
}
