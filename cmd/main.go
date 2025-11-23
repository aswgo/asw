package cmd

import (
	"flag"
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

var Root = &cobra.Command{
	Use:   "asw",
	Short: "Make your holiday happen today!",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	PreRun: func(cmd *cobra.Command, args []string) { setup() },
	RunE:   service.Run,
}

func Configure() {
	configureInitCmd()
	configureRunCmd()
	configureGenCmd()

	flag.Bool("watch", false, "reload every code changed")

	gowok.CMD.Command = gowok.CMD.Wrap(Root)
	gowok.CMD.AddCommand(
		Init,
		Gen,
	)
}
