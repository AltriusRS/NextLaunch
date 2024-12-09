package telemetry

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
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

	featureFlags map[string]interface{}
}

func NewTelemetry(phToken string) *Telemetry {
	l := logging.NewLogger("Telemetry")
	if phToken == "" {
		l.Debugf("No Posthog Token Provided")
		return nil
	}

	var client posthog.Client
	var err error

	hasApiKey := false

	if config.PHKey != "unset" {
		hasApiKey = true
		client, err = posthog.NewWithConfig(phToken, posthog.Config{
			Endpoint:       "https://eu.i.posthog.com",
			PersonalApiKey: config.PHKey,
		})
		if err != nil {
			l.Errorf("Failed to create Posthog client: %s", err)
			return nil
		}
	} else {
		l.Debugf("No Posthog Key Provided")
		client, err = posthog.NewWithConfig(phToken, posthog.Config{Endpoint: "https://eu.i.posthog.com"})
		if err != nil {
			l.Errorf("Failed to create Posthog client: %s", err)
			return nil
		}
	}

	did, err := GetDistinctIdentifier(l)
	if err != nil {
		l.Errorf("Failed to get distinct identifier: %s", err)
	}

	l.Debugf("Did: %s", did)
	l.Debugf("PHToken: %s", phToken)

	return &Telemetry{
		Logger:       l,
		PHToken:      phToken,
		did:          did,
		client:       client,
		hasApiKey:    hasApiKey,
		featureFlags: make(map[string]interface{}),
	}
}

func (t *Telemetry) Init() {
	t.Debugf("Initializing Telemetry")

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
		setData["language.selected"] = config.Config.General.Language
		initData["cpu.model"] = CPU.BrandName
		setData["cpu.model"] = CPU.BrandName
		initData["cpu.vendor"] = CPU.VendorString
		setData["cpu.vendor"] = CPU.VendorString
		initData["cpu.cores"] = CPU.PhysicalCores
		setData["cpu.cores"] = CPU.PhysicalCores
		initData["cpu.threads"] = CPU.LogicalCores
		setData["cpu.threads"] = CPU.LogicalCores
		initData["cpu.frequency"] = CPU.Hz
		initData["cpu.cache.l1"] = CPU.Cache.L1I
		initData["cpu.cache.l2"] = CPU.Cache.L2
		initData["cpu.cache.l3"] = CPU.Cache.L3

		for _, feature := range CPU.FeatureSet() {
			initData["cpu.feature."+feature] = true
		}
	}

	initData["$set"] = setData

	t.Trigger("configuration.init", 0, initData)

	//if t.hasApiKey {
	//	t.GetFeatureFlags()
	//}
}

func (t *Telemetry) Trigger(event string, level uint16, properties map[string]interface{}) {
	if !config.Config.Telemetry.EnableTelemetry {
		t.Debugf("Telemetry disabled, not triggering event %s", event)
		return
	}

	if level > config.Config.Telemetry.TelemetryLevel {
		t.Debugf("Telemetry level is %d, not triggering event with level %s", level, event)
		return
	}

	t.Debugf("Triggering event %s", event)

	props := posthog.NewProperties()
	for k, v := range properties {
		props.Set(k, v)
	}

	t.Debugf("Distinct ID: %s", t.did)
	t.Debugf("Event: %s", event)
	t.Debugf("Properties: %v", props)

	err := t.client.Enqueue(posthog.Capture{
		DistinctId: t.did,
		Event:      event,
		Properties: props,
	})

	if err != nil {
		t.Errorf("Error triggering event %s: %s", event, err)
	}
}

func (t *Telemetry) GetDistinctIdentifier() string {
	return t.did
}

//func (t *Telemetry) GetFeatureFlags() map[string]interface{} {
//	flags, err := t.client.GetAllFlags(posthog.FeatureFlagPayloadNoKey{
//		DistinctId: t.did,
//	})
//
//	if err != nil {
//		t.Errorf("Error getting feature flags: %s", err)
//		return nil
//	}
//
//	for _, flag := range flags {
//		println(fmt.Sprintf("%v", flag))
//		//t.featureFlags[flag.Key] = t.GetFeatureFlag(flag.Key)
//	}
//
//	return t.featureFlags
//}

//func (t *Telemetry) GetFeatureFlag(key string) interface{} {
//	payload := posthog.FeatureFlagPayload{
//		Key:        key,
//		DistinctId: t.did,
//	}
//
//	flagVariant, err := t.client.GetFeatureFlag(payload)
//
//	if err != nil {
//		t.Errorf("Error getting feature flag: %s", err)
//		return nil
//	}
//
//	return flagVariant
//}

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
