package main

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"Nextlaunch/src/tsd"
	"Nextlaunch/src/tui"
	"Nextlaunch/src/tui/widgets"
	"github.com/posthog/posthog-go"
)

var logger *logging.Logger

var ll2 tsd.LL2Client
var snapi tsd.SnapiClient

func init() {
	logger = logging.NewLogger("main")
	logger.Log("Starting")
	//logging.EnterTui()

	config.LoadConfig()
}

func main() {

	logger.Log("Waiting for initialization to finish")

	logger.Log("Initialized")

	logger.Log("Preparing caches")

	ll2 = *tsd.NewLL2Client()
	snapi = *tsd.NewSnapiClient()

	logger.Log("Starting application")

	config.LoadConfig()

	logging.EnterTui()

	window := widgets.NewWindow("NextLaunch", 80, 20, 1)

	ph, err := posthog.NewWithConfig(config.PHToken, posthog.Config{Endpoint: "https://eu.i.posthog.com"})

	if err != nil {
		logger.Errorf("Error initializing posthog")
		logger.Error(err)
	}

	defer func(ph posthog.Client) {
		err := ph.Close()
		if err != nil {
			logger.Errorf("Error closing posthog client")
			logger.Error(err)
		}
	}(ph)

	if config.Config.Telemetry.EnableTelemetry {
		logger.Logf("Telemetry enabled - Starting posthog client")
		initialProps := posthog.NewProperties()

		// Log system information
		initialProps.Set("system.os", config.BuildOS)
		initialProps.Set("system.arch", config.BuildArch)
		//initialProps.Set()

		initialProps.Set("ll2.has_api_key", config.Config.LaunchLibrary.LaunchLibraryKey != "")
		initialProps.Set("analytics.enabled", config.Config.Telemetry.EnableTelemetry)
		initialProps.Set("analytics.level", config.Config.Telemetry.TelemetryLevel)

		// Configure the analytics agent with a startup event trigger
		err = ph.Enqueue(posthog.Capture{
			Event:      "configuration.init",
			Properties: initialProps,
		})

		if err != nil {
			return
		}

	}

	context := tui.Model{
		Analytics:         &ph,
		KeybindingManager: tui.NewKeybindManager(),
		CursorPosition:    tui.CursorPosition{0, 0},
		CursorBlink:       false,
		CursorVisible:     false,
		CursorStyle:       tui.CursorStyleNone,
		Data:              make(map[string]interface{}),
		LL2:               ll2,
		Snapi:             snapi,
		Page:              0,
		LastPage:          0,
		Compositor:        tui.NewCompositor(window),
	}

	tui.StartBubbletea(&context)
	logging.ShouldExit = true
}
