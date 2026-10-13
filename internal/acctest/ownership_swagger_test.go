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

func TestOwnedSwaggerUsesExactVersionAndDigest(t *testing.T) {
	content := `{"openapi":"3.0.3","info":{"title":"Fixture","version":"1"},"paths":{}}`
	present := true
	var deleted string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if !present {
			w.WriteHeader(404)
			w.Write([]byte(`{"code":"NOT_FOUND"}`))
			return
		}
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(map[string]interface{}{"metadata": map[string]string{"namespace": "system", "name": "tf-acc-test-runone-swagger", "version": "v1"}, "string_value": content, "content_format": "json"})
		case "DELETE":
			deleted = r.URL.Path
			present = false
			w.Write([]byte(`{}`))
		default:
			t.Fatal("unexpected operation")
		}
	}))
	defer server.Close()
	c := client.NewClient(server.URL, "test")
	r := OwnedResource{Type: "xcsh_swagger_object", Name: "tf-acc-test-runone-swagger", Namespace: "system", Version: "v1", ContentSHA256: publicContentDigest(content)}
	j, err := CreateOwnershipJournal(filepath.Join(t.TempDir(), "journal.json"), "runone")
	if err != nil {
		t.Fatal(err)
	}
	if err = j.Intent(r); err != nil {
		t.Fatal(err)
	}
	if err = j.Confirm(r, "v1"); err != nil {
		t.Fatal(err)
	}
	receipt, err := j.Cleanup(context.Background(), apiOwnershipBackend{c: c}, 3, 0)
	if err != nil || !receipt.Complete || deleted != "/api/object_store/namespaces/system/stored_objects/swagger/tf-acc-test-runone-swagger/v1" {
		t.Fatalf("cleanup %#v %v %s", receipt, err, deleted)
	}
}
