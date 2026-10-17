package provider

import (
	"encoding/json"
	"reflect"
	"testing"
)

const routeOne = "0532376e-5dda-459e-9f3b-4e7be4998699"
const routeTwo = "69e95dbf-3b2f-48fc-95f4-5e1061364bd4"

func testRoute(path, identity string) map[string]interface{} {
	body := map[string]interface{}{"path": map[string]interface{}{"prefix": path}, "http_method": "ANY", "advanced_options": map[string]interface{}{"timeout": float64(10)}}
	if identity != "" {
		body["uuid"] = identity
	}
	return map[string]interface{}{"simple_route": body}
}
func testRouteSpec(routes ...map[string]interface{}) map[string]interface{} {
	items := make([]interface{}, len(routes))
	for i, r := range routes {
		items[i] = r
	}
	return map[string]interface{}{"routes": items}
}
func cloneRouteSpec(t *testing.T, spec map[string]interface{}) map[string]interface{} {
	t.Helper()
	raw, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]interface{}
	if err = json.Unmarshal(raw, &result); err != nil {
		t.Fatal(err)
	}
	return result
}
func TestRouteIdentityPreservesSelectorMatchesAcrossPolicyEditAndReorder(t *testing.T) {
	live := testRouteSpec(testRoute("/one/", routeOne), testRoute("/two/", routeTwo))
	original := cloneRouteSpec(t, live)
	raw, err := encodeRouteIdentities(live, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	desired := testRouteSpec(testRoute("/two/", ""), testRoute("/new/", ""), testRoute("/one/", ""))
	desired["routes"].([]interface{})[2].(map[string]interface{})["simple_route"].(map[string]interface{})["advanced_options"].(map[string]interface{})["timeout"] = float64(20)
	before := cloneRouteSpec(t, desired)
	privateState := map[string][]byte{routeIdentitiesPrivateSlot: raw}
	result, err := preserveRouteIdentities(privateState[routeIdentitiesPrivateSlot], desired, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	routes := result["routes"].([]interface{})
	want := []string{routeTwo, "", routeOne}
	for i, r := range routes {
		id, _ := r.(map[string]interface{})["simple_route"].(map[string]interface{})["uuid"].(string)
		if id != want[i] {
			t.Fatalf("route %d identity %q, want %q", i, id, want[i])
		}
	}
	if !reflect.DeepEqual(live, original) || !reflect.DeepEqual(desired, before) {
		t.Fatal("mutated source API or desired evidence")
	}
	observed := cloneRouteSpec(t, result)
	observed["routes"].([]interface{})[1].(map[string]interface{})["simple_route"].(map[string]interface{})["uuid"] = "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
	if err := verifyPreservedRouteIdentities(result, observed); err != nil {
		t.Fatal(err)
	}
	observed["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{})["uuid"] = "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"
	if err := verifyPreservedRouteIdentities(result, observed); err == nil {
		t.Fatal("server identity replacement accepted")
	}
}
func TestRouteIdentityRejectsInvalidProvenanceAndAmbiguity(t *testing.T) {
	live := testRouteSpec(testRoute("/one/", routeOne))
	raw, err := encodeRouteIdentities(live, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	desired := testRouteSpec(testRoute("/one/", ""))
	for _, tc := range []struct {
		name, ns, lb, token string
		raw                 []byte
	}{
		{"missing", "demo", "lb", "token", nil}, {"malformed", "demo", "lb", "token", []byte("{")}, {"namespace", "other", "lb", "token", raw}, {"name", "demo", "other", "token", raw}, {"token", "demo", "lb", "new-token", raw},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := preserveRouteIdentities(tc.raw, desired, tc.ns, tc.lb, tc.token); err == nil {
				t.Fatal("unverified provenance accepted")
			}
		})
	}
	for _, tc := range []struct {
		name string
		spec map[string]interface{}
	}{
		{"missing-uuid", testRouteSpec(testRoute("/one/", ""))}, {"malformed-uuid", testRouteSpec(testRoute("/one/", "bad"))},
		{"duplicate-selector", testRouteSpec(testRoute("/one/", routeOne), testRoute("/one/", routeTwo))},
		{"duplicate-uuid", testRouteSpec(testRoute("/one/", routeOne), testRoute("/two/", routeOne))},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := encodeRouteIdentities(tc.spec, "demo", "lb", "token"); err == nil {
				t.Fatal("invalid live identity accepted")
			}
		})
	}
	if _, err := preserveRouteIdentities(raw, testRouteSpec(testRoute("/one/", ""), testRoute("/one/", "")), "demo", "lb", "token"); err == nil {
		t.Fatal("ambiguous desired routes accepted")
	}
}
func TestRouteIdentitySelectorDefaultsAndRemovedRoutes(t *testing.T) {
	live := testRouteSpec(testRoute("/one/", routeOne))
	body := live["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{})
	body["headers"] = []interface{}{}
	body["query_params"] = nil
	body["incoming_port"] = nil
	raw, err := encodeRouteIdentities(live, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	desired := testRouteSpec(testRoute("/one/", ""))
	delete(desired["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{}), "http_method")
	result, err := preserveRouteIdentities(raw, desired, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	if result["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{})["uuid"] != routeOne {
		t.Fatal("equivalent selector defaults lost identity")
	}
	result, err = preserveRouteIdentities(raw, map[string]interface{}{}, "demo", "lb", "token")
	if err != nil || len(result) != 0 {
		t.Fatal("removed route restored")
	}
}

func TestRouteIdentityPolicyAndMatchBoundaries(t *testing.T) {
	live := testRouteSpec(testRoute("/one/", routeOne))
	body := live["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{})
	body["query_params"] = map[string]interface{}{"retain_all_params": map[string]interface{}{}}
	body["headers"] = []interface{}{map[string]interface{}{"name": "x-route", "exact": "one", "invert_match": false}}
	raw, err := encodeRouteIdentities(live, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	desired := cloneRouteSpec(t, live)
	body = desired["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{})
	delete(body, "uuid")
	delete(body["headers"].([]interface{})[0].(map[string]interface{}), "invert_match")
	body["query_params"] = map[string]interface{}{"remove_all_params": map[string]interface{}{}}
	result, err := preserveRouteIdentities(raw, desired, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	if result["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{})["uuid"] != routeOne {
		t.Fatal("query rewrite edit lost route identity")
	}
	body["headers"] = []interface{}{map[string]interface{}{"name": "x-route", "exact": "two"}}
	result, err = preserveRouteIdentities(raw, desired, "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := result["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{})["uuid"]; ok {
		t.Fatal("changed match borrowed old UUID")
	}
}

func TestRouteIdentityRejectsCorruptPrivateEvidence(t *testing.T) {
	raw, err := encodeRouteIdentities(testRouteSpec(testRoute("/one/", routeOne)), "demo", "lb", "token")
	if err != nil {
		t.Fatal(err)
	}
	var evidence routeIdentityEvidence
	if err := json.Unmarshal(raw, &evidence); err != nil {
		t.Fatal(err)
	}
	desired := testRouteSpec(testRoute("/one/", ""))
	for _, tc := range []struct {
		name   string
		mutate func(*routeIdentityEvidence)
	}{
		{"version", func(e *routeIdentityEvidence) { e.Version = 2 }},
		{"missing-routes", func(e *routeIdentityEvidence) { e.Routes = nil }},
		{"kind", func(e *routeIdentityEvidence) { e.Routes[0].Kind = "other" }},
		{"selector", func(e *routeIdentityEvidence) { e.Routes[0].Selector = "bad" }},
		{"uuid", func(e *routeIdentityEvidence) { e.Routes[0].UUID = "bad" }},
		{"duplicate", func(e *routeIdentityEvidence) { e.Routes = append(e.Routes, e.Routes[0]) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var copy routeIdentityEvidence
			if err := json.Unmarshal(raw, &copy); err != nil {
				t.Fatal(err)
			}
			tc.mutate(&copy)
			broken, err := json.Marshal(copy)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := preserveRouteIdentities(broken, desired, "demo", "lb", "token"); err == nil {
				t.Fatal("corrupt private evidence accepted")
			}
		})
	}
	for _, broken := range [][]byte{append(append([]byte{}, raw...), []byte(" {}")...), append([]byte(`{"extra":true,`), raw[1:]...)} {
		if _, err := preserveRouteIdentities(broken, desired, "demo", "lb", "token"); err == nil {
			t.Fatal("invalid JSON evidence accepted")
		}
	}
	if _, err := preserveRouteIdentities(raw, testRouteSpec(testRoute("/one/", routeTwo)), "demo", "lb", "token"); err == nil {
		t.Fatal("caller-supplied UUID accepted")
	}
}

func TestRouteIdentityInlineKindsAndReadbackRejection(t *testing.T) {
	for _, kind := range []string{"simple_route"} {
		t.Run(kind, func(t *testing.T) {
			route := testRoute("/one/", routeOne)
			body := route["simple_route"]
			delete(route, "simple_route")
			route[kind] = body
			live := testRouteSpec(route)
			raw, err := encodeRouteIdentities(live, "demo", "lb", "token")
			if err != nil {
				t.Fatal(err)
			}
			desired := cloneRouteSpec(t, live)
			delete(desired["routes"].([]interface{})[0].(map[string]interface{})[kind].(map[string]interface{}), "uuid")
			result, err := preserveRouteIdentities(raw, desired, "demo", "lb", "token")
			if err != nil {
				t.Fatal(err)
			}
			if err := verifyPreservedRouteIdentities(result, live); err != nil {
				t.Fatal(err)
			}
			if err := verifyPreservedRouteIdentities(result, map[string]interface{}{}); err == nil {
				t.Fatal("missing retained route accepted")
			}
			if err := verifyPreservedRouteIdentities(result, desired); err == nil {
				t.Fatal("missing returned UUID accepted")
			}
		})
	}
}

func TestRouteIdentityNonUUIDKindsRemainIntact(t *testing.T) {
	for _, kind := range []string{"redirect_route", "direct_response_route"} {
		t.Run(kind, func(t *testing.T) {
			route := testRoute("/other/", "")
			body := route["simple_route"]
			delete(route, "simple_route")
			route[kind] = body
			spec := testRouteSpec(testRoute("/one/", routeOne), route)
			raw, err := encodeRouteIdentities(spec, "demo", "lb", "token")
			if err != nil {
				t.Fatal(err)
			}
			var evidence routeIdentityEvidence
			if err := json.Unmarshal(raw, &evidence); err != nil {
				t.Fatal(err)
			}
			if len(evidence.Routes) != 1 || evidence.Routes[0].UUID != routeOne {
				t.Fatal("non-UUID route added to identity evidence")
			}
			desired := cloneRouteSpec(t, spec)
			delete(desired["routes"].([]interface{})[0].(map[string]interface{})["simple_route"].(map[string]interface{}), "uuid")
			result, err := preserveRouteIdentities(raw, desired, "demo", "lb", "token")
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(result, spec) {
				t.Fatal("non-UUID route changed")
			}
			if err := verifyPreservedRouteIdentities(result, spec); err != nil {
				t.Fatal(err)
			}
		})
	}
}
