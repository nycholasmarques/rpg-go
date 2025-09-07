package menu_principal

import (
	tea "github.com/charmbracelet/bubbletea"
)

type menuModel struct{
	cursor int
	choices []string
}

func InitialMenu() tea.Model {
	return menuModel{
		choices: []string{"New Game", "Load", "Leave"},
	}	
}

func (m menuModel) Init() tea.Cmd { return nil }

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				switch m.cursor {
				case 0:
					return InitialDifficultyMenu(), nil
				case 1:
					return InitialLoadViewMenu(), nil
				case 2:
					return m, tea.Quit
				}
			case "up":
				if m.cursor > 0 { m.cursor-- }
			case "down":
				if m.cursor < len(m.choices)-1 { m.cursor++ }
		}
	}	
	return m, nil
}

func (m menuModel) View() string {
	s := "=== RPG Terminal ===\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "👉"
		}
		s += cursor + " " + choice + "\n"
	}
	return s
}