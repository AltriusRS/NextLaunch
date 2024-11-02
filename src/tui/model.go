package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var mainStyle = lipgloss.NewStyle().MarginLeft(1).MarginRight(1)

type CursorStyle int

type CursorPosition [2]int

const (
	CursorStyleNone  CursorStyle = iota // Hidden cursor
	CursorStyleBlock                    // Cursor is a one character block
	CursorStyleBeam                     // Cursor is a pane-width beam
	CursorStylePane                     // Cursor is a full pane width and height (highlighting the outline of the pane itself
)

// Model is the state struct for global state management
type Model struct {
	KeybindingManager *KeybindingManager
	CursorPosition    CursorPosition
	CursorBlink       bool
	CursorVisible     bool
	CursorStyle       CursorStyle
	Data              map[string]interface{}
	Frame             Renderer
}
type (
	tickMsg  struct{}
	frameMsg struct{}
)

//func (m *Model) frame() tea.Cmd {
//	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
//		fmt.Println("bubbletea.frame")
//		return frameMsg{}
//	})
//}

func (m *Model) tick() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m *Model) Init() tea.Cmd {
	//return tea.Batch(m.tick(), m.frame())
	return m.tick()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Println("Window size changed")
		fmt.Println(msg.Width)
		fmt.Println(msg.Height)
		m.Frame.SetWidth(msg.Width)
		m.Frame.SetHeight(msg.Height)
		mainStyle.Width(msg.Width)
		mainStyle.Height(msg.Height)
		mainStyle.Render(m.View())
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
		}
	case frameMsg:
		//mainStyle.Render(m.View())
	case tickMsg:
		//mainStyle.Render(m.View())
	}

	//if m.cursorBlink {
	//	m.cursorVisible = !m.cursorVisible
	//}

	return m, m.tick()
}

func (m *Model) View() string {
	return m.Frame.Render(m)
}
