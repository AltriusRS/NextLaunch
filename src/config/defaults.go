package config

var DefaultConfig = Configuration{
	Spec: 1,
	General: General{
		Language: "en-US",
		LogLevel: "Info",
	},
	Database: Database{
		DatabaseLocation: "DEFAULT",
	},
	LaunchLibrary: LaunchLibrary{
		LaunchLibraryKey: "",
		CacheToDatabase:  true,
		UpdateFrequency:  15,
	},
	Snapi: Snapi{
		EnableSnapi:     true,
		UpdateFrequency: 15,
	},
	ErrorLogging: ErrorLogging{
		LogErrors:       true,
		LogFileLocation: "DEFAULT",
		AutoShareErrors: true,
	},
	Telemetry: Telemetry{
		EnableTelemetry: true,
		TelemetryLevel:  1,
	},
	/* Keybindings are defined using VIM notations - See https://vimdoc.sourceforge.net/htmldoc/intro.html#key-notation
	 * An action can be defined with multiple keys, but only one action may be defined for each pattern
	 * The order of the keys is important, as it defines the priority of the key
	 * actions are defined in the following format: <namespace>.<action> - <namespace> is the namespace of the action, and <action> is the name of the action
	 * the following namespaces are defined:
	 * program - actions that are related to the program
	 * news - actions that are related to the news
	 * launch - actions that are related to the launches
	 * settings - actions that are related to the settings
	 *
	 * More actions may be added in the future
	 */
	Keybindings: KeyBindings{
		"program.quit":     []string{"q"},
		"program.help":     []string{"h", "?"},
		"program.launch":   []string{"l"},
		"program.news":     []string{"n"},
		"program.settings": []string{"s"},
		"program.about":    []string{"a"},
		"program.exit":     []string{"<Esc>"},
		"program.up":       []string{"k", "<Up>"},
		"program.down":     []string{"j", "<Down>"},
		"program.left":     []string{"h", "<Left>"},
		"program.right":    []string{"l", "<Right>"},
		"program.pageup":   []string{"<PageUp>"},
		"program.pagedown": []string{"<PageDown>"},
		"program.home":     []string{"g", "<Home>"},
		"program.end":      []string{"G", "<End>"},
		"program.enter":    []string{"<CR>"},
		"program.open":     []string{"o", "<Space>"},
		"news.refresh":     []string{"nr"},
		"news.open":        []string{"no"},
		"news.next":        []string{"nn"},
		"news.previous":    []string{"np"},
		"news.detailed":    []string{"nd"},
		"launch.refresh":   []string{"lr"},
		"launch.open":      []string{"nr"},
		"launch.next":      []string{"nr"},
		"launch.previous":  []string{"nr"},
		"launch.detailed":  []string{"nd"},
		"settings.refresh": []string{"srl"},
		"settings.open":    []string{"sof"},
		"settings.reset":   []string{"sr"},
	},
}
