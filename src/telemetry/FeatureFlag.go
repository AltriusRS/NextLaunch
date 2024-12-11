package telemetry

import (
	"database/sql"
	"encoding/json"
)

type FeatureFlag struct {
	Key       string
	Available bool
	OptedIn   bool
	Metadata  map[string]interface{}
	Group     string
}

func (f *FeatureFlag) Save() error {
	metadata, err := json.Marshal(f.Metadata)
	if err != nil {
		return err
	}

	_, err = featureFlagDatabase.Exec(
		`
INSERT INTO feature_flags (display_name, available, opted_in, metadata, groups) 
	VALUES (?, ?, ?, ?, ?) 
	ON CONFLICT(display_name) DO UPDATE 
	    SET available = ?, 
	        opted_in = ?, 
	        metadata = ?, 
			groups = ?`,
		f.Key,
		f.Available,
		f.OptedIn,
		string(metadata),
		f.Group,
		f.Available,
		f.OptedIn,
		string(metadata),
		f.Group,
	)
	if err != nil {
		return err
	}
	return nil
}

func (f *FeatureFlag) Load() error {
	rows, err := featureFlagDatabase.Query(
		`
SELECT display_name, available, opted_in, metadata, groups 
FROM feature_flags 
WHERE display_name = ?`,
		f.Key,
	)
	if err != nil {
		return err
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
			return err
		}

		meta := make(map[string]interface{})
		err = json.Unmarshal([]byte(metadata), &meta)
		if err != nil {
			return err
		}

		f.Key = displayName
		f.Available = available
		f.OptedIn = optedIn
		f.Metadata = meta
		f.Group = group
		return nil
	}
	return nil
}

func (f *FeatureFlag) Delete() error {
	_, err := featureFlagDatabase.Exec(
		`
DELETE FROM feature_flags 
WHERE display_name = ?`,
		f.Key,
	)
	if err != nil {
		return err
	}
	return nil
}

func (f *FeatureFlag) LoadOptedIn() error {
	rows, err := featureFlagDatabase.Query(
		`
SELECT opted_in FROM feature_flags 
WHERE display_name = ?`,
		f.Key,
	)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	for rows.Next() {
		var optedIn bool

		err := rows.Scan(
			&optedIn,
		)
		if err != nil {
			return err
		}

		f.OptedIn = optedIn
		return nil
	}
	return nil
}
