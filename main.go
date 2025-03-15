package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GianlucaTurra/GitCleaner/ui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const listHeight = 14
const defaultWidth = 20

var Items []list.Item

func main() {

	l := list.New(Items, ui.ItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Branches"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = ui.TitleStyle
	l.Styles.PaginationStyle = ui.PaginationStyle
	l.Styles.HelpStyle = ui.HelpStyle

	m := ui.Model{List: l}

	f, err := tea.LogToFile("log.txt", "debug")
	if err != nil {
		log.Println("Fatal :", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
