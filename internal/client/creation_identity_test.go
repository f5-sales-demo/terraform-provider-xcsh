package client

import (
	"encoding/json"
	"testing"
)

func TestCreationIdentityResponse(t *testing.T) {
	for _, response := range []string{`{"system_metadata":{"uid":"original-uid"},"spec":{"content":"different-sensitive-bootstrap"}}`, `{"spec":{"content":"different-sensitive-bootstrap"}}`} {
		var token Token
		if err := json.Unmarshal([]byte(response), &token); err != nil {
			t.Fatal(err)
		}
		if token.SystemMetadata != nil && token.ImmutableUID() != "original-uid" {
			t.Fatal("bootstrap substituted for original UID")
		}
		if token.SystemMetadata == nil && token.ImmutableUID() != "" {
			t.Fatal("missing UID acquired inferred ownership")
		}
	}
}
