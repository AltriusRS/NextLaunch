package telemetry

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"testing"
)

var (
	phToken string
	phKey   string
)

func init() {
	log.Println("Setting up test suite: Telemetry")

	// Load the environment variables if in development mode
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error loading env file")
		log.Println(err)
	} else {
		log.Println("Loaded env file")
		phToken = os.Getenv("NLPH_TOKEN")
		phKey = os.Getenv("NLPH_KEY")
	}
}

func TestGetDistinctIdentifier(t *testing.T) {
	id, err := GetDistinctIdentifier(nil)
	if err != nil {
		t.Errorf("Error getting distinct identifier: %s", err)
	}

	if id == "" {
		t.Errorf("Distinct identifier is empty")
	} else if len(id) != 36 {
		t.Errorf("Distinct identifier length is not 36")
	} else if !regexp.MustCompile(`^[a-zA-Z0-9\-]+$`).MatchString(id) {
		t.Errorf("Distinct identifier contains invalid characters")
	} else {
		t.Logf("Distinct identifier is valid")
		t.Logf("Distinct identifier: %s", id)
		t.Logf("PHToken: %s", phToken)
		t.Logf("PHKey: %s", phKey)
	}
}

func TestTelemetryInitialization(t *testing.T) {
	testTelemetry, err := NewTelemetry(phToken, phKey, 0, true)

	if err != nil {
		t.Errorf("Error while initializing telemetry client")
		t.Error(err)
	}

	simulatedLevel := testTelemetry.Init()

	if simulatedLevel != 0 {
		t.Errorf("Telemetry event triggered with improper level %d", simulatedLevel)
	}
}

func TestTelemetryFeatureFlagNotFound(t *testing.T) {
	testTelemetry, err := NewTelemetry(phToken, phKey, 0, true)

	if err != nil {
		t.Errorf("Error while initializing telemetry client")
		t.Error(err)
	}

	_, err = testTelemetry.GetFeatureFlag("did_test_feature_flag")

	if err != nil {
		message := err.Error()
		if message != "feature flag 'did_test_feature_flag' not found" {
			t.Errorf("Error message is not correct")
			t.Error(message)
		}
	} else {
		t.Errorf("Error not returned")
	}
}

func TestTelemetryGetFeatureFlag(t *testing.T) {
	testTelemetry, err := NewTelemetry(phToken, phKey, 0, true)

	if err != nil {
		t.Errorf("Error while initializing telemetry client")
		t.Error(err)
	}

	testTelemetry.Init()

	//fmt.Printf("Feature flags: \n%v\n", testTelemetry.GetFeatureFlags())

	flag, err := testTelemetry.GetFeatureFlag("did_test_feature_flag_2")

	if err != nil {
		t.Errorf("Error while getting feature flag")
		t.Error(err)
	}

	t.Logf("Feature flag: %s", flag.Key)
	t.Logf("Feature flag metadata: %v", flag.Metadata)

	if flag.Key != "did_test_feature_flag_2" {
		t.Errorf("Feature flag key is not correct")
	}

	if flag.Available != true {
		t.Errorf("Feature flag is not enabled")
	}

	if flag.Metadata["key"] != "a5b0" {
		t.Errorf("Feature flag metadata key is not correct")
		t.Errorf("Expected: 'a5b0', Got: '%s'", flag.Metadata["key"])
	}
}
