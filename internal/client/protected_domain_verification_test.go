package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestProtectedDomainVerification(t *testing.T) {
	cases := []struct {
		name, body         string
		single, collection int
		wantError, absent  bool
	}{
		{"single", "{\"metadata\":{},\"spec\":{\"protected_domain\":\"example.com\"}}", 200, 200, false, false},
		{"single_absent", "", 404, 200, true, true},
		{"deceptive_404_501", "", 500, 200, true, false},
		{"unique", "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\",\"date_added\":\"observed\"}}]}", 501, 200, false, false},
		{"empty_complete", "{\"items\":[]}", 501, 200, true, true},
		{"blank_item", "{\"items\":[{}]}", 501, 200, true, false},
		{"spec_not_get_spec", "{\"items\":[{\"spec\":{\"protected_domain\":\"example.com\"}}]}", 501, 200, true, false},
		{"duplicate", "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\"}},{\"get_spec\":{\"protected_domain\":\"example.com\"}}]}", 501, 200, true, false},
		{"scope", "{\"items\":[{\"namespace\":\"other\",\"get_spec\":{\"protected_domain\":\"example.com\"}}]}", 501, 200, true, false},
		{"identifier", "{\"items\":[{\"name\":\"other\",\"get_spec\":{\"protected_domain\":\"example.com\"}}]}", 501, 200, true, false},
		{"errors", "{\"items\":[],\"errors\":[{\"message\":\"private unrelated information\"}]}", 501, 200, true, false},
		{"unknown_coverage", "{\"items\":[],\"next_page_token\":\"more\"}", 501, 200, true, false},
		{"missing_items", "{}", 501, 200, true, false},
		{"null_items", "{\"items\":null}", 501, 200, true, false},
		{"invalid_json", "{", 501, 200, true, false},
		{"collection_404", "", 501, 404, true, false},
		{"unauthorized", "", 401, 200, true, false},
		{"forbidden", "", 403, 200, true, false},
		{"rate_limit", "", 429, 200, true, false},
		{"unavailable", "", 503, 200, true, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var calls atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				if r.Method != "GET" {
					t.Errorf("verification mutated: %s", r.Method)
				}
				if r.URL.Path == "/api/shape/csd/namespaces/fixture/protected_domains/owned" {
					w.WriteHeader(tc.single)
					if tc.single == 200 {
						fmt.Fprint(w, tc.body)
					} else {
						fmt.Fprint(w, "message containing 404 and 501")
					}
					return
				}
				if r.URL.Path != "/api/shape/csd/namespaces/fixture/protected_domains" {
					t.Errorf("scope: %s", r.URL.Path)
				}
				if !reflect.DeepEqual(r.URL.Query()["report_fields"], []string{"get_spec", "metadata", "name", "namespace", "uid"}) {
					t.Errorf("projection: %v", r.URL.Query())
				}
				w.WriteHeader(tc.collection)
				fmt.Fprint(w, tc.body)
			}))
			defer server.Close()
			c := NewClient(server.URL, "fixture-token", WithMaxRetries(0))
			result, err := c.VerifyProtectedDomain(context.Background(), "fixture", "owned", "example.com")
			if (err != nil) != tc.wantError {
				t.Fatalf("error=%v", err)
			}
			if HasHTTPStatus(err, 404) != tc.absent {
				t.Fatalf("absence=%v error=%v", HasHTTPStatus(err, 404), err)
			}
			expected := int32(1)
			if tc.single == 501 {
				expected = 2
			}
			if calls.Load() != expected {
				t.Fatalf("calls=%d want=%d", calls.Load(), expected)
			}
			if !tc.wantError && result.Spec["protected_domain"] != "example.com" {
				t.Fatal(result)
			}
			if tc.name == "unique" && result.Metadata.Name != "" {
				t.Fatal("invented backend identity")
			}
		})
	}
}

func TestProtectedDomainCollectionBoundary(t *testing.T) {
	for _, count := range []int{255, 256, 257} {
		t.Run(fmt.Sprint(count), func(t *testing.T) {
			items := make([]map[string]interface{}, count)
			for i := range items {
				items[i] = map[string]interface{}{"get_spec": map[string]interface{}{"protected_domain": fmt.Sprintf("fixture-%d.example.com", i)}}
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RawQuery == "" {
					w.WriteHeader(501)
					return
				}
				_ = json.NewEncoder(w).Encode(map[string]interface{}{"items": items})
			}))
			defer server.Close()
			c := NewClient(server.URL, "fixture-token", WithMaxRetries(0))
			_, err := c.VerifyProtectedDomain(context.Background(), "fixture", "owned", "absent.example.com")
			if !HasHTTPStatus(err, 404) {
				t.Fatalf("inferred cap cannot change original List contract: %v", err)
			}
		})
	}
}

func TestProtectedDomainCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c := NewClient("http://127.0.0.1:1", "fixture-token", WithMaxRetries(0))
	_, err := c.VerifyProtectedDomain(ctx, "fixture", "owned", "example.com")
	if err == nil || HasHTTPStatus(err, 404) {
		t.Fatal(err)
	}
}

func TestProtectedDomainCreateResponse(t *testing.T) {
	for _, tc := range []struct {
		name      string
		result    *ProtectedDomain
		wantError bool
	}{
		{"empty", &ProtectedDomain{}, false},
		{"matching", &ProtectedDomain{Spec: map[string]interface{}{"protected_domain": "example.com"}}, false},
		{"different_root", &ProtectedDomain{Spec: map[string]interface{}{"protected_domain": "other.example"}}, true},
		{"different_name", &ProtectedDomain{Metadata: Metadata{Name: "other"}}, true},
		{"different_scope", &ProtectedDomain{Metadata: Metadata{Namespace: "demo-app"}}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateProtectedDomainCreateResponse(tc.result, "fixture", "owned", "example.com")
			if (err != nil) != tc.wantError {
				t.Fatal(err)
			}
		})
	}
}

func TestProtectedDomainSuccessfulMalformedCreateRetainsOwnership(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		calls++
		if req.Method != "POST" {
			t.Error(req.Method)
		}
		fmt.Fprint(w, "{")
	}))
	defer server.Close()
	c := NewClient(server.URL, "fixture-token", WithMaxRetries(3))
	_, landed, err := c.CreateProtectedDomainRegistration(context.Background(), &ProtectedDomain{Metadata: Metadata{Name: "owned", Namespace: "demo"}})
	if err == nil || !landed || calls != 1 {
		t.Fatalf("landed=%v calls=%d error=%v", landed, calls, err)
	}
}

func TestProtectedDomainDomainKeyDelete(t *testing.T) {
	for _, remains := range []bool{false, true} {
		t.Run(fmt.Sprint(remains), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method == "DELETE" {
					if req.URL.Path != "/api/shape/csd/namespaces/demo/protected_domains/example.com" {
						t.Error(req.URL.Path)
					}
					fmt.Fprint(w, "{}")
					return
				}
				if req.URL.RawQuery == "" {
					w.WriteHeader(501)
					return
				}
				if remains {
					fmt.Fprint(w, "{\"items\":[{\"get_spec\":{\"protected_domain\":\"example.com\"}}]}")
				} else {
					fmt.Fprint(w, "{\"items\":[]}")
				}
			}))
			defer server.Close()
			c := NewClient(server.URL, "fixture-token", WithMaxRetries(0))
			err := c.DeleteProtectedDomainRegistration(context.Background(), "demo", "owned", "example.com")
			if (err != nil) != remains {
				t.Fatalf("error=%v", err)
			}
		})
	}
}
