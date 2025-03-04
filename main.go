package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	branches := strings.Split(executeShellScript("./getAllBranches.sh"), "\n")
	fmt.Println(filterStringSlice(branches))
}

func executeShellScript(scriptName string) string {
	cmd := exec.Command(scriptName)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return output.String()
}

func filterStringSlice(slice []string) []string {
	var newSlice []string
	for _, s := range slice {
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}
		newSlice = append(newSlice, s)
	}
	return newSlice
}
