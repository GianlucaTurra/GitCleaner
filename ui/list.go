package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var Items []list.Item

type Model struct {
	List     list.Model
	Choice   string
	Quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil
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
			m.List.SetItems(pushSelectedBranch(string(i)))
		case "enter":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, tea.Quit
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

func pushSelectedBranch(branchName string) []list.Item {
	var newItems []list.Item
	for _, s := range Items {
		if s == Item(branchName) {
			continue
		}
		newItems = append(newItems, s)
	}
	Items = newItems
	return Items
	// cmd := exec.Command("git", "push", branchName)
	// var output bytes.Buffer
	// cmd.Stdout = &output
	// cmd.Stderr = &output
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%s pushed to remote repo.", branchName)
	// return nil
}
