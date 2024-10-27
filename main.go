package main

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"Nextlaunch/src/tsd"
	"strconv"
)

var logger *logging.Logger

var ll2 tsd.LL2Client
var snapi tsd.SnapiClient

func init() {
	logger = logging.NewLogger("main")
	logger.Log("Starting")

	config.LoadConfig()
}

func main() {

	logger.Log("Waiting for initialization to finish")

	logger.Log("Initialized")

	logger.Log("Preparing caches")

	ll2 = *tsd.NewLL2Client()
	snapi = *tsd.NewSnapiClient()

	logger.Log("Starting application")

	launches := ll2.GetLaunches(1, 0)

	if launches == nil {
		logger.Log("No launches found")
		return
	}

	logger.Log("Found " + strconv.Itoa(len(*launches)) + " launches")

	for _, launch := range *launches {
		logger.Log(launch.ID + " - " + launch.Name + " - " + launch.Status.Name)
	}
	//logging.EnterTui()
}
