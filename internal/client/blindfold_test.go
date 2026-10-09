package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNativeBlindfoldUncertainWrite(t *testing.T) {
	desired := map[string]any{"metadata": map[string]any{"name": "synthetic", "annotations": map[string]any{"f5-sales-demo.com/blindfold": "synthetic"}}, "spec": map[string]any{"certificate_url": "string:///public", "private_key": map[string]any{"blindfold_secret_info": map[string]any{"location": "string:///encrypted"}}}}
	writes := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			writes++
			w.WriteHeader(503)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"replace_form": desired, "resource_version": "version"})
	}))
	defer server.Close()
	c := NewClient(server.URL, "synthetic")
	var result map[string]any
	if err := c.Post(context.Background(), "/api/config/namespaces/synthetic/certificates", desired, &result); err != nil {
		t.Fatal(err)
	}
	if writes != 1 {
		t.Fatal("mutation retried")
	}
	if result["resource_version"] != "version" {
		t.Fatal("token lost")
	}
}
