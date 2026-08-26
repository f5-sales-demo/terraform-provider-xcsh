// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package mocks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestConfigObjectConcurrencyLifecycle(t *testing.T) {
	server := NewServer()
	defer server.Close()

	collection := server.URL() + "/api/config/namespaces/system/widgets"
	create := []byte(`{"metadata":{"name":"example"},"spec":{"value":"original"}}`)
	response, err := http.Post(collection, "application/json", bytes.NewReader(create)) //nolint:noctx // local test server
	if err != nil {
		t.Fatal(err)
	}
	created := decodeMockResponse(t, response)
	firstToken, ok := created["resource_version"].(string)
	if !ok || firstToken == "" {
		t.Fatal("create response has no opaque resource_version")
	}

	objectURL := collection + "/example"
	current := []byte(`{"metadata":{"name":"example"},"spec":{"value":"current"},"resource_version":` + mustJSON(t, firstToken) + `}`)
	request, err := http.NewRequest(http.MethodPut, objectURL, bytes.NewReader(current))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	updated := decodeMockResponse(t, response)
	advancedToken, ok := updated["resource_version"].(string)
	if !ok || advancedToken == "" || advancedToken == firstToken {
		t.Fatal("successful replace did not advance resource_version")
	}

	stale := []byte(`{"metadata":{"name":"example"},"spec":{"value":"stale-overwrite"},"resource_version":` + mustJSON(t, firstToken) + `}`)
	request, err = http.NewRequest(http.MethodPut, objectURL, bytes.NewReader(stale))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusConflict {
		body, _ := io.ReadAll(response.Body)
		_ = response.Body.Close()
		t.Fatalf("stale replace status = %d, want 409: %s", response.StatusCode, body)
	}
	_ = response.Body.Close()

	stored, found := server.GetResource("/api/config/namespaces/system/widgets/example")
	if !found {
		t.Fatal("updated resource disappeared")
	}
	storedObject := stored.(map[string]interface{})
	if storedObject["resource_version"] != advancedToken || storedObject["spec"].(map[string]interface{})["value"] != "current" {
		t.Fatalf("stale replace mutated the stored object: %#v", storedObject)
	}
}

func TestConfigObjectPathFamiliesAreVersioned(t *testing.T) {
	t.Parallel()

	for _, path := range []string{
		"/api/config/namespaces/system/widgets",
		"/api/web/namespaces",
		"/api/register/namespaces/system/tokens",
	} {
		if !isConfigObjectPath(path) {
			t.Errorf("%s was not classified as a versioned API object path", path)
		}
	}
	if isConfigObjectPath("/healthz") {
		t.Fatal("non-API health endpoint was classified as a config object")
	}
}

func decodeMockResponse(t *testing.T, response *http.Response) map[string]interface{} {
	t.Helper()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		t.Fatalf("response status = %d: %s", response.StatusCode, body)
	}
	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	return result
}

func mustJSON(t *testing.T, value interface{}) string {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return string(encoded)
}
