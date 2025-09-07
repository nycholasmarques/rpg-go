package menu_principal

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycholasmarques/rpg-go/internal/game/model"
)

type classModel struct{
	cursor int
	choices []model.Class
}

func InitialClassMenu() tea.Model {
	return classModel{
		choices: []model.Class{
			{
				Name: "Guerreiro",
				Hp: 200,
				Atk: 15,
				Def: 10,
			},
			{
				Name: "Mago",
				Hp: 100,
				Atk: 25,
				Def: 5,
			},
			{
				Name: "Arqueiro",
				Hp: 150,
				Atk: 20,
				Def: 2,
			},
		},
	}	
}

func (m classModel) Init() tea.Cmd { return nil }

func (m classModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				switch m.cursor {
				case 0:
					return InitialCreateCharacterMenu(m.choices[0]), nil
				case 1:
					return InitialCreateCharacterMenu(m.choices[1]), nil
				case 2:
					return InitialCreateCharacterMenu(m.choices[2]), nil
				case 3:
					return InitialDifficultyMenu(), nil
				}
			case "up":
				if m.cursor > 0 { m.cursor-- }
			case "down":
				if m.cursor < len(m.choices)-1 { m.cursor++ }
		}
	}	
	return m, nil
}

func (m classModel) View() string {
	s := "=== RPG Terminal ===\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "👉"
		}
		s += cursor + " " + choice.Name + ": atk: " + strconv.Itoa(choice.Atk) + " - def: " + strconv.Itoa(choice.Def) + " - hp: " + strconv.Itoa(choice.Hp) + "\n"
	}
	return s
}