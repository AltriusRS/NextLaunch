package tui

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"fmt"
)

type KeybindingManager struct {
	bindings map[string]string
	logger   *logging.Logger
}

func NewKeybindManager() *KeybindingManager {
	kbm := &KeybindingManager{
		bindings: make(map[string]string),
		logger:   logging.NewLogger("Keybinding Manager"),
	}

	for s := range config.Config.Keybindings {
		for _, v := range config.Config.Keybindings[s] {
			err := kbm.AddBinding(TranslateBinding(v), s)
			if err != nil {
				return nil
			}
		}
	}

	return kbm
}

func TranslateBinding(keys string) string {
	// TODO: Implement VIM notation to bubble tea translations
	return keys
}

func (k *KeybindingManager) AddBinding(key string, action string) error {
	if existingAction, ok := k.bindings[key]; ok {
		k.logger.Errorf("The pattern %s is already bound to action %s, cannot add binding for this %s", key, existingAction, action)
		return fmt.Errorf("the pattern %s is already bound to action %s", key, existingAction)
	}

	k.logger.Debugf("Adding binding %s -> %s", key, action)

	k.bindings[key] = action
	return nil
}

func (k *KeybindingManager) RemoveBinding(key string) error {
	if _, ok := k.bindings[key]; !ok {
		k.logger.Errorf("The pattern %s is not bound to any action", key)
		return fmt.Errorf("the pattern %s is not bound to any action", key)
	}

	delete(k.bindings, key)
	return nil
}

func (k *KeybindingManager) GetBinding(key string) (string, error) {
	if action, ok := k.bindings[key]; ok {
		return action, nil
	}
	return "", fmt.Errorf("the pattern %s is not bound to any action", key)
}

func (k *KeybindingManager) GetBindings() map[string]string {
	return k.bindings
}
