package main

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"Nextlaunch/src/telemetry"
	"Nextlaunch/src/tsd"
	"time"
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

	t, err := telemetry.NewTelemetry(config.PHToken, config.PHKey, config.Config.Telemetry.TelemetryLevel, config.Config.Telemetry.EnableTelemetry)

	if err != nil {
		logger.Errorf("Error while initializing telemetry client")
		logger.Error(err)
	}

	//println(t.GetDistinctIdentifier())

	_ = t.Init()

	//window := widgets.NewWindow("NextLaunch", 80, 20, 1)
	//
	//context := tui.Model{
	//	Telemetry:         t,
	//	KeybindingManager: tui.NewKeybindManager(),
	//	CursorPosition:    tui.CursorPosition{0, 0},
	//	CursorBlink:       false,
	//	CursorVisible:     false,
	//	CursorStyle:       tui.CursorStyleNone,
	//	Data:              make(map[string]interface{}),
	//	LL2:               ll2,
	//	Snapi:             snapi,
	//	Page:              0,
	//	LastPage:          0,
	//	Compositor:        tui.NewCompositor(window),
	//}

	time.Sleep(time.Second * 15)

	//logging.EnterTui()
	//tui.StartBubbletea(&context)
	logging.ShouldExit = true

}
