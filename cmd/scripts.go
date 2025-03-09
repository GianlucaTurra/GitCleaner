package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	"github.com/GianlucaTurra/GitCleaner/ui"
	"github.com/charmbracelet/bubbles/list"
)

func ReadFromShellScript(scriptName string) string {
	cmd := exec.Command(scriptName)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return output.String()
}

func ParseShellOutput(strOutput string) []list.Item {
	var items []list.Item
	l := strings.Split(strOutput, "\n")
	for _, s := range l {
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}
		items = append(items, ui.Item(s))
	}
	return items
}
