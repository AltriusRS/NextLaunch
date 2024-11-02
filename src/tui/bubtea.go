package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/lipgloss"
)

func StartBubbletea(ctx *Model) {

	program := tea.NewProgram(ctx)
	//program := tea.NewProgram(ctx, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

type Renderer interface {
	Render(m *Model) string
	Clear()
	SetWidth(width int)
	SetHeight(height int)
	SetActive(active bool)
	Size() (int, int)
}
