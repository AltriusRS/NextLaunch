package config

//import (
//	"fmt"
//	"github.com/fatih/structs"
//	"reflect"
//	"strings"
//)
//
//type MigrationEntity struct {
//	key         string
//	transformer func(any) any
//	newKey      string
//}
//
//func (m MigrationEntity) apply(data any) any {
//	return m.transformer(data)
//}
//
//type Migration struct {
//	StartVersion int
//	EndVersion   int
//	Migrations   []MigrationEntity
//}
//
//func (m Migration) Apply(data *map[string]interface{}) {
//	if m.StartVersion == m.EndVersion {
//		return
//	}
//
//	if len(m.Migrations) == 0 {
//		return
//	}
//
//	start := data
//
//	end := structs.New(Configuration{})
//
//	logger.Debugf("Migrating from %v to %v", m.StartVersion, m.EndVersion)
//
//	for _, migration := range m.Migrations {
//		logger.Debugf("Applying migration %v", migration.key)
//		key := migration.key
//		parts := strings.Split(key, ".")
//		levels := len(parts)
//		level := 0
//		l2 := level
//		var breadcrumbs []interface{}
//		var workspace map[string]interface{}
//		for level < levels {
//			logger.Debugf("%s | %d | %d, %v", key, level, levels, workspace)
//			// if we are at the bottom level, we can just set the value
//			if level == levels && workspace != nil {
//				field := workspace[parts[level]]
//
//				workspace[parts[level]] = migration.apply(field)
//
//				// Migrate changes back up the tree to the root of the config
//				for l2 > 0 {
//					temp := workspace
//					workspace = breadcrumbs[level-1].(map[string]interface{})
//					breadcrumbs = breadcrumbs[:level-1]
//					workspace[parts[level]] = temp[parts[level+1]]
//					l2--
//				}
//
//				level++
//				l2 = level
//				continue
//			}
//
//			// if we are not at the bottom level, we should apply the migration to the workspace variable (which should always be up-to-date)
//			logger.Debugf("Setting field %s", parts[level])
//
//			if workspace == nil {
//				workspace = start[parts[level]]
//				breadcrumbs = append(breadcrumbs, workspace)
//			} else {
//				workspace = workspace[parts[level]]
//				breadcrumbs = append(breadcrumbs, workspace)
//			}
//			level++
//		}
//	}
//
//	data = end
//
//	return
//}
//
//// ApplyMigrations applies all migrations to the config
//// It does this by converting the config to a map, via json,
//// and then applying the migrations before converting
//// it back to its original type
//func ApplyMigrations(config *Configuration) {
//
//	intermediate, err := StructToMap(config)
//
//	if err != nil {
//		logger.Fatal(err)
//	}
//
//	if intermediate["Spec"] == nil {
//		intermediate["Spec"] = 0
//	}
//
//	for _, migration := range Migrations {
//		logger.Debugf("Applying migration %+v", migration)
//		logger.Debugf("Versions %+v -> %+v", migration.StartVersion, migration.EndVersion)
//		logger.Debugf("Spec version %v", intermediate["Spec"])
//		logger.Debugf("Migration Applicable? %v", intermediate["Spec"].(int) == migration.StartVersion && specVersion.Value().(int) <= migration.EndVersion)
//
//		if intermediate["Spec"].(int) == migration.StartVersion && intermediate["Spec"].(int) <= migration.EndVersion {
//			migration.Apply(&intermediate)
//		}
//	}
//
//	if intermediate["Spec"].(int) == Migrations[len(Migrations)-1].EndVersion {
//		logger.Debugf("Config is up to date")
//		err = MapToStruct(intermediate, config)
//		if err != nil {
//			logger.Fatal(err)
//		}
//	} else {
//		logger.Debugf("Cofig is out of date")
//		logger.Warningf("Config is out of date, migrations have not been applied successfully")
//	}
//}
//
//const tagName = "nextlaunch"
//
//// StructToMap converts a struct to a map using custom tags
//func StructToMap(s interface{}) (map[string]interface{}, error) {
//	result := make(map[string]interface{})
//	v := reflect.ValueOf(s)
//
//	// Dereference pointer if necessary
//	if v.Kind() == reflect.Ptr {
//		v = v.Elem()
//	}
//
//	if v.Kind() != reflect.Struct {
//		return nil, fmt.Errorf("input is not a struct")
//	}
//
//	t := v.Type()
//	for i := 0; i < v.NumField(); i++ {
//		field := t.Field(i)
//		tag := field.Tag.Get(tagName)
//		if tag == "" {
//			// Skip fields without a tag
//			continue
//		}
//		result[tag] = v.Field(i).Interface()
//	}
//
//	return result, nil
//}
//
//// MapToStruct converts a map to a struct using custom tags
//func MapToStruct(m map[string]interface{}, s interface{}) error {
//	v := reflect.ValueOf(s)
//
//	// Dereference pointer if necessary
//	if v.Kind() == reflect.Ptr {
//		v = v.Elem()
//	}
//
//	if v.Kind() != reflect.Struct {
//		return fmt.Errorf("input is not a pointer to a struct")
//	}
//
//	t := v.Type()
//	for i := 0; i < v.NumField(); i++ {
//		field := t.Field(i)
//		tag := field.Tag.Get(tagName)
//		if tag == "" {
//			// Skip fields without a tag
//			continue
//		}
//
//		mapValue, ok := m[tag]
//		if !ok {
//			// Skip fields not present in the map
//			continue
//		}
//
//		fieldValue := reflect.ValueOf(mapValue)
//		if fieldValue.Type().ConvertibleTo(field.Type) {
//			v.Field(i).Set(fieldValue.Convert(field.Type))
//		} else {
//			return fmt.Errorf("cannot convert map value to struct field %s", field.Name)
//		}
//	}
//
//	return nil
//}
