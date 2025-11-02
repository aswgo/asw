package service

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/aswgo/asw/pkg"
	"github.com/fsnotify/fsnotify"
	"github.com/gowok/fp/slices"
	"github.com/spf13/cobra"
)

func cmdGoRun(args []string) *exec.Cmd {
	p := pkg.NewCommandInDir(args[0], "go", "run", ".", "--config", "config.toml")
	return p
}

func Run(cmd *cobra.Command, args []string) error {
	slog := slog.With("context", "Run")

	watchedExt := []string{".go", ".toml"}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(args[0])
	if err != nil {
		return err
	}

	filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})

	p := cmdGoRun(args)

	go func() {
		err := p.Start()
		if err != nil {
			fmt.Println(1, err)
			return
		}

		p.Wait()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM)
	signal.Notify(done, syscall.SIGINT)
	for {
		select {
		case event := <-watcher.Events:
			ext := strings.ToLower(filepath.Ext(event.Name))
			if !slices.Contains(watchedExt, ext) {
				continue
			}

			if event.Op != fsnotify.Write {
				continue
			}

			slog.Info("file(s) changed, restarting")

			err = pkg.CmdKill(p)
			if err != nil {
				return err
			}

			p = cmdGoRun(args)
			go func() {
				err := p.Start()
				if err != nil {
					fmt.Println(2, err)
					return
				}
			}()
		case <-done:
			return pkg.CmdKill(p)
		}
	}
}
