// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package suppress

import (
	"reflect"
	"testing"
)

func TestDerive(t *testing.T) {
	db := Database{
		Resources: map[string]*ResourceResult{
			"http_loadbalancer": {
				ResourceName: "http_loadbalancer",
				Status:       "discovered",
				Defaults: map[string]FieldDefault{
					"spec.disable_waf":               {Path: "spec.disable_waf", Type: "object", IsMarkerBlock: true},
					"spec.http.dns_volterra_managed": {Path: "spec.http.dns_volterra_managed", Type: "bool", DefaultValue: false},
					"spec.some_true_flag":            {Path: "spec.some_true_flag", Type: "bool", DefaultValue: true}, // user-meaningful, not suppressed
					"spec.domains":                   {Path: "spec.domains", Type: "array"},                          // not a marker/bool
				},
			},
			"failed_res": {ResourceName: "failed_res", Status: "failed"}, // ignored
		},
	}

	got := Derive(db)
	want := map[string][]string{
		"HTTPLoadBalancer": {"disable_waf", "dns_volterra_managed"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Derive() = %#v, want %#v", got, want)
	}
}

func TestMerge(t *testing.T) {
	existing := map[string][]string{"Healthcheck": {"headers"}, "HTTPLoadBalancer": {"disable_waf"}}
	derived := map[string][]string{"HTTPLoadBalancer": {"round_robin", "disable_waf"}, "AppFirewall": {"disable_x"}}
	got := Merge(existing, derived)
	want := map[string][]string{
		"Healthcheck":      {"headers"},
		"HTTPLoadBalancer": {"disable_waf", "round_robin"},
		"AppFirewall":      {"disable_x"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge() = %#v, want %#v", got, want)
	}
}
