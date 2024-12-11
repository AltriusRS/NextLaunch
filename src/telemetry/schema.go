package telemetry

type ApiFeatureFlag struct {
	Name       string   `json:"name"`
	IsOptional bool     `json:"is_optional"`
	Variants   []string `json:"variants"`
}
