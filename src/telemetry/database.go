package telemetry

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var featureFlagDatabase *sql.DB
var featureFlagDatabasePath string
var needsInit bool

var featureFlagTableName = "feature_flags"

var initSql = `CREATE TABLE IF NOT EXISTS feature_flags (
	display_name TEXT PRIMARY KEY unique,
	available BOOLEAN not null default true,
	opted_in BOOLEAN not null default false,
	metadata TEXT NOT NULL default '{}',
	groups TEXT
)`

func InitFeatureFlagDatabase() error {
	fmt.Printf("Initializing feature flag database\n")
	// Create the feature flag database if it doesn't already exist
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	featureFlagDatabasePath = filepath.Join(cfgDir, "NextLaunch", "cache.db")

	fmt.Printf("Feature flag database path: %s\n", featureFlagDatabasePath)

	// Check if the feature flag database exists
	_, err = os.Stat(featureFlagDatabasePath)
	if os.IsNotExist(err) {
		fmt.Printf("Feature flag database does not exist, creating\n")
		needsInit = true
	}

	featureFlagDatabase, err = sql.Open("sqlite3", featureFlagDatabasePath)
	if err != nil {
		fmt.Printf("Error opening feature flag database: %v\n", err)
		return err
	}

	defer func(featureFlagDatabase *sql.DB) {
		err := featureFlagDatabase.Close()
		if err != nil {
			panic(err)
		}
	}(featureFlagDatabase)

	if needsInit {
		// Create the feature flag database
		_, err = featureFlagDatabase.Exec(initSql)
		if err != nil {
			fmt.Printf("Error initializing feature flag database: %v\n", err)
			return err
		}
	}

	fmt.Printf("Feature flag database initialized\n")

	return nil
}

func GetFeatureFlagList() []FeatureFlag {
	fmt.Printf("Getting feature flag list\n")
	var flags []FeatureFlag

	rows, err := featureFlagDatabase.Query(
		`
SELECT display_name, available, opted_in, metadata, groups 
FROM feature_flags 
ORDER BY display_name`,
	)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	for rows.Next() {
		var displayName string
		var available bool
		var optedIn bool
		var metadata string
		var group string

		err := rows.Scan(
			&displayName,
			&available,
			&optedIn,
			&metadata,
			&group,
		)
		if err != nil {
			panic(err)
		}

		meta := make(map[string]interface{})
		err = json.Unmarshal([]byte(metadata), &meta)
		if err != nil {
			panic(err)
		}
		flags = append(flags, FeatureFlag{
			Key:       displayName,
			Available: available,
			OptedIn:   optedIn,
			Metadata:  meta,
			Group:     group,
		})
	}

	return flags
}
