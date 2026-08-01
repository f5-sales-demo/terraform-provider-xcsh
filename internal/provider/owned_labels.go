// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"encoding/json"
	"sort"
)

// ownedLabelKeysPrivateKey is the private-state key holding the label keys the
// configuration declared at the last Create or Update.
//
// Private state is the only place this can live. Read receives prior state and the API
// response and nothing else, so it cannot tell a label the configuration asked for from
// one the platform added — and prior state is not a substitute: every provider release up
// to 3.81.1 wrote the platform's own labels into state, which is exactly why the #1391
// attempt at inferring ownership from prior state had to be reverted.
const ownedLabelKeysPrivateKey = "ownedLabelKeys"

// ownedLabelKeysPayload is the private-state document. A named field rather than a bare
// array so a future addition does not have to change the encoding.
type ownedLabelKeysPayload struct {
	Keys []string `json:"ownedLabelKeys"`
}

// encodeOwnedLabelKeys serialises the configuration's label keys for private state.
//
// Keys are sorted so the encoding depends only on the set, not on map iteration order.
// Without that, every apply would rewrite private state and Terraform would report a
// change on a configuration nobody touched.
//
// nolint:unused // Used by generated resource Create and Update methods
func encodeOwnedLabelKeys(keys []string) ([]byte, error) {
	sorted := make([]string, len(keys))
	copy(sorted, keys)
	sort.Strings(sorted)
	return json.Marshal(ownedLabelKeysPayload{Keys: sorted})
}

// decodeOwnedLabelKeys reads the private-state entry written by encodeOwnedLabelKeys.
//
// Anything unreadable — absent, empty, not JSON, or the wrong shape — means "owns
// nothing". That is deliberate and it is the upgrade path: every resource created before
// this change has no entry, and treating it as owning nothing reproduces the #1391
// behaviour, where the platform's labels are filtered and the plan is clean. Failing
// loudly here would break every existing resource on the first refresh after upgrade.
//
// nolint:unused // Used by generated resource Read methods
func decodeOwnedLabelKeys(raw []byte) map[string]struct{} {
	owned := make(map[string]struct{})
	if len(raw) == 0 {
		return owned
	}
	var payload ownedLabelKeysPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		return owned
	}
	for _, k := range payload.Keys {
		owned[k] = struct{}{}
	}
	return owned
}

// filterSystemLabelsOwning is filterSystemLabels with an exemption list: a platform-
// namespaced or discovery key the configuration declared is kept, so Terraform can see
// what it wrote and the plan converges.
//
// Passing a nil or empty owned set is identical to filterSystemLabels, which is what
// makes this safe for resources whose state predates the private-state entry.
//
// nolint:unused // Used by generated resource Read methods
func filterSystemLabelsOwning(labels map[string]string, siteDiscovery bool, owned map[string]struct{}) map[string]string {
	filtered := make(map[string]string, len(labels))
	for k, v := range labels {
		if _, isOwned := owned[k]; isOwned {
			// The configuration declared this key, so it is the configuration's to
			// manage even though it sits in a namespace the platform also uses.
			filtered[k] = v
			continue
		}
		if !isPlatformLabel(k, siteDiscovery) {
			filtered[k] = v
		}
	}
	return filtered
}

// configLabelKeys returns the keys of the labels a configuration declares, for recording
// as the owned set. A nil map yields an empty slice rather than nil so the encoded
// payload is the same shape whether or not any labels were declared.
//
// nolint:unused // Used by generated resource Create and Update methods
func configLabelKeys(labels map[string]string) []string {
	keys := make([]string, 0, len(labels))
	for k := range labels {
		keys = append(keys, k)
	}
	return keys
}
