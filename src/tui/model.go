package tui

import (
	"Nextlaunch/src/telemetry"
	"Nextlaunch/src/translations"
	"Nextlaunch/src/tsd"
	"Nextlaunch/src/tui/screens"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

//var mainStyle = lipgloss.NewStyle().MarginLeft(1).MarginRight(1)

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
	NeedsRepaint      bool
	Telemetry         *telemetry.Telemetry
	KeybindingManager *KeybindingManager
	Translations      *translations.TranslationManager
	CursorPosition    CursorPosition
	CursorBlink       bool
	CursorVisible     bool
	CursorStyle       CursorStyle
	Data              map[string]interface{}
	Page              int
	LastPage          int
	Compositor        *Compositor
	LL2               *tsd.LL2Client
	Snapi             *tsd.SnapiClient
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
	return tea.Tick(time.Second/2, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m *Model) Init() tea.Cmd {
	//return tea.Batch(m.tick(), m.frame())

	screen := screens.LandingScreen(screens.RenderContext{Width: m.Compositor.width, Height: m.Compositor.height})
	screenId := screen.Id()
	m.Compositor.AddWidget(screen)
	m.Compositor.FocusEntity(screenId)

	fmt.Println("Spawning runtime model")

	return m.tick()

}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Println("Window size changed")
		fmt.Println(msg.Width)
		fmt.Println(msg.Height)
		m.Compositor.height = msg.Height
		m.Compositor.width = msg.Width
		print("\x1b[2J")
		m.View()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		//case "ctrl+h", "?":
		//m.KeybindingManager.ShowHelp()
		default:
		}
	case frameMsg:
		//m.View()
	case tickMsg:
		m.CheckLL2Data()
		m.View()
	}

	return m, m.tick()
}

func (m *Model) View() string {
	if !m.NeedsRepaint {
		return ""
	}

	print("\x1b[H" + m.Compositor.Render(m.Compositor.width, m.Compositor.height))
	m.NeedsRepaint = false
	return ""
}
