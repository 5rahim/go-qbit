//go:build windows

package qbittorrent_util

import (
	"context"
	"os/exec"
	"syscall"
)

// NewCmd creates a new exec.Cmd object with the given arguments.
// Since for Windows, the app is built as a GUI application, we need to hide the console windows launched when running commands.
func NewCmd(arg string, args ...string) *exec.Cmd {
	cmd := exec.Command(arg, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: 0x08000000,
	}
	return cmd
}

func NewCmdCtx(ctx context.Context, arg string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, arg, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: 0x08000000,
	}
	return cmd
}
