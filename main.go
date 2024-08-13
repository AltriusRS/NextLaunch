package main

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
)

var logger *logging.Logger

func init() {
	logger = logging.NewLogger()
	logger.Log("Starting")

	config.LoadConfig()
}

func main() {
}
