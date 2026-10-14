// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestAPIDiagnosticDoesNotExposeResponse(t *testing.T) {
	old := log.Writer()
	defer log.SetOutput(old)
	var captured bytes.Buffer
	log.SetOutput(&captured)
	body := `{"message":"Quota not configured for kind SECRET_SENTINEL; namespace PRIVATE_SENTINEL"}`
	response := &http.Response{StatusCode: 429, Body: io.NopCloser(strings.NewReader(body))}
	logAcceptanceAPIDiagnostic(response)
	if strings.Contains(captured.String(), "SECRET_SENTINEL") || strings.Contains(captured.String(), "PRIVATE_SENTINEL") || !strings.Contains(captured.String(), "[API_DIAGNOSTIC:quota_unconfigured]") {
		t.Fatal("unsafe or missing diagnostic")
	}
	data, err := io.ReadAll(response.Body)
	if err != nil || !strings.Contains(string(data), "SECRET_SENTINEL") || !strings.Contains(string(data), "[API_DIAGNOSTIC:quota_unconfigured]") {
		t.Fatal("response mutated")
	}
}
