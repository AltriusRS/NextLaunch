package tui

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"strings"
)

type CursorStyle int

const (
	CursorStyleNone  CursorStyle = iota // Hidden cursor
	CursorStyleBlock                    // Cursor is a one character block
	CursorStyleBeam                     // Cursor is a pane-width beam
	CursorStylePane                     // Cursor is a full pane width and height (highlighting the outline of the pane itself
)

type Model struct {
	KeybindingManager *KeybindingManager
	queue             *chan string
	display           string
	cursorPosition    struct {
		x int
		y int
	}
	cursorBlink   bool
	cursorVisible bool
	cursorStyle   int
}

type RuntimeCommand struct {
	command string
	data    []interface{}
}

type Runtime struct {
	*logging.Logger
	Configuration     config.Configuration
	KeybindingManager KeybindingManager
	queue             chan RuntimeCommand
}

func NewRuntime(config config.Configuration) *Runtime {
	return &Runtime{
		Logger:            logging.NewLogger("Runtime"),
		Configuration:     config,
		KeybindingManager: *NewKeybindManager(),
		queue:             make(chan RuntimeCommand, 100),
	}
}

func (r *Runtime) Enqueue(command string) {
	r.queue <- RuntimeCommand{command: command}
}

func (r *Runtime) Run() {
	go func() {
		for cmd := range r.queue {
			command := cmd.command
			data := cmd.data

			if command == "application.quit" {
				r.Logger.Info("Quitting application")
				break
			}

			parts := strings.SplitN(command, ".", 2)

			namespace := parts[0]
			action := parts[1]

			switch namespace {
			case "program":
				r.executeCommand(action)
			case "keybinding":
				switch action {
				case "add":
					err := r.KeybindingManager.AddBinding(data[0].(string), data[1].(string))
					if err != nil {
						r.Logger.Error(err)
						return
					}
				case "remove":
					err := r.KeybindingManager.RemoveBinding(data[0].(string))
					if err != nil {
						r.Logger.Error(err)
						return
					}
				case "query":

					binding, err := r.KeybindingManager.GetBinding(data[0].(string))
					if err != nil {
						r.Logger.Error(err)
						return
					}

					r.Logger.Infof("Found binding: %s - %s", data[0].(string), binding)

				case "list":
					r.KeybindingManager.GetBindings()
				}
			}
		}
	}()
}

func (r *Runtime) executeCommand(command string) {
	switch command {
	// TODO: Implement some program commands
	}
}
