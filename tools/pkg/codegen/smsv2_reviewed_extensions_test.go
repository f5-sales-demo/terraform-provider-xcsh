package codegen

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func contract73Fixture(t *testing.T) map[string]any {
	t.Helper()
	var contract map[string]any
	if err := json.Unmarshal([]byte(syntheticSMSv2V7Contract), &contract); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join("..", "..", "..", "tests", "fixtures", "smsv2-contract-7.3-extensions.json"))
	if err != nil {
		t.Fatal(err)
	}
	var additions map[string]any
	if err = json.Unmarshal(raw, &additions); err != nil {
		t.Fatal(err)
	}
	providers := contract["providers"].(map[string]any)
	aws := providers["aws"].(map[string]any)
	azure := providers["azure"].(map[string]any)
	aws["runtime"].(map[string]any)["physical_interfaces"] = additions["aws_physical_interfaces"]
	azure["runtime"].(map[string]any)["health"] = additions["azure_health"]
	azure["slo_bgp_relay"] = additions["azure_slo_bgp_relay"]
	contract["version"] = "7.3.0"
	return contract
}

func TestSMSv2Reviewed73ExtensionsPreserveProviderSurfaces(t *testing.T) {
	contract := contract73Fixture(t)
	raw, err := json.Marshal(contract)
	if err != nil {
		t.Fatal(err)
	}
	got, err := SMSv2DataSourceTemplates(raw)
	if err != nil {
		t.Fatal(err)
	}
	baseline, err := SMSv2DataSourceTemplates([]byte(syntheticSMSv2V7Contract))
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(got) != fmt.Sprint(baseline) {
		t.Fatal("held additions changed public provider surfaces")
	}
}

func TestSMSv2Reviewed73ExtensionsRejectPromotion(t *testing.T) {
	for _, kind := range []string{"unknown-version", "downgrade", "missing", "physical-freshness", "health-ha", "relay-live"} {
		t.Run(kind, func(t *testing.T) {
			c := contract73Fixture(t)
			p := c["providers"].(map[string]any)
			aws := p["aws"].(map[string]any)
			azure := p["azure"].(map[string]any)
			switch kind {
			case "unknown-version":
				c["version"] = "7.4.0"
			case "downgrade":
				c["version"] = "7.0.0"
			case "missing":
				delete(aws["runtime"].(map[string]any), "physical_interfaces")
			case "physical-freshness":
				aws["runtime"].(map[string]any)["physical_interfaces"].(map[string]any)["qualification"].(map[string]any)["freshness"] = true
			case "health-ha":
				azure["runtime"].(map[string]any)["health"].(map[string]any)["qualification"].(map[string]any)["ha_health_acceptance"] = true
			case "relay-live":
				azure["slo_bgp_relay"].(map[string]any)["qualification"].(map[string]any)["live_acceptance"] = true
			}
			raw, err := json.Marshal(c)
			if err != nil {
				t.Fatal(err)
			}
			if _, err = SMSv2DataSourceTemplates(raw); err == nil {
				t.Fatal("unreviewed or promoted additions accepted")
			}
		})
	}
}
