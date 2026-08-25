// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import "testing"

func TestConcurrencyTokenPrivateStateRoundTrip(t *testing.T) {
	raw, err := encodeConcurrencyToken("opaque-token")
	if err != nil {
		t.Fatalf("encodeConcurrencyToken: %v", err)
	}
	if string(raw) == "opaque-token" {
		t.Fatal("private-state value must be valid JSON, not an unquoted raw token")
	}
	got, err := decodeConcurrencyToken(raw)
	if err != nil {
		t.Fatalf("decodeConcurrencyToken: %v", err)
	}
	if got != "opaque-token" {
		t.Fatalf("decoded token = %q, want opaque-token", got)
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
