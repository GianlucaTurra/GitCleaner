package ui

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Quit  key.Binding
	Del   key.Binding
	Push  key.Binding
}

var Keys = keyMap{
	Up:    key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "move up")),
	Down:  key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "move down")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "select")),
	Quit:  key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	Del:   key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete branch")),
	Push:  key.NewBinding(key.WithKeys("p"), key.WithHelp("d", "push branch")),
}

// type CustomHelpDelegate struct{}

func /*(d CustomHelpDelegate)*/ ShortHelp() []key.Binding {
	return []key.Binding{Keys.Push, Keys.Del}
}

// func (d CustomHelpDelegate) FullHelp() [][]key.Binding {
// 	return [][]key.Binding{
// 		{Keys.Up, Keys.Down, Keys.Enter, Keys.Quit},
// 		{Keys.Push, Keys.Del},
// 	}
// }
