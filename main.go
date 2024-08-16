package main

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
)

var logger *logging.Logger

//var ll2 tsd.LL2Client
//var snapi tsd.SnapiClient

var initialized chan bool = make(chan bool)

func init() {
	logger = logging.NewLogger("main")
	logger.Log("Starting")

	config.LoadConfig()

	initialized <- true
}

func main() {

	logger.Log("Waiting for initialization to finish")
	<-initialized

	logger.Log("Initialized")

	logger.Log("Preparing caches")

	//ll2 = *tsd.NewLL2Client()
	//snapi = *tsd.NewSnapiClient()
	//
	//logger.Log("Starting application")
	//
	//launches := ll2.GetLaunches(1, 0)
	//
	//if launches == nil {
	//	logger.Log("No launches found")
	//	return
	//}
	//
	//logger.Log("Found " + strconv.Itoa(len(*launches)) + " launches")
	//
	//for _, launch := range *launches {
	//	logger.Log(launch.ID + " - " + launch.Name + " - " + launch.Status.Name)
	//}
	logging.EnterTui()
	logger.Fatal(nil)
}
