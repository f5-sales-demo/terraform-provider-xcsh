// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
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
	facts := []string{}
	for name, patterns := range signatures {
		for _, pattern := range patterns {
			if strings.Contains(text, pattern) {
				facts = append(facts, name)
				log.Printf("[API_DIAGNOSTIC:%s]", name)
				logDiagnosticFact(name)
				break
			}
		}
	}
	sort.Strings(facts)
	var body map[string]interface{}
	if json.Unmarshal(data, &body) == nil {
		message, _ := body["message"].(string)
		for _, fact := range facts {
			message += " [API_DIAGNOSTIC:" + fact + "]"
		}
		body["message"] = message
		if serialized, err := json.Marshal(body); err == nil {
			response.Body = io.NopCloser(bytes.NewReader(serialized))
			response.ContentLength = int64(len(serialized))
		}
	}
}
