package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycholasmarques/rpg-go/internal/ui/menu_principal"
)

func main() {
		p := tea.NewProgram(menu_principal.InitialMenu(), tea.WithAltScreen())
		_, err := p.Run()
		if err != nil {
			panic(err)
		}
}