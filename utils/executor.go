package utils

import (
	"os/exec"

	"github.com/poggybitz/gotomate/value_objects"
)

func Exectutor(command value_objects.Command) (string, error) {
	// prepapre a command with args
	cmd := exec.Command(command.App, command.Args...)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}