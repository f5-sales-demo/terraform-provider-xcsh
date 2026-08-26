// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import "testing"

func TestConcurrencyTokenPrivateStateRoundTrip(t *testing.T) {
	for _, token := range []string{"current-token", "advanced:v2/opaque==", "opaque-世界"} {
		t.Run(token, func(t *testing.T) {
			raw, err := encodeConcurrencyToken(token)
			if err != nil {
				t.Fatalf("encodeConcurrencyToken: %v", err)
			}
			if string(raw) == token {
				t.Fatal("private-state value must be valid JSON, not an unquoted raw token")
			}
			got, err := decodeConcurrencyToken(raw)
			if err != nil {
				t.Fatalf("decodeConcurrencyToken: %v", err)
			}
			if got != token {
				t.Fatalf("decoded token = %q, want %q", got, token)
			}
		})
	}
}

func TestConcurrencyTokenPrivateStateRejectsMissingOrInvalid(t *testing.T) {
	for name, raw := range map[string][]byte{
		"missing":      nil,
		"invalid JSON": []byte("not-json"),
		"empty token":  []byte(`""`),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := decodeConcurrencyToken(raw); err == nil {
				t.Fatal("expected invalid private token to fail")
			}
		})
	}
	if _, err := encodeConcurrencyToken(""); err == nil {
		t.Fatal("expected empty API token to fail encoding")
	}
}
