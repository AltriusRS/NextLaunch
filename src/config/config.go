package config

import (
	"Nextlaunch/src/errors"
	"Nextlaunch/src/logging"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path"
	"path/filepath"
	"strconv"
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

var Config Configuration = Configuration{}
var logger *logging.Logger

func LoadConfig() {
	logger = logging.NewLogger("config")
	logger.Log("Loading config")
	// Prepare the configuration directory
	configPath := path.Join(PrepConfigDirectory(), "config.toml")
	logger.Debug("Checking config file at " + configPath)

	stat, err := os.Stat(configPath)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Debug("Checking config file at " + stat.Name())
	logger.Debug("Config file size is " + strconv.Itoa(int(stat.Size())))

	if stat.Size() == 0 {
		logger.Log("Config file is empty, creating a new one")
		Config = DefaultConfig
		WriteConfig(configPath)
		return
	}

	// Load the config file
	file, err := os.OpenFile(configPath, os.O_RDWR, 0644)

	if err != nil {
		logger.Fatal(err)
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

	ApplyMigrations(&Config)
}

// PrepConfigDirectory prepares the config directory for the application
func PrepConfigDirectory() string {
	logger.Log("Preparing config directory")
	configDir, err := os.UserConfigDir()

	if err != nil {
		logger.Fatal(err)
	}

	configDir = filepath.Join(configDir, "NextLaunch")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.Mkdir(configDir, 0644)
		if err != nil {
			logger.Fatal(errors.NewError(errors.ErrorConfigDirectoryNotFound, err, true))
		}
	}

	_, err = os.Create(filepath.Join(configDir, "config.toml"))

	if err != nil {
		logger.Fatal(err)
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
