// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Trimmed copy of a real registrations_by_site response for a registered
// single-node CE site (ar-bgp-eastus01). Only the fields the data source
// exposes are retained.
const registrationsBySiteOneItem = `{
  "items": [
    {
      "tenant": "example-tenant",
      "namespace": "system",
      "name": "r-dcec2400-52d5-4154-9fd0-4b042d3fe18d",
      "uid": "dcec2400-52d5-4154-9fd0-4b042d3fe18d",
      "system_metadata": {
        "uid": "dcec2400-52d5-4154-9fd0-4b042d3fe18d"
      },
      "object": {
        "status": {
          "current_state": "ONLINE"
        }
      },
      "get_spec": {
        "token": "0d2c8f4a-0000-0000-0000-000000000000",
        "infra": {
          "provider": "AZURE",
          "hostname": "f5-xc-ce-vm-01",
          "instance_id": "/subscriptions/example/resourceGroups/demo/providers/Microsoft.Compute/virtualMachines/f5-xc-ce-vm-01"
        },
        "passport": {
          "cluster_name": "ar-bgp-eastus01",
          "cluster_size": 1
        }
      }
    }
  ],
  "errors": []
}`

// A site whose CE has not registered yet (or does not exist) returns HTTP 200
// with an empty items array — never a 404 and never an error.
const registrationsBySiteEmpty = `{"items":[],"errors":[]}`

func TestListRegistrationsBySite_OneItem(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(registrationsBySiteOneItem))
	}))
	defer server.Close()

	c := NewClient(server.URL, "test-token")

	resp, err := c.ListRegistrationsBySite(context.Background(), "system", "ar-bgp-eastus01")
	if err != nil {
		t.Fatalf("ListRegistrationsBySite() error = %v", err)
	}

	wantPath := "/api/register/namespaces/system/registrations_by_site/ar-bgp-eastus01"
	if gotPath != wantPath {
		t.Errorf("request path = %q, want %q", gotPath, wantPath)
	}

	if len(resp.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(resp.Items))
	}
	item := resp.Items[0]

	if item.Name != "r-dcec2400-52d5-4154-9fd0-4b042d3fe18d" {
		t.Errorf("Name = %q, want %q", item.Name, "r-dcec2400-52d5-4154-9fd0-4b042d3fe18d")
	}
	if item.UID != "dcec2400-52d5-4154-9fd0-4b042d3fe18d" {
		t.Errorf("UID = %q, want %q", item.UID, "dcec2400-52d5-4154-9fd0-4b042d3fe18d")
	}
	if item.SystemMetadata.UID != "dcec2400-52d5-4154-9fd0-4b042d3fe18d" {
		t.Errorf("SystemMetadata.UID = %q, want the item uid", item.SystemMetadata.UID)
	}
	if item.GetSpec.Passport.ClusterName != "ar-bgp-eastus01" {
		t.Errorf("GetSpec.Passport.ClusterName = %q, want %q", item.GetSpec.Passport.ClusterName, "ar-bgp-eastus01")
	}
	if item.GetSpec.Passport.ClusterSize != 1 {
		t.Errorf("GetSpec.Passport.ClusterSize = %d, want 1", item.GetSpec.Passport.ClusterSize)
	}
	if item.GetSpec.Infra.Hostname != "f5-xc-ce-vm-01" {
		t.Errorf("GetSpec.Infra.Hostname = %q, want %q", item.GetSpec.Infra.Hostname, "f5-xc-ce-vm-01")
	}
	if item.GetSpec.Infra.Provider != "AZURE" {
		t.Errorf("GetSpec.Infra.Provider = %q, want %q", item.GetSpec.Infra.Provider, "AZURE")
	}
	if item.GetSpec.Infra.InstanceID != "/subscriptions/example/resourceGroups/demo/providers/Microsoft.Compute/virtualMachines/f5-xc-ce-vm-01" {
		t.Errorf("GetSpec.Infra.InstanceID = %q, want the live infrastructure identity", item.GetSpec.Infra.InstanceID)
	}
	if item.Object.Status.CurrentState != "ONLINE" {
		t.Errorf("Object.Status.CurrentState = %q, want %q", item.Object.Status.CurrentState, "ONLINE")
	}
	if len(resp.Errors) != 0 {
		t.Errorf("len(Errors) = %d, want 0", len(resp.Errors))
	}
}

// An unregistered or unknown site yields HTTP 200 with items: [] — zero items
// and NO error. Callers rely on this to report "not found" without failing.
func TestListRegistrationsBySite_EmptyItems(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(registrationsBySiteEmpty))
	}))
	defer server.Close()

	c := NewClient(server.URL, "test-token")

	resp, err := c.ListRegistrationsBySite(context.Background(), "system", "no-such-site")
	if err != nil {
		t.Fatalf("ListRegistrationsBySite() error = %v, want nil for an unregistered site", err)
	}
	if len(resp.Items) != 0 {
		t.Errorf("len(Items) = %d, want 0", len(resp.Items))
	}
}

// Registration objects can carry raw control bytes inside embedded string
// values (certificates, logs). The client must sanitize them (GetLenient)
// rather than fail the whole read.
func TestListRegistrationsBySite_ToleratesRawControlChars(t *testing.T) {
	body := "{\"items\":[{\"name\":\"r-1\",\"get_spec\":{\"infra\":{\"hostname\":\"a\x01b\"},\"passport\":{\"cluster_name\":\"s1\"}}}],\"errors\":[]}"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()

	c := NewClient(server.URL, "test-token")

	resp, err := c.ListRegistrationsBySite(context.Background(), "system", "s1")
	if err != nil {
		t.Fatalf("ListRegistrationsBySite() error = %v", err)
	}
	if len(resp.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(resp.Items))
	}
	if got := resp.Items[0].GetSpec.Infra.Hostname; got != "ab" {
		t.Errorf("Hostname = %q, want %q (control char stripped)", got, "ab")
	}
}
