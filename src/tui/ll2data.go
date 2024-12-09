package tui

import (
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
		launches := map[string]interface{}{}

		apiLaunches := m.LL2.GetLaunches(20, 0)
		for _, launch := range *apiLaunches {
			launches[launch.ID] = &launch
		}
		cache["launches"] = launches
		cache["last_launch_sync"] = time.Now()
		needsLaunchSync = false
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
