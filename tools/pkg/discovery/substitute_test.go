// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package discovery

import (
	"reflect"
	"testing"
)

func TestSubstitutePlaceholders(t *testing.T) {
	spec := map[string]interface{}{
		"domains": []interface{}{"tf-discover.example.com"},
		"default_route_pools": []interface{}{
			map[string]interface{}{
				"pool": map[string]interface{}{
					"name":      "@prereq:origin_pool",
					"namespace": "@prereq-ns:origin_pool",
				},
				"weight": 1,
			},
		},
		"advertise_on_public_default_vip": map[string]interface{}{},
		"unmatched":                       "@prereq:missing", // no name -> unchanged
	}
	names := map[string]string{"origin_pool": "tf-discover-origin-pool-42"}

	got := SubstitutePlaceholders(spec, names, "tf-discover-ns-1").(map[string]interface{})

	pool := got["default_route_pools"].([]interface{})[0].(map[string]interface{})["pool"].(map[string]interface{})
	if pool["name"] != "tf-discover-origin-pool-42" {
		t.Errorf("pool.name = %v, want created prereq name", pool["name"])
	}
	if pool["namespace"] != "tf-discover-ns-1" {
		t.Errorf("pool.namespace = %v, want test ns", pool["namespace"])
	}
	if got["unmatched"] != "@prereq:missing" {
		t.Errorf("unmatched placeholder should be unchanged, got %v", got["unmatched"])
	}
	// Untouched values preserved.
	if !reflect.DeepEqual(got["domains"], []interface{}{"tf-discover.example.com"}) {
		t.Errorf("domains altered: %v", got["domains"])
	}
}
