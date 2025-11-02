package pkg

import (
	"context"
	"os"
	"os/exec"
	"syscall"

	"github.com/gowok/gowok/errors"
)

func NewCommandWithCtx(ctx context.Context, cmdName string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, cmdName, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return cmd
}

func NewCommandInDir(dir, cmdName string, args ...string) *exec.Cmd {
	cmd := NewCommandWithCtx(context.Background(), cmdName, args...)
	cmd.Dir = dir
	return cmd
}

func ExecCommandInDir(dir, cmdName string, args ...string) error {
	cmd := NewCommandInDir(dir, cmdName, args...)
	return cmd.Run()
}

func CmdKill(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return errors.New("failed to kill not started process")
	}

	pgid, _ := syscall.Getpgid(cmd.Process.Pid)
	err := syscall.Kill(-pgid, syscall.SIGKILL)
	if err != nil {
		return err
	}

	return nil
}
