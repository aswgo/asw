package service

import (
	"cmp"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"syscall"

	"github.com/aswgo/asw/pkg"
	"github.com/fsnotify/fsnotify"
	"github.com/gowok/gowok"
	"github.com/spf13/cobra"
)

func cmdGoRunWatch(args []string) *exec.Cmd {
	p := pkg.NewCommandInDir(args[0], "go", "run", ".")
	return p
}

func Run(cmd *cobra.Command, args []string) error {
	flagWatch, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return err
	}

	flagConfig, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	pathConfig := path.Join(cwd, "asw.toml")

	if flagWatch {
		err = runWatch(cmp.Or(flagConfig, pathConfig), args...)
		if err != nil {
			return err
		}
	} else {
		run(cmp.Or(flagConfig, pathConfig))
	}

	return nil
}

func run(config string) {
	gowok.Run(config)
}

func runWatch(config string, args ...string) error {
	watchedExt := []string{".go", ".toml"}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	args = append(args, cwd)

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

	p := cmdGoRunWatch(args)

	go func() {
		err := p.Start()
		if err != nil {
			fmt.Println(1, err)
			return
		}

		p.Wait()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
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

			p = cmdGoRunWatch(args)
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
