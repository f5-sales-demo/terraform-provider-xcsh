// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package mocks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestReplacePreservesOriginalCreationIdentity(t *testing.T) {
	for _, family := range []string{"config", "web", "register"} {
		t.Run(family, func(t *testing.T) {
			s := NewServer()
			defer s.Close()
			endpoint := s.URL() + "/api/" + family + "/namespaces/system/securemesh_site_v2s"
			request := map[string]interface{}{
				"metadata": map[string]interface{}{"name": "identity-fixture"},
				"spec":     map[string]interface{}{"description": "original"},
			}
			send := func(method, url string, body interface{}) map[string]interface{} {
				t.Helper()
				data, err := json.Marshal(body)
				if err != nil {
					t.Fatal(err)
				}
				req, err := http.NewRequest(method, url, bytes.NewReader(data))
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Set("Content-Type", "application/json")
				resp, err := s.Client().Do(req)
				if err != nil {
					t.Fatal(err)
				}
				defer resp.Body.Close()
				var result map[string]interface{}
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatal(err)
				}
				if resp.StatusCode != http.StatusOK {
					t.Fatalf("%s: status %d", method, resp.StatusCode)
				}
				return result
			}
			original := send(http.MethodPost, endpoint, request)
			originalMeta := original["metadata"].(map[string]interface{})
			originalSystem := original["system_metadata"].(map[string]interface{})
			// Ensure timestamp preservation is tested without timing-dependent sleeps.
			originalSystem["creation_timestamp"] = "2020-01-01T00:00:00Z"
			s.SetResource("/api/"+family+"/namespaces/system/securemesh_site_v2s/identity-fixture", original)
			request["resource_version"] = original["resource_version"]
			request["metadata"].(map[string]interface{})["uid"] = "forged-uid"
			request["system_metadata"] = map[string]interface{}{"uid": "forged-uid", "creation_timestamp": "forged"}
			request["spec"] = map[string]interface{}{"description": "updated"}
			updated := send(http.MethodPut, endpoint+"/identity-fixture", request)
			if updated["metadata"].(map[string]interface{})["uid"] != originalMeta["uid"] {
				t.Error("metadata UID changed on replace")
			}
			for _, key := range []string{"uid", "creation_timestamp", "creator_class", "creator_id", "tenant"} {
				if updated["system_metadata"].(map[string]interface{})[key] != originalSystem[key] {
					t.Errorf("server-owned %s changed on replace", key)
				}
			}
			if updated["resource_version"] == original["resource_version"] {
				t.Error("replace did not advance resource_version")
			}
			if updated["spec"].(map[string]interface{})["description"] != "updated" {
				t.Error("replace lost changed spec")
			}
			read := send(http.MethodGet, endpoint+"/identity-fixture", nil)
			if read["system_metadata"].(map[string]interface{})["uid"] != originalSystem["uid"] {
				t.Error("stored UID changed")
			}
			send(http.MethodDelete, endpoint+"/identity-fixture", nil)
			delete(request, "resource_version")
			delete(request, "system_metadata")
			delete(request["metadata"].(map[string]interface{}), "uid")
			recreated := send(http.MethodPost, endpoint, request)
			if recreated["system_metadata"].(map[string]interface{})["uid"] == originalSystem["uid"] {
				t.Error("recreation reused original identity")
			}
		})
	}
}
