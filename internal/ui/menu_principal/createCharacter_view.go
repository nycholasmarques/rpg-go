package menu_principal

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nycholasmarques/rpg-go/internal/game"
	"github.com/nycholasmarques/rpg-go/internal/game/exploration"
	"github.com/nycholasmarques/rpg-go/internal/game/model"
	"github.com/nycholasmarques/rpg-go/internal/pkg"
)

type createCharacterMenu struct {
	textInput textinput.Model
	class     model.Class
	err       error
}

type (
	errMsg error
)

func InitialCreateCharacterMenu(class model.Class) tea.Model {
	ti := textinput.New()
	ti.Placeholder = "sr golang"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 20
	return createCharacterMenu{
		textInput: ti,
		class:     class,
		err:       nil,
	}
}

func (m createCharacterMenu) Init() tea.Cmd {
	return textinput.Blink
}

func (m createCharacterMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			fileForSave, err := pkg.CreateRandomFile("saves", "save")
			if err != nil {
				panic(err)
			}

			character := model.Character{
				Name:  m.textInput.Value(),
				Class: m.class,
				Level: model.Level_1,
				Hp:    m.class.Hp,
				Xp:    0,
			}
			gameState := model.GameState{
				Character:     character,
				PosX:          0,
				PosY:          0,
				Filename_save: fileForSave,
			}
			game.Save(gameState, "")

			ebiten.SetWindowSize(640, 480)
			ebiten.SetWindowTitle("Exploration map")
			if err := ebiten.RunGame(exploration.NewEbitenGameExploration(&gameState, exploration.ScreenMenu)); err != nil {
				panic(err)
			}
			return m, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m createCharacterMenu) View() string {
	return fmt.Sprintf(
		"Write your character name\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to return to choose class)",
	) + "\n"
}
