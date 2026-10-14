// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestOwnershipTransportJournalsIntentAndReadback(t *testing.T) {
	var stored map[string]interface{}
	var journal *OwnershipJournal
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			if len(journal.Resources) != 1 || journal.Resources[0].Status != "intent" {
				t.Fatal("create was sent before intent")
			}
			if err := json.NewDecoder(r.Body).Decode(&stored); err != nil {
				t.Fatal(err)
			}
			metadata := stored["metadata"].(map[string]interface{})
			metadata["uid"] = "uid-one"
			json.NewEncoder(w).Encode(stored)
		case "GET":
			if stored == nil {
				w.WriteHeader(404)
				w.Write([]byte(`{"code":"NOT_FOUND"}`))
				return
			}
			json.NewEncoder(w).Encode(stored)
		default:
			t.Fatal("unexpected mutation")
		}
	}))
	defer server.Close()
	j, err := CreateOwnershipJournal(filepath.Join(t.TempDir(), "journal.json"), "runone")
	if err != nil {
		t.Fatal(err)
	}
	journal = j
	c := client.NewClient(server.URL, "test")
	c.HTTPClient.Transport = &ownershipTransport{base: http.DefaultTransport, c: c, j: j}
	var result client.OriginPool
	err = c.Post(context.Background(), "/api/config/namespaces/system/origin_pools", map[string]interface{}{"metadata": map[string]interface{}{"name": "tf-acc-test-runone-pool", "namespace": "system"}, "spec": map[string]interface{}{}}, &result)
	if err != nil || len(j.Resources) != 1 || j.Resources[0].Status != "created" || j.Resources[0].UID != ownershipUIDDigest("uid-one") {
		t.Fatalf("journal %#v %v", j.Resources, err)
	}
	// Provider-facing reads omit only the internal ownership marker.
	got, err := c.GetOriginPool(context.Background(), "system", "tf-acc-test-runone-pool")
	if err != nil || got.Metadata.Annotations[ownershipAnnotation] != "" {
		t.Fatal("internal marker leaked into Terraform")
	}
	if _, err = c.CreateOriginPool(context.Background(), &client.OriginPool{Metadata: client.Metadata{Name: "unrelated", Namespace: "system"}}); err == nil {
		t.Fatal("unowned creation allowed")
	}
}

// Ownership prerequisites must bypass provider-facing annotation stripping.
func TestOwnershipReferenceReadContextPreservesRunMarker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"metadata":{"name":"selector","namespace":"system","annotations":{"f5-sales-demo.com/acceptance-run":"runone"}},"system_metadata":{"uid":"immutable"}}`))
	}))
	defer server.Close()
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	c := client.NewClient(server.URL, "synthetic")
	c.HTTPClient.Transport = &ownershipTransport{base: http.DefaultTransport, c: c, j: j}
	ctx := context.WithValue(context.Background(), ownershipReadContextKey{}, true)
	got, err := readOwnershipProjection(ctx, c, "/api/config/namespaces/system/virtual_sites/selector", "system", "selector")
	if err != nil || got.Annotations[ownershipAnnotation] != "runone" {
		t.Fatal("ownership read lost run marker")
	}
	normal, err := c.GetVirtualSite(context.Background(), "system", "selector")
	if err != nil || normal.Metadata.Annotations[ownershipAnnotation] != "" {
		t.Fatal("provider read leaked ownership marker")
	}
}
