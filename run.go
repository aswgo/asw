package asw

import (
	"github.com/aswgo/asw/cmd"
	"github.com/gowok/gowok"
)

func Run() {
	cmd.Configure()
	gowok.CMD.Execute()
}
