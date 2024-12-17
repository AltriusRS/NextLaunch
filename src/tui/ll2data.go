package tui

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/tui/screens"
	"time"
)

func (m *Model) CheckLL2Data() {
	cache, ok := m.Data["ll2"].(map[string]interface{})

	needsLaunchSync := false
	//needsUpdateSync := false

	if ok {
		lastLaunchSync := cache["last_launch_sync"].(time.Time)
		//lastUpdateSync := cache["last_update_sync"].(time.Time)
		if time.Now().Sub(lastLaunchSync) > time.Minute*15 {
			// Update the launch cache every 15 minutes
			// This ensures that launches are always up to date. In the future, we may make this more dynamic
			// Allowing us to update the cache more frequently the closer we get to the T-0 time
			needsLaunchSync = true
			//needsUpdateSync = true
		}
		//else if time.Now().Sub(lastUpdateSync) > time.Minute*10 {
		//	// Update sync is every 10 minutes, though this will have to change to use the proxy server
		//	needsUpdateSync = true
		//}
	} else {
		needsLaunchSync = true
		//needsUpdateSync = true
		cache = map[string]interface{}{
			"last_launch_sync": time.Now(),
			"last_update_sync": time.Now(),
			"launches":         map[string]interface{}{},
			"updates":          map[string]interface{}{},
		}
	}

	if needsLaunchSync {
		job := screens.NewLoadingState("Syncing launches", time.Second*5)
		job.SetWorking(true).SetProgress(0).SetContext("Querying API")

		m.Compositor.QueueLoadingState(job)
		launches := map[string]interface{}{}

		apiLaunches := m.LL2.GetLaunches(20, 0)

		job.SetProgress(len(*apiLaunches))
		job.SetContext("Processing Response")

		for _, launch := range *apiLaunches {
			launches[launch.ID] = &launch
		}
		job.SetProgress(98)
		job.SetContext(m.Translations.Translate("tasks.update_launches.progress"))
		cache["launches"] = launches
		cache["last_launch_sync"] = time.Now()
		needsLaunchSync = false
		job.SetProgress(100)
		_ = m.Telemetry.Trigger("launches.sync", 0, map[string]interface{}{
			"count":       len(launches),
			"time":        time.Now(),
			"api_version": config.LL2Version,
			"using_token": config.Config.LaunchLibrary.LaunchLibraryKey != "",
		})

		m.Compositor.RemoveLoadingState(job.ID)
	}

	/*if needsUpdateSync {
		var updates []*tsd.LL2LaunchUpdate
		updatedUpdates := m.LL2.GetUpdates(20, 0)
		for _, update := range *updatedUpdates {
			updates = append(updates, &update)
		}
		cache["updates"] = updates
		cache["last_update_sync"] = time.Now()
	}*/

	m.Data["ll2"] = cache
}
