// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"encoding/json"
	"fmt"
)

// concurrencyTokenPrivateKey stores the last server-returned optimistic-
// concurrency token. The value stays outside Terraform configuration and schema state.
const concurrencyTokenPrivateKey = "f5xcConcurrencyToken"

// encodeConcurrencyToken produces valid JSON because terraform-plugin-framework
// requires private-state values to be JSON and UTF-8 safe.
//
// nolint:unused // Used conditionally by generated resources.
func encodeConcurrencyToken(token string) ([]byte, error) {
	if token == "" {
		return nil, fmt.Errorf("server returned an empty concurrency token")
	}
	return json.Marshal(token)
}

// decodeConcurrencyToken fails closed. An older provider state has no token and
// must be refreshed before an update; fetching a newer token during the write would
// defeat optimistic concurrency by silently adopting changes the plan never saw.
//
// nolint:unused // Used conditionally by generated resources.
func decodeConcurrencyToken(raw []byte) (string, error) {
	if len(raw) == 0 {
		return "", fmt.Errorf("concurrency token is missing from private state")
	}
	var token string
	if err := json.Unmarshal(raw, &token); err != nil {
		return "", fmt.Errorf("decoding concurrency token from private state: %w", err)
	}
	if token == "" {
		return "", fmt.Errorf("concurrency token in private state is empty")
	}
	return token, nil
}
