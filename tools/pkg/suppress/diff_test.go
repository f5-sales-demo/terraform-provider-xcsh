// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package suppress

import "testing"

// hasMarker reports whether ExtractDefaults found a marker-block default whose leaf
// (final dotted segment) equals name.
func hasMarker(defaults map[string]FieldDefault, name string) bool {
	for _, fd := range defaults {
		if leaf(fd.Path) == name && fd.IsMarkerBlock {
			return true
		}
	}
	return false
}

// #1103: the F5 XC API echoes labels {} on every origin_servers[] element. The
// origin_servers key IS in the request, so the default lives one level down inside a
// LIST element — exactly what the map-only walker skipped. The differ must recurse
// into list elements and record spec.origin_servers.labels as a marker block.
func TestExtractDefaults_OriginServersLabelsMarker_Issue1103(t *testing.T) {
	request := map[string]interface{}{
		"spec": map[string]interface{}{
			"origin_servers": []interface{}{
				map[string]interface{}{
					"public_ip": map[string]interface{}{"ip": "20.98.232.135"},
				},
			},
		},
	}
	response := map[string]interface{}{
		"spec": map[string]interface{}{
			"origin_servers": []interface{}{
				map[string]interface{}{
					"public_ip": map[string]interface{}{"ip": "20.98.232.135"},
					"labels":    map[string]interface{}{},
				},
			},
		},
	}
	got := ExtractDefaults(request, response)
	if !hasMarker(got, "labels") {
		t.Errorf("expected a marker default with leaf 'labels' (origin_servers[].labels {}); got %#v", got)
	}
}

// #1103: default_route_pools[].endpoint_subsets {} — same class, different resource.
func TestExtractDefaults_EndpointSubsetsMarker_Issue1103(t *testing.T) {
	request := map[string]interface{}{
		"spec": map[string]interface{}{
			"default_route_pools": []interface{}{
				map[string]interface{}{
					"pool": map[string]interface{}{"name": "p", "namespace": "ns"},
				},
			},
		},
	}
	response := map[string]interface{}{
		"spec": map[string]interface{}{
			"default_route_pools": []interface{}{
				map[string]interface{}{
					"pool":             map[string]interface{}{"name": "p", "namespace": "ns"},
					"endpoint_subsets": map[string]interface{}{},
				},
			},
		},
	}
	got := ExtractDefaults(request, response)
	if !hasMarker(got, "endpoint_subsets") {
		t.Errorf("expected a marker default with leaf 'endpoint_subsets'; got %#v", got)
	}
}

// A marker nested inside a map inside a list element must also be found (full
// recursion, not just one level into list elements).
func TestExtractDefaults_MarkerNestedInsideListElementMap(t *testing.T) {
	request := map[string]interface{}{
		"spec": map[string]interface{}{
			"routes": []interface{}{
				map[string]interface{}{
					"simple_route": map[string]interface{}{
						"path": map[string]interface{}{"prefix": "/"},
					},
				},
			},
		},
	}
	response := map[string]interface{}{
		"spec": map[string]interface{}{
			"routes": []interface{}{
				map[string]interface{}{
					"simple_route": map[string]interface{}{
						"path":             map[string]interface{}{"prefix": "/"},
						"endpoint_subsets": map[string]interface{}{},
					},
				},
			},
		},
	}
	got := ExtractDefaults(request, response)
	if !hasMarker(got, "endpoint_subsets") {
		t.Errorf("expected marker 'endpoint_subsets' nested in routes[].simple_route; got %#v", got)
	}
}

// Behavior-preservation: a top-level marker block absent from the request (the
// original map-only case) is still discovered after the list-recursion change.
func TestExtractDefaults_TopLevelMarkerStillFound(t *testing.T) {
	request := map[string]interface{}{"spec": map[string]interface{}{"http": map[string]interface{}{"port": float64(80)}}}
	response := map[string]interface{}{"spec": map[string]interface{}{
		"http":        map[string]interface{}{"port": float64(80)},
		"disable_waf": map[string]interface{}{},
	}}
	got := ExtractDefaults(request, response)
	if !hasMarker(got, "disable_waf") {
		t.Errorf("expected top-level marker 'disable_waf'; got %#v", got)
	}
}
