package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	output string
}

type resultMsg string

func main() {
	p := tea.NewProgram(model{})

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running Bubble Tea program:", err)
		os.Exit(1)
	}
	// branches := strings.Split(executeShellScript("./getAllBranches.sh"), "\n")
	// fmt.Println(filterStringSlice(branches))
	// cmd.Execute()
}

func (m model) Init() tea.Cmd {
	return runShellScript
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case resultMsg:
		m.output = string(msg)
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Command Output:\n\n%s\n\n[Press q to quit]", m.output)
}

func runShellScript() tea.Msg {
	return resultMsg(executeShellScript("./getAllBranches.sh"))
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
