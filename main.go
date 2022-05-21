package main

import (
	"fmt"

	"github.com/poggybitz/gotomate/utils"
	"github.com/poggybitz/gotomate/value_objects"
)


func main() {

	command := value_objects.Command {
		App: "ls",
		Args: []string{"-la"},
	}

	output, _ := utils.Exectutor(command)
	fmt.Println(output)

}
