package tui

import (
	"Nextlaunch/src/config"
	"strings"
)

type Runtime struct {
	Configuration     config.Configuration
	KeybindingManager KeybindingManager
	queue             chan string
}

func NewRuntime(config config.Configuration) *Runtime {
	return &Runtime{
		Configuration:     config,
		KeybindingManager: *NewKeybindManager(),
		queue:             make(chan string, 100),
	}
}

func (r *Runtime) Enqueue(command string) {
	r.queue <- command
}

func (r *Runtime) Run() {
	go func() {
		for command := range r.queue {

			if command == "application.quit" {
				r.KeybindingManager.Stop()
				break
			}

			parts := strings.SplitN(command, ".", 2)

			namespace := parts[0]
			action := parts[1]

			switch namespace {
			case "program":
				r.executeCommand(action)
			case "keybinding":
			}
		}
	}()
}
