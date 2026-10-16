package normalization

import (
	"reflect"
	"testing"
)

func TestDefaultViolations(t *testing.T) {
	values := make([]interface{}, len(defaultDisabledViolations))
	for i, v := range defaultDisabledViolations {
		values[len(values)-1-i] = v
	}
	original := map[string]interface{}{"violation_settings": map[string]interface{}{"disabled_violation_types": values}, "other": "retained"}
	got := AppFirewallDetectionSettings(original, true)
	if _, ok := got["default_violation_settings"]; !ok {
		t.Fatal("equivalent defaults not normalized")
	}
	if _, ok := got["violation_settings"]; ok {
		t.Fatal("conflicting alternate retained")
	}
	if _, ok := original["violation_settings"]; !ok {
		t.Fatal("API evidence was modified")
	}
	if got["other"] != "retained" {
		t.Fatal("unrelated setting lost")
	}
	if !reflect.DeepEqual(AppFirewallDetectionSettings(original, false), original) {
		t.Fatal("explicit configuration changed")
	}
	for _, name := range []string{"missing", "extra", "duplicate", "changed", "malformed", "extra field"} {
		t.Run(name, func(t *testing.T) {
			v := append([]interface{}{}, values...)
			violations := map[string]interface{}{"disabled_violation_types": v}
			switch name {
			case "missing":
				violations["disabled_violation_types"] = v[1:]
			case "extra":
				violations["disabled_violation_types"] = append(v, "VIOL_OTHER")
			case "duplicate":
				v[0] = v[1]
			case "changed":
				v[0] = "VIOL_OTHER"
			case "malformed":
				v[0] = 123
			case "extra field":
				violations["enabled_violation_types"] = []interface{}{"VIOL_OTHER"}
			}
			input := map[string]interface{}{"violation_settings": violations}
			if !reflect.DeepEqual(AppFirewallDetectionSettings(input, true), input) {
				t.Fatal("real drift normalized away")
			}
		})
	}
}
