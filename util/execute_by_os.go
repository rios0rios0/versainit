package util

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
)

func ExecCommand(cmd string) error {
	if runtime.GOOS == "windows" {
		execwindows(cmd)
	} else {
		execlinux(cmd)
	}
	return nil
}

func execlinux(cmd string) error {
	logger.Info("o comando executado é " + cmd)
	command := exec.Command("sh", "-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("exec linux - can't run. Here the reason: %w", err)
	}
	logger.Infof("Command output: %s", string(output))
	return nil
}
func execwindows(command string) error {
	logger.Info("o comando executado é " + command)
	cmd := exec.Command("cmd", "/c", command)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("exec windows - can't run. Here the reason: %w", err)
	}
	logger.Infof("Command output: %s", string(output))
	return nil
}
