package config

///* NOTE: This file contains every migration for every version of the configuration, ever.
// * Please do NOT tamper with versions from prior releases, as this will break
// * our efforts to maintain backward compatability with prior versions of the config.
// *
// * To introduce a new Migration version, simply iterate the StartVersion and EndVersion values by 1
// * so that they may be used for future releases.
// */
//
//var Migrations = []Migration{
//	Migration{
//		StartVersion: 0,
//		EndVersion:   1,
//		Migrations: []MigrationEntity{
//			MigrationEntity{
//				key:         "Spec",
//				transformer: func(data any) any { return 1 },
//				newKey:      "Spec",
//			},
//			MigrationEntity{
//				key: "General.Language",
//				transformer: func(data any) any {
//					return "en-US"
//				},
//				newKey: "General.Language",
//			},
//			MigrationEntity{
//				key: "General.LogLevel",
//				transformer: func(data any) any {
//					return "Info"
//				},
//				newKey: "General.LogLevel",
//			},
//			MigrationEntity{
//				key: "Database.DatabaseLocation",
//				transformer: func(data any) any {
//					return "DEFAULT"
//				},
//				newKey: "Database.DatabaseLocation",
//			},
//			MigrationEntity{
//				key: "LaunchLibrary.LaunchLibraryKey",
//				transformer: func(data any) any {
//					return ""
//				},
//				newKey: "LaunchLibrary.LaunchLibraryKey",
//			},
//			MigrationEntity{
//				key: "LaunchLibrary.CacheToDatabase",
//				transformer: func(data any) any {
//					return true
//				},
//				newKey: "LaunchLibrary.CacheToDatabase",
//			},
//			MigrationEntity{
//				key: "LaunchLibrary.UpdateFrequency",
//				transformer: func(data any) any {
//					return 15
//				},
//				newKey: "LaunchLibrary.UpdateFrequency",
//			},
//			MigrationEntity{
//				key: "Snapi.EnableSnapi",
//				transformer: func(data any) any {
//					return true
//				},
//				newKey: "Snapi.EnableSnapi",
//			},
//			MigrationEntity{
//				key: "Snapi.UpdateFrequency",
//				transformer: func(data any) any {
//					return 15
//				},
//				newKey: "Snapi.UpdateFrequency",
//			},
//			MigrationEntity{
//				key: "ErrorLogging.LogErrors",
//				transformer: func(data any) any {
//					return true
//				},
//				newKey: "ErrorLogging.LogErrors",
//			},
//			MigrationEntity{
//				key: "ErrorLogging.LogFileLocation",
//				transformer: func(data any) any {
//					return "DEFAULT"
//				},
//				newKey: "ErrorLogging.LogFileLocation",
//			},
//			MigrationEntity{
//				key: "ErrorLogging.AutoShareErrors",
//				transformer: func(data any) any {
//					return true
//				},
//				newKey: "ErrorLogging.AutoShareErrors",
//			},
//			MigrationEntity{
//				key: "Telemetry.EnableTelemetry",
//				transformer: func(data any) any {
//					return true
//				},
//				newKey: "Telemetry.EnableTelemetry",
//			},
//			MigrationEntity{
//				key: "Telemetry.TelemetryLevel",
//				transformer: func(data any) any {
//					return 1
//				},
//				newKey: "Telemetry.TelemetryLevel",
//			},
//			MigrationEntity{
//				key: "Keybindings",
//				transformer: func(data any) any {
//					return KeyBindings{
//						"program.quit":     []string{"q"},
//						"program.help":     []string{"h", "?"},
//						"program.launch":   []string{"l"},
//						"program.news":     []string{"n"},
//						"program.settings": []string{"s"},
//						"program.about":    []string{"a"},
//						"program.exit":     []string{"<Esc>"},
//						"program.up":       []string{"k", "<Up>"},
//						"program.down":     []string{"j", "<Down>"},
//						"program.left":     []string{"h", "<Left>"},
//						"program.right":    []string{"l", "<Right>"},
//						"program.pageup":   []string{"<PageUp>"},
//						"program.pagedown": []string{"<PageDown>"},
//						"program.home":     []string{"g", "<Home>"},
//						"program.end":      []string{"G", "<End>"},
//						"program.enter":    []string{"<CR>"},
//						"program.open":     []string{"o", "<Space>"},
//						"news.refresh":     []string{"nr"},
//						"news.open":        []string{"no"},
//						"news.next":        []string{"nn"},
//						"news.previous":    []string{"np"},
//						"news.detailed":    []string{"nd"},
//						"launch.refresh":   []string{"lr"},
//						"launch.open":      []string{"nr"},
//						"launch.next":      []string{"nr"},
//						"launch.previous":  []string{"nr"},
//						"launch.detailed":  []string{"nd"},
//						"settings.refresh": []string{"srl"},
//						"settings.open":    []string{"sof"},
//						"settings.reset":   []string{"sr"},
//					}
//				},
//				newKey: "Keybindings",
//			},
//		},
//	},
//}
