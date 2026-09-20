// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// responseOperationPrerequisite mirrors the immutable prerequisite contract
// attached to a generated response operation. It deliberately has no CRUD
// fields: a runtime lookup failure does not establish object ownership.
type responseOperationPrerequisite struct {
	ID              string
	Resource        string
	Exactly         int
	Enforcement     string
	Availability    string
	Reason          string
	SourceKind      string
	SourceOperation string
	SourceImmutable bool
	LookupScope     string
	LookupCount     string
	SourceCommit    string
	SpecSHA256      string
	ReceiptPath     string
	ReceiptSHA256   string
}

const mauriceConfigCardinalityError = "number of maurice_config object is not one"

// responseOperationPrerequisiteDiagnostic converts only a verified server
// rejection into a non-remediating diagnostic without echoing the raw response
// or inferring tenant ownership, actual count, or a corrective operation.
func responseOperationPrerequisiteDiagnostic(err error, prerequisites []responseOperationPrerequisite) (string, string, bool) {
	if err == nil {
		return "", "", false
	}
	for _, prerequisite := range prerequisites {
		if prerequisite.ID != "maurice_config_cardinality_exactly_one" ||
			prerequisite.Resource != "maurice_config" ||
			prerequisite.Exactly != 1 ||
			prerequisite.Enforcement != "server" ||
			prerequisite.Availability != "unresolved_server_lookup" ||
			prerequisite.LookupScope != "unknown" || prerequisite.LookupCount != "unknown" ||
			!validResponseOperationDigest(prerequisite.SourceCommit, 40) ||
			!validResponseOperationDigest(prerequisite.SpecSHA256, 64) ||
			!validResponseOperationDigest(prerequisite.ReceiptSHA256, 64) ||
			!strings.HasPrefix(prerequisite.ReceiptPath, "config/evidence/") ||
			!strings.HasSuffix(prerequisite.ReceiptPath, ".json") ||
			strings.Contains(prerequisite.ReceiptPath, "..") ||
			prerequisite.SourceKind != "runtime_api_error" ||
			prerequisite.SourceOperation != "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl" ||
			!prerequisite.SourceImmutable ||
			!strings.Contains(strings.ToLower(err.Error()), mauriceConfigCardinalityError) {
			continue
		}
		return "Image lookup unresolved", fmt.Sprintf(
			"The image API reported that its %s lookup did not resolve exactly one object. Lookup scope, actual count, and corrective operation are unknown. This does not establish tenant ownership. Consult the source contract and its evidence receipt %q (SHA-256 %s).",
			prerequisite.Resource, prerequisite.ReceiptPath, prerequisite.ReceiptSHA256), true
	}
	return "", "", false
}

func validResponseOperationDigest(value string, length int) bool {
	if len(value) != length || strings.ToLower(value) != value {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}
