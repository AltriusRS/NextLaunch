package config

import "strings"

var DefaultConfig = Configuration{
	Spec: "0",
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
}

type MigrationEntity struct {
	key         string
	transformer func(any) any
	newKey      string
}

func (m MigrationEntity) apply(data any) any {
	return m.transformer(data)
}

type Migration struct {
	StartVersion int
	EndVersion   int
	Migrations   []MigrationEntity
}

func (m Migration) Apply(data map[string]any) map[string]any {
	if m.StartVersion == m.EndVersion {
		return data
	}

	if len(m.Migrations) == 0 {
		return data
	}

	start := data

	for key, val := range start {
		logger.Debugf("Migrating %s from %v to %v", key, val, m.Migrations[0].newKey)
	}

	end := make(map[string]any)

	for _, migration := range m.Migrations {
		end[migration.newKey] = migration.apply(start[migration.key])
	}

	return end
}

func ApplyMigrations(config any) {
	intermediate := config.(map[string]any)
	for _, migration := range Migrations {
		if intermediate["Spec"].(int) == migration.StartVersion && intermediate["Spec"].(int) <= migration.EndVersion {
			intermediate = migration.Apply(intermediate)
		}
	}

	working := DefaultConfig

	// apply migrations to the config variable
	for key, val := range intermediate {
		parts := strings.Split(key, ".")
		levels := len(parts)
		level := levels
		var workspace any
		for level > 0 {
			// if we are at the bottom level, we can just set the value
			if level == levels {
				workspace.(map[string]any)[parts[level]] = val
				level--
				continue
			}

			// if we are not at the bottom level, we should apply the migration to the workspace variable (which should always be up-to-date)
			workspace = working
			level--
		}
		if strings.ContainsRune(key, '.') {
			config.(map[string]any)[key] = val
		}
	}

	config = working
}

var Migrations = []Migration{
	Migration{
		StartVersion: 0,
		EndVersion:   1,
		Migrations: []MigrationEntity{
			MigrationEntity{key: "Spec", transformer: func(data any) any { return 1 }, newKey: "Spec"},
			MigrationEntity{
				key: "General.Language",
				transformer: func(data any) any {
					return "en-US"
				},
				newKey: "General.Language",
			},
			MigrationEntity{
				key: "General.LogLevel",
				transformer: func(data any) any {
					return "Info"
				},
				newKey: "General.LogLevel",
			},
			MigrationEntity{
				key: "Database.DatabaseLocation",
				transformer: func(data any) any {
					return "DEFAULT"
				},
				newKey: "Database.DatabaseLocation",
			},
			MigrationEntity{
				key: "LaunchLibrary.LaunchLibraryKey",
				transformer: func(data any) any {
					return ""
				},
				newKey: "LaunchLibrary.LaunchLibraryKey",
			},
			MigrationEntity{
				key: "LaunchLibrary.CacheToDatabase",
				transformer: func(data any) any {
					return true
				},
				newKey: "LaunchLibrary.CacheToDatabase",
			},
			MigrationEntity{
				key: "LaunchLibrary.UpdateFrequency",
				transformer: func(data any) any {
					return 15
				},
				newKey: "LaunchLibrary.UpdateFrequency",
			},
			MigrationEntity{
				key: "Snapi.EnableSnapi",
				transformer: func(data any) any {
					return true
				},
				newKey: "Snapi.EnableSnapi",
			},
			MigrationEntity{
				key: "Snapi.UpdateFrequency",
				transformer: func(data any) any {
					return 15
				},
				newKey: "Snapi.UpdateFrequency",
			},
			MigrationEntity{
				key: "ErrorLogging.LogErrors",
				transformer: func(data any) any {
					return true
				},
				newKey: "ErrorLogging.LogErrors",
			},
			MigrationEntity{
				key: "ErrorLogging.LogFileLocation",
				transformer: func(data any) any {
					return "DEFAULT"
				},
				newKey: "ErrorLogging.LogFileLocation",
			},
			MigrationEntity{
				key: "ErrorLogging.AutoShareErrors",
				transformer: func(data any) any {
					return true
				},
				newKey: "ErrorLogging.AutoShareErrors",
			},
			MigrationEntity{
				key: "Telemetry.EnableTelemetry",
				transformer: func(data any) any {
					return true
				},
				newKey: "Telemetry.EnableTelemetry",
			},
			MigrationEntity{
				key: "Telemetry.TelemetryLevel",
				transformer: func(data any) any {
					return 1
				},
				newKey: "Telemetry.TelemetryLevel",
			},
		},
	},
}
