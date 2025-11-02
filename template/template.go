package template

import (
	"embed"
	"text/template"
)

//go:embed *.tmpl
var T embed.FS

func Get(path string) (*template.Template, error) {
	return template.ParseFS(T, path)
}
