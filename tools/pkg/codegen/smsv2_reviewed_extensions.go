package codegen

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// This reviewed projection is shared with the release validator. It includes
// exact mappings and all explicit holds; new sections do not enable surfaces.
// Reviewable input: tests/fixtures/smsv2-contract-7.3-extensions.json.
const smsv2Reviewed73ExtensionSHA256 = "3f28bebbc52ae40b1ab5b1e6b44f6473883f10fa8a73c5d14b2ae154c0782d1a"

func validateSMSv2ReviewedExtensions(raw []byte, version string) error {
	var source struct {
		Providers struct {
			AWS struct {
				Runtime map[string]any `json:"runtime"`
			} `json:"aws"`
			Azure struct {
				Runtime map[string]any `json:"runtime"`
				Relay   map[string]any `json:"slo_bgp_relay"`
			} `json:"azure"`
		} `json:"providers"`
	}
	if err := json.Unmarshal(raw, &source); err != nil {
		return fmt.Errorf("decode reviewed SMSv2 additions: %w", err)
	}
	projection := map[string]any{
		"aws_physical_interfaces": source.Providers.AWS.Runtime["physical_interfaces"],
		"azure_health":            source.Providers.Azure.Runtime["health"],
		"azure_slo_bgp_relay":     source.Providers.Azure.Relay,
	}
	if version == "7.0.0" {
		if projection["aws_physical_interfaces"] != nil || projection["azure_health"] != nil || source.Providers.Azure.Relay != nil {
			return fmt.Errorf("SMSv2 additions require their reviewed version")
		}
		return nil
	}
	canonical, err := json.Marshal(projection)
	if err != nil {
		return fmt.Errorf("encode reviewed SMSv2 additions: %w", err)
	}
	if fmt.Sprintf("%x", sha256.Sum256(canonical)) != smsv2Reviewed73ExtensionSHA256 {
		return fmt.Errorf("SMSv2 additions differ from reviewed mappings and qualification holds")
	}
	return nil
}
