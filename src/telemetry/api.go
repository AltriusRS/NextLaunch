package telemetry

import (
	"encoding/json"
	"fmt"
	"github.com/posthog/posthog-go"
	"io"
	"net/http"
)

func (t *Telemetry) GetFeatureFlags() map[string]*FeatureFlag {
	t.Debugf("Getting feature flags from backend")

	// Use HTTP to get the feature flags from the server
	res, err := http.Get("https://backend.nextlaunch.org/feature_flags")

	if err != nil {
		t.Errorf("Error getting feature flags: %s", err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(res.Body)

	var flags []ApiFeatureFlag

	err = json.NewDecoder(res.Body).Decode(&flags)

	if err != nil {
		t.Errorf("Error decoding feature flags: %s", err)
		return nil
	}

	for _, flag := range flags {
		t.Debugf("Testing for feature flag '%s'", flag.Name)

		flagPayload := posthog.FeatureFlagPayload{
			Key:        flag.Name,
			DistinctId: t.did,
		}

		flagVariant, err := t.client.GetFeatureFlag(flagPayload)

		if err != nil {
			t.Errorf("Error getting feature flag '%s': %s", flag.Name, err)
			continue
		}

		if flagVariant != false {
			payload, err := t.client.GetFeatureFlagPayload(flagPayload)

			if err != nil {
				t.Errorf("Error getting feature flag payload '%s': %s", flag.Name, err)
				continue
			}

			if payload != "" {
				parsedPayload := make(map[string]interface{})
				err := json.Unmarshal([]byte(payload), &parsedPayload)
				if err != nil {
					t.Errorf("Error parsing feature flag payload '%s': %s", payload, err)
					continue
				}

				//t.Debugf("Feature flag '%s' payload: %s", flag.Name, parsedPayload)

				group := ""

				if flagVariant == true {
					group = "enabled"
				} else {
					group = fmt.Sprintf("%s", flagVariant)
				}

				t.featureFlags[flag.Name] = &FeatureFlag{
					Key:       flag.Name,
					Available: true,
					Metadata:  parsedPayload,
					Group:     group,
				}
			}
		} else {
			t.featureFlags[flag.Name] = &FeatureFlag{
				Key:       flag.Name,
				Available: false,
				Metadata: map[string]interface{}{
					"description": "This feature flag is disabled",
				},
				Group: "disabled",
			}
		}
	}

	for k, flag := range t.featureFlags {
		err := flag.LoadOptedIn()
		if err != nil {
			t.Errorf("Error loading opted in status for feature flag '%s': %s", k, err)
		}

		t.featureFlags[k] = flag
	}

	fmt.Printf("Feature flags: %v\n", t.featureFlags)

	return t.featureFlags
}
