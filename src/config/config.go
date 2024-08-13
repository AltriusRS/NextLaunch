package config

import (
	"log"
	"os"
	"path/filepath"
)

var Version = "0.0.1"
var LL2Version = "2.2.0"
var LL2BaseURL = "https://ll.thespacedevs.com/"
var SNAPIVersion = "4"
var SNAPIBaseURL = "https://api.spaceflightnewsapi.net/v"
var BuildDate = "UNSET"
var DevBuild = "TRUE"
var IsDev = DevBuild == "TRUE" // This is not a constant because it can be changed at compile time

// Configuration is the main configuration struct
type Configuration struct {
	APIKey string `toml:"api_key"`
}

var Config Configuration = Configuration{}

func LoadConfig() {
	// Prepare the configuration and cahce directories
	PrepConfigDirectory()

	// Load the config file
	//err := toml.DecodeFile("config.toml", &Config)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func PrepConfigDirectory() {
	configDir, err := os.UserConfigDir()

	if err != nil {
		log.Fatal(err)
	}

	configDir = filepath.Join(configDir, "NextLaunch")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.Mkdir(configDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

}
