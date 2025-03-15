package ui

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Message type for initializing the view
type branches string

// Message types for input handling
type errorMsg struct{ err error }
type pushMsg struct{ branchName string }
type deleteMsg struct{ branchName string }

// Wrapper functions for executing commands with arguements
func deleteBranch(branchName string) tea.Cmd {
	return func() tea.Msg {
		return deleteMsg{branchName}
	}
}

func pushBranch(branchName string) tea.Cmd {
	return func() tea.Msg {
		return pushMsg{branchName}
	}
}

type Model struct {
	List     list.Model
	Choice   string
	Quitting bool
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		return readBranches()
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil
	case branches:
		m.List.SetItems(parseShellOutput(string(msg)))
		return m, nil
	case deleteMsg:
		items := m.deleteSelectedBranch(msg.branchName)
		m.List.SetItems(items)
	case pushMsg:
		items := m.pushSelectedBranch(msg.branchName)
		m.List.SetItems(items)
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		case "p":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, pushBranch(string(i))
		case "d":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, deleteBranch(string(i))
		case "enter":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.Quitting {
		return QuitTextStyle.Render("Goodbye!")
	}
	return "\n" + m.List.View()
}

func (m Model) pushSelectedBranch(branchName string) []list.Item {
	cmd := exec.Command("/bin/bash", "./cmd/push.sh", branchName)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}
	var newItems []list.Item
	for _, s := range m.List.Items() {
		if s == Item(branchName) {
			continue
		}
		newItems = append(newItems, s)
	}
	return newItems
}

func (m Model) deleteSelectedBranch(branchName string) []list.Item {
	cmd := exec.Command("/bin/bash", "./cmd/delete.sh", branchName)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}
	var newItems []list.Item
	for _, s := range m.List.Items() {
		if s == Item(branchName) {
			continue
		}
		newItems = append(newItems, s)
	}
	return newItems
}

func readBranches() tea.Msg {
	cmd := exec.Command("cmd/getBranches.sh")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}
	return branches(output.String())
}

func parseShellOutput(output string) []list.Item {
	var items []list.Item
	l := strings.Split(output, "\n")
	for _, s := range l {
		if len(strings.TrimSpace(s)) == 0 {
			continue
		}
		items = append(items, Item(s))
	}
	return items
}
