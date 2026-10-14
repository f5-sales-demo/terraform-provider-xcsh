// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func TestOwnershipCreateReleasesResponseBeforeReadback(t *testing.T) {
	var mu sync.Mutex
	var stored []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPost:
			var payload map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				t.Error(err)
				return
			}
			metadata := payload["metadata"].(map[string]interface{})
			metadata["uid"] = "uid-one"
			data, err := json.Marshal(payload)
			if err != nil {
				t.Error(err)
				return
			}
			mu.Lock()
			stored = data
			mu.Unlock()
			_, _ = w.Write(data)
		case http.MethodGet:
			mu.Lock()
			data := append([]byte(nil), stored...)
			mu.Unlock()
			if data == nil {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"code":5}`))
				return
			}
			_, _ = w.Write(data)
		default:
			t.Error("unexpected mutation")
		}
	}))
	defer server.Close()
	j, err := CreateOwnershipJournal(t.TempDir()+"/journal.json", "runone")
	if err != nil {
		t.Fatal(err)
	}
	base := &http.Transport{MaxConnsPerHost: 1}
	defer base.CloseIdleConnections()
	c := client.NewClient(server.URL, "synthetic", client.WithTimeout(500*time.Millisecond))
	c.HTTPClient.Transport = &ownershipTransport{base: base, c: c, j: j}
	payload := map[string]interface{}{"metadata": map[string]interface{}{"name": "tf-acc-test-runone-pool", "namespace": "system"}, "spec": map[string]interface{}{"description": "body-preserved"}}
	var result map[string]interface{}
	if err = c.Post(context.Background(), "/api/config/namespaces/system/origin_pools", payload, &result); err != nil {
		t.Fatal("ownership readback blocked by unconsumed create response")
	}
	if len(j.Resources) != 1 || j.Resources[0].Status != "created" {
		t.Fatal("ownership confirmation missing")
	}
	if result["spec"].(map[string]interface{})["description"] != "body-preserved" {
		t.Fatal("create response changed while releasing connection")
	}
}
