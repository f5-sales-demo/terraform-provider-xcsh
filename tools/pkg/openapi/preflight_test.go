// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"strings"
	"testing"
)

// parsePreflightsJSON must skip the string "_comment" field (same regression that
// silently disabled import-default-suppressions.json) and load resource entries.
func TestParsePreflightsJSON_SkipsComment(t *testing.T) {
	data := []byte(`{"_comment":"docs here","HTTPLoadBalancer":[{"when_field":"client_side_defense","list_path":"/api/shape/csd/namespaces/%s/protected_domains","requires":"needs a protected_domain","error_title":"missing","error_detail":"none in %s"}]}`)
	got := parsePreflightsJSON(data)
	if _, isComment := got["_comment"]; isComment {
		t.Error("_comment must be skipped, not parsed as a resource")
	}
	pf, ok := got["HTTPLoadBalancer"]
	if !ok || len(pf) != 1 {
		t.Fatalf("want exactly 1 HTTPLoadBalancer preflight, got %#v", got)
	}
	if pf[0].WhenField != "client_side_defense" {
		t.Errorf("WhenField = %q, want client_side_defense", pf[0].WhenField)
	}
	if pf[0].ListPath != "/api/shape/csd/namespaces/%s/protected_domains" {
		t.Errorf("ListPath = %q", pf[0].ListPath)
	}
	if pf[0].ErrorDetail != "none in %s" {
		t.Errorf("ErrorDetail = %q", pf[0].ErrorDetail)
	}
}

// End-to-end: the shipped tools/preflight-requirements.json must yield the
// client_side_defense -> protected_domains preflight for HTTPLoadBalancer,
// proving the file is actually loaded (not just the parser).
func TestLoadPreflights_HTTPLoadBalancerFromFile(t *testing.T) {
	got := LoadPreflights("HTTPLoadBalancer")
	if len(got) == 0 {
		t.Fatal("HTTPLoadBalancer must have at least one preflight from the data file")
	}
	found := false
	for _, p := range got {
		if p.WhenField == "client_side_defense" && strings.Contains(p.ListPath, "protected_domains") {
			found = true
			if n := strings.Count(p.ErrorDetail, "%s"); n != 1 {
				t.Errorf("error_detail must contain exactly one namespace verb, got %d", n)
			}
		}
	}
	if !found {
		t.Errorf("client_side_defense -> protected_domains preflight not loaded; got %#v", got)
	}
}

// A resource with no declared preflights returns an empty slice (no panic, no code emitted).
func TestLoadPreflights_None(t *testing.T) {
	if got := LoadPreflights("NoSuchResourceXYZ"); len(got) != 0 {
		t.Errorf("want no preflights for unknown resource, got %#v", got)
	}
}

// Guard the data file: every declared preflight must be well-formed so the codegen
// emits valid, correctly-argumented Sprintf calls. error_detail and list_path each take
// exactly the namespace, so each must contain exactly one %s; the human-facing fields
// must be non-empty. This prevents the malformed-format class of bug at the source.
func TestLoadPreflights_AllEntriesWellFormed(t *testing.T) {
	preflightOnce.Do(loadPreflights)
	if len(preflightMap) == 0 {
		t.Fatal("preflight-requirements.json loaded no resources")
	}
	for resource, entries := range preflightMap {
		for i, p := range entries {
			if p.WhenField == "" || p.ListPath == "" || p.ErrorTitle == "" || p.ErrorDetail == "" || p.Requires == "" {
				t.Errorf("%s[%d]: required field empty: %+v", resource, i, p)
			}
			if n := strings.Count(p.ListPath, "%s"); n != 1 {
				t.Errorf("%s[%d]: list_path must contain exactly one %%s, got %d", resource, i, n)
			}
			if n := strings.Count(p.ErrorDetail, "%s"); n != 1 {
				t.Errorf("%s[%d]: error_detail must contain exactly one %%s, got %d", resource, i, n)
			}
		}
	}
}
