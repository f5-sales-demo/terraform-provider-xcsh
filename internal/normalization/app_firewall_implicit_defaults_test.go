package normalization

import (
	"encoding/json"
	"reflect"
	"testing"
)

func implicitFixture() map[string]interface{} {
	return map[string]interface{}{
		"enable_suppression":      map[string]interface{}{},
		"enable_threat_campaigns": map[string]interface{}{},
		"signature_selection_setting": map[string]interface{}{
			"high_medium_accuracy_signatures": map[string]interface{}{},
			"signature_settings_by_accuracy":  map[string]interface{}{"high_accuracy_action": "SIG_BLOCK", "medium_accuracy_action": "SIG_BLOCK", "low_accuracy_action": "SIG_IGNORE"},
		},
		"unrelated": "retained",
	}
}

func TestImplicitDefaults(t *testing.T) {
	input := implicitFixture()
	before, _ := json.Marshal(input)
	got := AppFirewallImplicitDefaults(input, true, true, true)
	if _, present := got["enable_suppression"]; present {
		t.Fatal("implicit suppression default retained")
	}
	if _, present := got["enable_threat_campaigns"]; present {
		t.Fatal("implicit threat default retained")
	}
	selection := got["signature_selection_setting"].(map[string]interface{})
	if _, present := selection["signature_settings_by_accuracy"]; present {
		t.Fatal("implicit action defaults retained")
	}
	if _, present := selection["high_medium_accuracy_signatures"]; !present || got["unrelated"] != "retained" {
		t.Fatal("unrelated policy lost")
	}
	after, _ := json.Marshal(input)
	if string(before) != string(after) {
		t.Fatal("API evidence mutated")
	}
	if !reflect.DeepEqual(AppFirewallImplicitDefaults(input, false, false, false), input) {
		t.Fatal("explicit policy changed")
	}
	for _, which := range []string{"suppression", "threat", "signatures"} {
		got := AppFirewallImplicitDefaults(input, which != "suppression", which != "threat", which != "signatures")
		switch which {
		case "suppression":
			if _, ok := got["enable_suppression"]; !ok {
				t.Fatal("explicit suppression stripped")
			}
		case "threat":
			if _, ok := got["enable_threat_campaigns"]; !ok {
				t.Fatal("explicit threat stripped")
			}
		case "signatures":
			if _, ok := got["signature_selection_setting"].(map[string]interface{})["signature_settings_by_accuracy"]; !ok {
				t.Fatal("explicit actions stripped")
			}
		}
	}
}

func TestImplicitDefaultRejectsNonEquivalentValues(t *testing.T) {
	for _, name := range []string{"high", "medium", "low", "missing", "extra", "malformed", "suppression-content", "threat-content", "suppression-conflict", "threat-conflict"} {
		t.Run(name, func(t *testing.T) {
			input := implicitFixture()
			selection := input["signature_selection_setting"].(map[string]interface{})
			actions := selection["signature_settings_by_accuracy"].(map[string]interface{})
			target := "signature_settings_by_accuracy"
			switch name {
			case "high":
				actions["high_accuracy_action"] = "SIG_IGNORE"
			case "medium":
				actions["medium_accuracy_action"] = "SIG_IGNORE"
			case "low":
				actions["low_accuracy_action"] = "SIG_BLOCK"
			case "missing":
				delete(actions, "medium_accuracy_action")
			case "extra":
				actions["unknown_action"] = "SIG_BLOCK"
			case "malformed":
				actions["high_accuracy_action"] = map[string]interface{}{}
			case "suppression-content":
				target = "enable_suppression"
				input[target] = map[string]interface{}{"unexpected": true}
			case "threat-content":
				target = "enable_threat_campaigns"
				input[target] = map[string]interface{}{"unexpected": true}
			case "suppression-conflict":
				target = "enable_suppression"
				input["disable_suppression"] = map[string]interface{}{}
			case "threat-conflict":
				target = "enable_threat_campaigns"
				input["disable_threat_campaigns"] = map[string]interface{}{}
			}
			before, _ := json.Marshal(input)
			got := AppFirewallImplicitDefaults(input, true, true, true)
			if target == "signature_settings_by_accuracy" {
				if !reflect.DeepEqual(got["signature_selection_setting"].(map[string]interface{})[target], actions) {
					t.Fatal("nondefault actions concealed")
				}
			} else if !reflect.DeepEqual(got[target], input[target]) {
				t.Fatal("non-equivalent marker concealed")
			}
			after, _ := json.Marshal(input)
			if string(before) != string(after) {
				t.Fatal("wire evidence mutated")
			}
		})
	}
}
