package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSecretPolicyReplaceFormAndExactWire(t *testing.T) {
	calls := []string{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls = append(calls, r.Method+" "+r.URL.Path)
		if !strings.HasPrefix(r.URL.Path, "/api/secret_management/namespaces/demo/secret_policys") {
			t.Error("wrong API route")
		}
		if r.Method == "GET" {
			if r.URL.Query().Get("response_format") != "" && r.URL.Query().Get("response_format") != "2" {
				t.Error("missing replace-form normalization")
			}
			_, _ = w.Write([]byte(`{"metadata":{"name":"policy","namespace":"demo"},"system_metadata":{"uid":"uid-1"},"resource_version":"17","spec":{"algo":"FIRST_MATCH","marked_for_delete":true,"deletion_time":"2030-01-01T00:00:00Z"},"replace_form":{"metadata":{"name":"policy","namespace":"demo","annotations":{"external":"keep"}},"spec":{"algo":"FIRST_MATCH","allow_f5xc":true,"decrypt_cache_timeout":"30s","rule_list":{"rules":[{"metadata":{"name":"allow"},"spec":{"action":"ALLOW","client_selector":{"expressions":["app=demo"]}}},{"metadata":{"name":"deny"},"spec":{"action":"DENY"}}]}}}}`))
			return
		}
		var body map[string]json.RawMessage
		if r.Method != "DELETE" {
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if _, ok := body["system_metadata"]; ok {
				t.Error("readonly identity submitted")
			}
			if strings.Contains(string(body["spec"]), "FIRST_MATCH") || strings.Contains(string(body["spec"]), "marked_for_delete") {
				t.Error("invented or readonly write fields")
			}
		}
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()
	c := NewClient(server.URL, "fixture", WithMaxRetries(0))
	p, err := c.GetSecretPolicy(context.Background(), "demo", "policy")
	if err != nil || p.SystemMetadata.UID != "uid-1" || !p.Spec.MarkedForDelete || len(p.Spec.RuleList.Rules) != 2 {
		t.Fatalf("normalization failed: %v", err)
	}
	p.Spec.MarkedForDelete = false
	p.Spec.DeletionTime = ""
	if err := c.CreateSecretPolicy(context.Background(), p); err != nil {
		t.Fatal(err)
	}
	if err := c.ReplaceSecretPolicy(context.Background(), p); err != nil {
		t.Fatal(err)
	}
	if err := c.DeleteSecretPolicy(context.Background(), "demo", "policy"); err != nil {
		t.Fatal(err)
	}
	if len(calls) != 5 {
		t.Fatal(calls)
	}
}
func TestSecretPolicyLegacyRuleNormalization(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "secret_policy_rules") {
			_, _ = w.Write([]byte(`{"metadata":{"name":"allow","namespace":"demo"},"spec":{"action":"ALLOW","client_selector":{"expressions":["app=demo"]}}}`))
			return
		}
		_, _ = w.Write([]byte(`{"metadata":{"name":"policy","namespace":"demo"},"spec":{"algo":"FIRST_MATCH","allow_f5xc":true,"decrypt_cache_timeout":"30s","legacy_rule_list":{"rules":[{"name":"allow","namespace":"demo"}]}}}`))
	}))
	defer server.Close()
	p, err := NewClient(server.URL, "fixture", WithMaxRetries(0)).GetSecretPolicy(context.Background(), "demo", "policy")
	if err != nil || p.Spec.LegacyRuleList != nil || len(p.Spec.RuleList.Rules) != 1 || p.Spec.RuleList.Rules[0].Spec.ClientSelector.Expressions[0] != "app=demo" {
		t.Fatal("legacy rule normalization failed", err)
	}
}
func TestSecretPolicyWritesNeverReplay(t *testing.T) {
	for _, method := range []string{"POST", "PUT", "DELETE"} {
		t.Run(method, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				http.Error(w, "unavailable", http.StatusServiceUnavailable)
			}))
			defer server.Close()
			c := NewClient(server.URL, "fixture", WithMaxRetries(3))
			p := &SecretPolicy{Metadata: SecretPolicyMetadata{Name: "policy", Namespace: "demo"}}
			var err error
			switch method {
			case "POST":
				err = c.CreateSecretPolicy(context.Background(), p)
			case "PUT":
				err = c.ReplaceSecretPolicy(context.Background(), p)
			case "DELETE":
				err = c.DeleteSecretPolicy(context.Background(), "demo", "policy")
			}
			if err == nil || calls != 1 {
				t.Fatal("write replayed", calls, err)
			}
		})
	}
}
func TestSecretPolicyRejectsForeignIdentity(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"metadata":{"name":"another-policy","namespace":"demo"}}`))
	}))
	defer server.Close()
	if _, err := NewClient(server.URL, "fixture", WithMaxRetries(0)).GetSecretPolicy(context.Background(), "demo", "policy"); err == nil {
		t.Fatal("foreign identity accepted")
	}
}

func TestSecretPolicyLifecycleUsesFullSpecWhenViewSpecIsNull(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("response_format") == "2" {
			_, _ = w.Write([]byte(`{"metadata":{"name":"policy","namespace":"demo"},"spec":null,"replace_form":{"metadata":{"name":"policy","namespace":"demo"},"spec":{"allow_f5xc":true,"decrypt_cache_timeout":"30s"}}}`))
			return
		}
		_, _ = w.Write([]byte(`{"metadata":{"name":"policy","namespace":"demo"},"system_metadata":{"uid":"owned"},"spec":{"algo":"FIRST_MATCH","allow_f5xc":true,"decrypt_cache_timeout":"30s","marked_for_delete":true,"deletion_time":"2030-01-01T00:00:00Z"}}`))
	}))
	defer server.Close()
	p, err := NewClient(server.URL, "fixture", WithMaxRetries(0)).GetSecretPolicy(context.Background(), "demo", "policy")
	if err != nil || !p.Spec.MarkedForDelete || p.Spec.DeletionTime != "2030-01-01T00:00:00Z" || p.Spec.Algorithm != "FIRST_MATCH" {
		t.Fatal("lifecycle fields were lost from the full read", err)
	}
}

func TestSecretPolicyRejectsUnverifiedAlgorithm(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		_, _ = w.Write([]byte(`{"metadata":{"name":"policy","namespace":"demo"},"spec":{"algo":"ALLOW_OVERRIDES"}}`))
	}))
	defer server.Close()
	if _, err := NewClient(server.URL, "fixture", WithMaxRetries(0)).GetSecretPolicy(context.Background(), "demo", "policy"); err == nil || calls != 1 {
		t.Fatal("algorithm was asserted without verified FIRST_MATCH readback")
	}
}
