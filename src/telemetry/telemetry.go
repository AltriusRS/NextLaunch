package telemetry

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/errors"
	"Nextlaunch/src/logging"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	. "github.com/klauspost/cpuid/v2"
	"github.com/posthog/posthog-go"
	"os"
)

type Telemetry struct {
	*logging.Logger
	PHToken   string
	did       string
	client    posthog.Client
	hasApiKey bool

	featureFlags      map[string]*FeatureFlag
	maxTelemetryLevel uint16
	enabled           bool
	db                *sql.DB
	metadata          posthog.Properties
}

var Testing bool = true

func NewTelemetry(phToken string, phApiKey string, maxTelemetryLevel uint16, enabled bool) (*Telemetry, error) {
	l := logging.NewLogger("Telemetry")
	if phToken == "" {
		l.Debugf("No Posthog Token Provided")
		return nil, errors.NewError(errors.TelemetryTokenNotFound, fmt.Errorf("no posthog token provided"), true)
	}

	var client posthog.Client
	var err error

	hasApiKey := false

	if phApiKey != "" {
		l.Debugf("Using Posthog API Key")
		hasApiKey = true
		client, err = posthog.NewWithConfig(phToken, posthog.Config{
			Endpoint:       "https://eu.i.posthog.com",
			PersonalApiKey: config.PHKey,
		})
		if err != nil {
			l.Errorf("Failed to create Posthog client: %s", err)
			return nil, errors.NewError(errors.TelemetryKeyNotFound, fmt.Errorf("failed to create posthog client"), true)
		}
	} else {
		l.Debugf("No Posthog Key Provided")
		client, err = posthog.NewWithConfig(phToken, posthog.Config{Endpoint: "https://eu.i.posthog.com"})
		if err != nil {
			l.Errorf("Failed to create Posthog client: %s", err)
			return nil, errors.NewError(errors.TelemetryKeyNotFound, fmt.Errorf("failed to create posthog client"), true)
		}
	}

	did, err := GetDistinctIdentifier(l)
	if err != nil {
		l.Errorf("Failed to get distinct identifier: %s", err)
	}

	return &Telemetry{
		Logger:            l,
		PHToken:           phToken,
		maxTelemetryLevel: maxTelemetryLevel,
		did:               did,
		client:            client,
		hasApiKey:         hasApiKey,
		featureFlags:      make(map[string]*FeatureFlag),
		enabled:           enabled,
		metadata:          posthog.NewProperties(),
	}, nil
}

func (t *Telemetry) Init() uint16 {
	err := InitFeatureFlagDatabase()

	if err != nil {
		return 0
	}

	t.Debugf("Initializing Telemetry")

	targetLevel := 0

	if t.maxTelemetryLevel > 0 {
		initData := map[string]interface{}{}

		initData["system.os"] = config.BuildOS
		initData["system.arch"] = config.BuildArch
		initData["ll2.has_api_key"] = config.Config.LaunchLibrary.LaunchLibraryKey != ""
		initData["analytics.enabled"] = config.Config.Telemetry.EnableTelemetry
		initData["analytics.level"] = config.Config.Telemetry.TelemetryLevel
		initData["build.commit"] = config.BuildCommit
		initData["build.date"] = config.BuildDate
		initData["build.os"] = config.BuildOS
		initData["build.arch"] = config.BuildArch
		initData["build.version"] = config.Version

		// Set the data that will be tied to the analytics profile
		setData := map[string]interface{}{}
		setData["system.os"] = config.BuildOS
		setData["system.arch"] = config.BuildArch
		setData["ll2.has_api_key"] = config.Config.LaunchLibrary.LaunchLibraryKey != ""
		setData["analytics.enabled"] = config.Config.Telemetry.EnableTelemetry
		setData["analytics.level"] = config.Config.Telemetry.TelemetryLevel
		setData["build.commit"] = config.BuildCommit
		setData["build.date"] = config.BuildDate
		setData["build.os"] = config.BuildOS
		setData["build.arch"] = config.BuildArch
		setData["build.version"] = config.Version

		if config.Config.Telemetry.TelemetryLevel == 2 {
			initData["language.selected"] = config.Config.General.Language
			initData["cpu.model"] = CPU.BrandName
			initData["cpu.vendor"] = CPU.VendorString
			initData["cpu.cores"] = CPU.PhysicalCores
			initData["cpu.threads"] = CPU.LogicalCores
			initData["cpu.frequency"] = CPU.Hz
			initData["cpu.cache.l1"] = CPU.Cache.L1I
			initData["cpu.cache.l2"] = CPU.Cache.L2
			initData["cpu.cache.l3"] = CPU.Cache.L3

			// Set the data that will be tied to the analytics profile
			setData["language.selected"] = config.Config.General.Language
			setData["cpu.model"] = CPU.BrandName
			setData["cpu.vendor"] = CPU.VendorString
			setData["cpu.cores"] = CPU.PhysicalCores
			setData["cpu.threads"] = CPU.LogicalCores

			for _, feature := range CPU.FeatureSet() {
				initData["cpu.feature."+feature] = true
			}

			targetLevel = 2
		}

		// Configured at compile time, ignore IDE warnings of "always true"
		if Testing {
			setData["test.user"] = true

		}

		// Set the analytics profile data behind the $set key so that
		// posthog can use it to create a new analytics profile (or update the existing one)
		initData["$set"] = setData

		// We can ignore this error, as it is not fatal, and is logged in other places already
		_ = t.Trigger("configuration.init", 0, initData)

		//if t.hasApiKey {
		t.GetFeatureFlags()
		//}
	}

	return uint16(targetLevel)
}

func (t *Telemetry) Trigger(event string, level uint16, properties map[string]interface{}) error {
	if !t.enabled {
		t.Debugf("Telemetry disabled, not triggering event %s", event)
		return errors.NewError(errors.TelemetryDisabled, fmt.Errorf("telemetry disabled, not triggering event %s", event), true)
	}

	if level > t.maxTelemetryLevel {
		t.Debugf("Telemetry level is %d, not triggering event with level %s", level, event)
		return errors.NewError(errors.TelemetryLevelTooLow, fmt.Errorf("telemetry level is %d, not triggering event with level %s", level, event), true)
	}

	t.Debugf("Triggering event %s", event)

	props := posthog.NewProperties()
	for k, v := range properties {
		if event == "configuration.init" && k == "$set" {
			// Process the $set data and map it to the telemetry metadata
			for k2, v2 := range v.(map[string]interface{}) {
				t.metadata.Set(k2, v2)
			}
		}
		props.Set(k, v)
	}

	err := t.client.Enqueue(posthog.Capture{
		DistinctId: t.did,
		Event:      event,
		Properties: props,
	})

	if err != nil {
		t.Errorf("Error triggering event %s: %s", event, err)
	}

	return err
}

func (t *Telemetry) GetDistinctIdentifier() string {
	return t.did
}

func (t *Telemetry) GetFeatureFlag(key string) (*FeatureFlag, error) {
	fmt.Printf("Getting feature flag '%s'\n", key)
	fmt.Printf("Feature flags: %v\n", t.featureFlags)
	flag, ok := t.featureFlags[key]
	if !ok {
		return nil, errors.NewError(errors.TelemetryFeatureFlagNotFound, fmt.Errorf("feature flag '%s' not found", key), true)
	}
	return flag, nil
}

func GetDistinctIdentifier(l *logging.Logger) (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		l.Errorf("Failed to get user config dir: %s", err)
	}

	// Read the distinct id file - or create one if it doesn't exist.

	rawDistinctIdentifier, err := os.ReadFile(cfgDir + "/NextLaunch/.did")
	if err != nil {
		l.Errorf("Failed to read did.txt: %s", err)
	}

	var did string

	if len(rawDistinctIdentifier) == 0 {
		l.Errorf("Failed to read did.txt")
		// If we can't read the did.txt, generate a new one using uuidv7
		u, err := uuid.NewV7()
		if err != nil {
			l.Fatalf("Failed to get UUID: %s", err)
		} else {
			did = u.String()
		}
	} else {
		did = string(rawDistinctIdentifier)
	}

	err = os.WriteFile(cfgDir+"/NextLaunch/.did", []byte(did), 0644)

	if err != nil {
		l.Errorf("Failed to write did.txt: %s", err)
	}

	return did, nil
}
