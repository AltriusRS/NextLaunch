package tui

import (
	"Nextlaunch/src/config"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/lipgloss"
)

func StartBubbletea(ctx *Model) {
	ctx.Telemetry.Debugf("Spawning bubbletea runtime")

	program := tea.NewProgram(ctx)
	//program := tea.NewProgram(ctx, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Println("could not start program:", err)
		ctx.Telemetry.Errorf("Bubbletea runtime exited with error: %s", err)
		_ = ctx.Telemetry.Trigger("tui.error", 0, map[string]interface{}{
			"version": config.Version,
			"error":   err,
		})
	} else {
		ctx.Telemetry.Debugf("Bubbletea runtime exited")
		_ = ctx.Telemetry.Trigger("tui.close", 0, map[string]interface{}{
			"version": config.Version,
		})
	}
}
