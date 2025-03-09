package main

import (
	"fmt"
	"os"

	"github.com/GianlucaTurra/GitCleaner/cmd"
	"github.com/GianlucaTurra/GitCleaner/ui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const listHeight = 14

func main() {
	shellOutput := cmd.ReadFromShellScript("./cmd/getLocalNonUpstream.sh")
	ui.Items = cmd.ParseShellOutput(shellOutput)

	const defaultWidth = 20

	l := list.New(ui.Items, ui.ItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Branches"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = ui.TitleStyle
	l.Styles.PaginationStyle = ui.PaginationStyle
	l.Styles.HelpStyle = ui.HelpStyle

	m := ui.Model{List: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
