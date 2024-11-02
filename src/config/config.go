package config

import (
	"Nextlaunch/src/errors"
	"Nextlaunch/src/logging"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path"
	"path/filepath"
)

var Version = "0.0.1"
var LL2Version = "2.3.0"
var LL2BaseURL = "https://lldev.thespacedevs.com/"
var LL2FullBaseURL = LL2BaseURL + LL2Version + "/"
var SNAPIVersion = "4"
var SNAPIBaseURL = "https://api.spaceflightnewsapi.net/v"
var SNAPIFullBaseURL = SNAPIBaseURL + SNAPIVersion + "/"
var BuildDate = "unset"
var DevBuild = "true"

//goland:noinspection GoBoolExpressions
var IsDev = DevBuild == "true" // This is not a constant because it can be changed at compile time

var Config Configuration
var logger *logging.Logger

func LoadConfig() {
	logger = logging.NewLogger("config")
	logger.Log("Loading config")
	// Prepare the configuration directory
	configPath, err := filepath.Abs(path.Join(PrepConfigDirectory(), "config.toml"))

	if err != nil {
		logger.Fatal(err)
	}

	logger.Debug("Checking config file at " + configPath)

	var file *os.File

	stat, err := os.Stat(configPath)

	if err != nil {
		logger.Fatal(err)
	}

	if os.IsNotExist(err) {
		logger.Debug("Creating config file")
		file, err = os.Create(configPath)
		if err != nil {
			logger.Error(err)
		}
	} else if err != nil {
		logger.Error(err)
		file, err = os.Open(configPath)
		if err != nil {
			logger.Error(err)
		}
	} else {
		logger.Debug("Loading config file")
		file, err = os.Open(configPath)
		if err != nil {
			logger.Error(err)
		}
	}

	logger.Debug("Checking config file at " + stat.Name())
	logger.Debugf("Config file size is %d", stat.Size())

	if stat.Size() == 0 {
		logger.Log("Config file is empty, creating a new one")
		Config = DefaultConfig
		WriteConfig(configPath)
		return
	}

	// Close the file once we're done with it
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(file)

	// Decode the file into the Config struct
	content := toml.NewDecoder(file)

	err = content.Decode(&Config)

	if err != nil {
		logger.Fatal(err)
	}

	if Config.Spec == 0 {
		logger.Debug("Config file is empty")
		Config = DefaultConfig
		WriteConfig(configPath)
	}

	logger.Log("Config loaded")
	logger.Debugf("Config is version %d", Config.Spec)

	//ApplyMigrations(&Config)

	//structMap, err := StructToMap(DefaultConfig, "")

	//if err != nil {
	//	logger.Fatal(err)
	//}
	//
	//for k, v := range structMap {
	//	logger.Debugf("%s: %v", k, v)
	//}
}

/*
func StructToMap(s interface{}, prefix string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	t := v.Type()

	if prefix == "" {
		prefix = "root"
	}
	prefix = prefix + "."

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("nextlaunch")
		if tag == "" {
			// Skip fields without a tag
			continue
		}

		if typeOf := v.FieldByIndex([]int{i}).Type(); typeOf.Kind() == reflect.Struct {
			r, err := StructToMap(v.FieldByIndex([]int{i}), prefix+tag)
			if err != nil {
				return nil, err
			}
			for k, v := range r {
				result[k] = v
			}
			continue
		} /* else if typeOf := v.FieldByIndex([]int{i}).Type(); typeOf.Kind() == reflect.Map {
			v2 := v.FieldByIndex([]int{i}).Interface()
			v3 := reflect.ValueOf(v2)

			if v3.Kind() != reflect.Map {
				return nil, fmt.Errorf("input is not a map")
			}

			for k, v := range v3.Interface().(map[string]interface{}) {
				result[prefix+tag+"."+k] = v
			}
			continue
		}*\/

		result[prefix+tag] = v.FieldByIndex([]int{i}).Interface()
	}

	return result, nil
}
*/

// PrepConfigDirectory prepares the config directory for the application
func PrepConfigDirectory() string {
	logger.Log("Preparing config directory")
	configDir, err := os.UserConfigDir()
	os.TempDir()

	if err != nil {
		logger.Fatal(err)
	}

	configDir = filepath.Join(configDir, "NextLaunch")

	var file *os.File

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.Mkdir(configDir, 0644)
		if err != nil {
			logger.Fatal(errors.NewError(errors.ErrorConfigDirectoryNotFound, err, true))
		}
		file, err = os.Create(filepath.Join(configDir, "config.toml"))
	} else if err != nil {
		logger.Fatal(err)
	}

	if err != nil {
		logger.Fatal(err)
	}

	if file == nil {
		file, err = os.Open(filepath.Join(configDir, "config.toml"))
	}

	if err != nil {
		logger.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		logger.Fatal(err)
		return ""
	}

	return configDir
}

// WriteConfig writes the current config to the config file
func WriteConfig(dir string) {
	logger.Debug("Creating config file at " + dir)
	logger.Log("Creating config file")
	file, err := os.Create(dir)

	if err != nil {
		logger.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(file)

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(Config)

	if err != nil {
		logger.Fatal(err)
	}
}
