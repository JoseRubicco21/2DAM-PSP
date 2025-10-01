package main

import (
	"fmt"
	"os/exec"
)

func main() {
	lsCmd := exec.Command("ls", "-l")
	lsCmd.Run()
	output, err := lsCmd.Output()
	if err != nil {
		fmt.Println(output)
	} else {
		fmt.Println(err)
	}
}
