package qbittorrent

import (
	"errors"
	"github.com/5rahim/go-qbit/util"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func (c *Client) getBinaryName() string {
	switch runtime.GOOS {
	case "windows":
		return "qbittorrent.exe"
	default:
		return "qbittorrent"
	}
}

func (c *Client) getBinaryPath() string {

	if len(c.BinaryPath) > 0 {
		return c.BinaryPath
	}

	switch runtime.GOOS {
	case "windows":
		return "C:/Program Files/qBittorrent/qbittorrent.exe"
	case "linux":
		return "/usr/bin/qbittorrent"
	case "darwin":
		return "/Applications/Client.app/Contents/MacOS/qBittorrent"
	default:
		return ""
	}
}

func (c *Client) Start() error {
	// If the path is empty, do not check if qBittorrent is running
	if c.BinaryPath == "" {
		return nil
	}

	name := c.getBinaryName()
	if ProgramIsRunning(name) {
		return nil
	}

	path := c.getBinaryPath()

	if path == "" {
		return errors.New("failed to get path to qBittorrent binary")
	}

	cmd := qbittorrent_util.NewCmd(path)
	err := cmd.Start()
	if err != nil {
		return errors.New("failed to start qBittorrent")
	}

	time.Sleep(1 * time.Second)

	return nil
}

func (c *Client) CheckStart() bool {
	if c == nil {
		return false
	}

	// If the path is empty, assume it's running
	if c.BinaryPath == "" {
		return true
	}

	_, err := c.Application.GetAppVersion()
	if err == nil {
		return true
	}

	err = c.Start()
	timeout := time.After(30 * time.Second)
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			_, err = c.Application.GetAppVersion()
			if err == nil {
				return true
			}
		case <-timeout:
			return false
		}
	}
}

func ProgramIsRunning(name string) bool {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = qbittorrent_util.NewCmd("tasklist")
	case "linux":
		cmd = qbittorrent_util.NewCmd("pgrep", name)
	case "darwin":
		cmd = qbittorrent_util.NewCmd("pgrep", name)
	default:
		return true
	}

	output, _ := cmd.Output()

	return strings.Contains(string(output), name)
}
