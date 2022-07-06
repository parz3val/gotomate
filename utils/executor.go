package utils

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/poggybitz/gotomate/value_objects"
)

func Exectutor(command value_objects.Command) (string, error) {
	// prepapre a command with args
	for _, arg := range command.Args {
		log.Println(arg)
	}
	cmd := exec.Command(command.App, command.Args...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Printf("Couldn't execute command: %s", err)
		log.Fatal(err)
		return "", err
	}
	return string(stdout), nil
}

func Runner(command value_objects.Command) error {
	// prepapre a command with args
	for _, arg := range command.Args {
		log.Println(arg)
	}
	cmd := exec.Command(command.App, command.Args...)
	output, err := cmd.CombinedOutput()
	
	// Ignore errors from the user hitting the cancel button
	if err != nil  {
		log.Printf("Couldn't execute command: %s", err)
		log.Fatal(err)
		return err
	}
fmt.Printf("%s", output)

	return  nil
}


