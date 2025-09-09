package menu_principal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nycholasmarques/rpg-go/internal/game"
	"github.com/nycholasmarques/rpg-go/internal/game/exploration"
)

type menuLoadViewModel struct{
	cursor int
	choices []string
}

func InitialLoadViewMenu() tea.Model {

	return menuLoadViewModel{
		choices: game.PrintSaves(),
	}	
}

func (m menuLoadViewModel) Init() tea.Cmd { return nil }

func (m menuLoadViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				gameState := game.LoadSave(m.choices[m.cursor])
				ebiten.SetWindowSize(640, 480)
				ebiten.SetWindowTitle("Exploration map")
				if err := ebiten.RunGame(exploration.NewEbitenGameExploration(&gameState, exploration.ScreenMenu)); err != nil {
					panic(err)
				}
			case "up":
				if m.cursor > 0 { m.cursor-- }
			case "down":
				if m.cursor < len(m.choices)-1 { m.cursor++ }
		}
	}	
	return m, nil
}

func (m menuLoadViewModel) View() string {
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