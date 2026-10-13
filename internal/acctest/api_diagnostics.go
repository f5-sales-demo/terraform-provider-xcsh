// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
)

func logAcceptanceAPIDiagnostic(response *http.Response) {
	if response == nil || response.StatusCode < 400 || response.StatusCode >= 500 {
		return
	}
	data, err := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		return
	}
	// Preserve the original response for the API client; only fixed facts are logged.
	response.Body = io.NopCloser(bytes.NewReader(data))
	text := strings.ToLower(string(data))
	signatures := map[string][]string{
		"quota_unconfigured":     {"quota not configured"},
		"quota_exceeded":         {"quota exceeded", "quota limit", "maximum number"},
		"swagger_reference":      {"swagger", "openapi"},
		"namespace_constraint":   {"namespace"},
		"missing_required_input": {"required", "must specify", "cannot be empty"},
		"schema_validation":      {"validation", "invalid request", "bad request"},
		"authentication_denied":  {"unauthorized", "authentication", "invalid token"},
	}
	for name, patterns := range signatures {
		for _, pattern := range patterns {
			if strings.Contains(text, pattern) {
				log.Printf("[API_DIAGNOSTIC:%s]", name)
				break
			}
		}
	}
}
