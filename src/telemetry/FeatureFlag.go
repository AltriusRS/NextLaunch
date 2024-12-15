package telemetry

type FeatureFlag struct {
	Key       string                 `bson:"key"`
	Available bool                   `bson:"available"`
	OptedIn   bool                   `bson:"opted_in"`
	Metadata  map[string]interface{} `bson:"metadata,inline"`
	Group     string                 `bson:"group"`
}

func (f *FeatureFlag) Save() error {
	return nil
}

func (f *FeatureFlag) Load() error {
	//	rows, err := featureFlagDatabase.Query(
	//		`
	//SELECT display_name, available, opted_in, metadata, groups
	//FROM feature_flags
	//WHERE display_name = ?`,
	//		f.Key,
	//	)
	//	if err != nil {
	//		return err
	//	}
	//	defer func(rows *sql.Rows) {
	//		err := rows.Close()
	//		if err != nil {
	//			panic(err)
	//		}
	//	}(rows)
	//
	//	for rows.Next() {
	//		var displayName string
	//		var available bool
	//		var optedIn bool
	//		var metadata string
	//		var group string
	//
	//		err := rows.Scan(
	//			&displayName,
	//			&available,
	//			&optedIn,
	//			&metadata,
	//			&group,
	//		)
	//		if err != nil {
	//			return err
	//		}
	//
	//		meta := make(map[string]interface{})
	//		err = json.Unmarshal([]byte(metadata), &meta)
	//		if err != nil {
	//			return err
	//		}
	//
	//		f.Key = displayName
	//		f.Available = available
	//		f.OptedIn = optedIn
	//		f.Metadata = meta
	//		f.Group = group
	//		return nil
	//	}
	return nil
}

func (f *FeatureFlag) Delete() error {
	//	_, err := featureFlagDatabase.Exec(
	//		`
	//DELETE FROM feature_flags
	//WHERE display_name = ?`,
	//		f.Key,
	//	)
	//	if err != nil {
	//		return err
	//	}
	return nil
}

func (f *FeatureFlag) LoadOptedIn() error {
	//	rows, err := featureFlagDatabase.Query(
	//		`
	//SELECT opted_in FROM feature_flags
	//WHERE display_name = ?`,
	//		f.Key,
	//	)
	//	if err != nil {
	//		return err
	//	}
	//	defer func(rows *sql.Rows) {
	//		err := rows.Close()
	//		if err != nil {
	//			panic(err)
	//		}
	//	}(rows)
	//
	//	for rows.Next() {
	//		var optedIn bool
	//
	//		err := rows.Scan(
	//			&optedIn,
	//		)
	//		if err != nil {
	//			return err
	//		}
	//
	//		f.OptedIn = optedIn
	//		return nil
	//	}
	return nil
}
