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

func TestAPIDiagnosticsClassifyStorageAndParsingWithoutPayload(t *testing.T) {
	old := log.Writer()
	defer log.SetOutput(old)
	var output bytes.Buffer
	log.SetOutput(&output)
	response := &http.Response{StatusCode: 400, Body: io.NopCloser(strings.NewReader(`{"message":"Unable to download file SECRET_SENTINEL: failed to parse specification"}`))}
	logAcceptanceAPIDiagnostic(response)
	for _, fact := range []string{"document_fetch_failed", "document_parse_failed"} {
		if !strings.Contains(output.String(), "[API_DIAGNOSTIC:"+fact+"]") {
			t.Errorf("missing %s", fact)
		}
	}
	if strings.Contains(output.String(), "SECRET_SENTINEL") {
		t.Fatal("payload leaked")
	}
}
