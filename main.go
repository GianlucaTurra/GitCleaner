package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	getBranchesWithRemote()
}

func getBranchesWithRemote() {
	cmd := exec.Command("./getBranchesWithRemote.sh")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(output.String())
}
