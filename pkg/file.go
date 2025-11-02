package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aswgo/asw/template"
	"golang.org/x/mod/modfile"
)

func FileCreateFromTemplate(path, tmplName string, data ...any) error {
	fMain, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fMain.Close()

	tmpl, err := template.Get(tmplName)
	if err != nil {
		return err
	}

	var d any
	if len(data) > 0 {
		d = data[0]
	}

	err = tmpl.Execute(fMain, d)
	if err != nil {
		return err
	}

	return nil
}

func FileWriteAfterMarker(target, marker, content string) error {
	f, err := os.Open(target)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	var lines []string
	reader := bufio.NewReader(f)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil && err2 != io.EOF {
			return fmt.Errorf("reading file: %w", err2)
		}
		lines = append(lines, line)
		if err2 == io.EOF {
			break
		}
	}

	found := false
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, line)
		if !found && strings.Contains(line, marker) {
			newLines = append(newLines, content+"\n")
			found = true
		}
	}

	if !found {
		return fmt.Errorf("marker %q not found in file %s", marker, target)
	}

	err = os.WriteFile(target, []byte(strings.Join(newLines, "")), 0644)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func FileGomodRead(target string) (*modfile.Module, error) {
	data, err := os.ReadFile(target)
	if err != nil {
		return nil, fmt.Errorf("error reading go.mod: %v", err)
	}

	mf, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, fmt.Errorf("error parsing go.mod: %v", err)
	}

	if mf.Module == nil {
		return nil, fmt.Errorf("no module directive found in go.mod")
	}

	return mf.Module, nil
}
