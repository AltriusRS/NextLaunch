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
		"launch.detailed":  []string{"ld"},
		"launch.next":      []string{"ln"},
		"launch.open":      []string{"lo"},
		"launch.previous":  []string{"lp"},
		"launch.refresh":   []string{"lr"},
		"launch.filter":    []string{"lf"},
		"news.detailed":    []string{"nd"},
		"news.next":        []string{"nn"},
		"news.open":        []string{"no"},
		"news.previous":    []string{"np"},
		"news.refresh":     []string{"nr"},
		"program.about":    []string{"a", "about"},
		"program.down":     []string{"j", "<Down>"},
		"program.end":      []string{"G", "<End>"},
		"program.enter":    []string{"<CR>"},
		"program.exit":     []string{"<Esc>"},
		"program.help":     []string{"?", "help"},
		"program.home":     []string{"g", "<Home>"},
		"program.launch":   []string{"launch"},
		"program.left":     []string{"h", "<Left>"},
		"program.news":     []string{"n", "news"},
		"program.open":     []string{"o", "open", "<Space>"},
		"program.pagedown": []string{"<PageDown>"},
		"program.pageup":   []string{"<PageUp>"},
		"program.quit":     []string{"q", "quit"},
		"program.right":    []string{"l", "<Right>"},
		"program.settings": []string{"s"},
		"program.up":       []string{"k", "<Up>"},
		"settings.open":    []string{"so", "settings"},
		"settings.refresh": []string{"srl", "refresh"},
		"settings.reset":   []string{"sr", "reset"},
	},
}
