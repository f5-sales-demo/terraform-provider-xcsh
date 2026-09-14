// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"fmt"
	"strings"
)

// responseOperationPrerequisite mirrors the immutable prerequisite contract
// attached to a generated response operation. It deliberately has no CRUD
// fields: the referenced tenant object remains externally managed.
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
}

const mauriceConfigCardinalityError = "number of maurice_config object is not one"

// responseOperationPrerequisiteDiagnostic converts only a verified server
// rejection into an actionable, non-remediating Terraform diagnostic. Unknown
// errors retain their original client diagnostic so unrelated failures are not
// misclassified as tenant prerequisites.
func responseOperationPrerequisiteDiagnostic(err error, prerequisites []responseOperationPrerequisite) (string, string, bool) {
	if err == nil {
		return "", "", false
	}
	for _, prerequisite := range prerequisites {
		if prerequisite.ID != "maurice_config_cardinality_exactly_one" ||
			prerequisite.Resource != "maurice_config" ||
			prerequisite.Exactly != 1 ||
			prerequisite.Enforcement != "server" ||
			prerequisite.Availability != "external_tenant_prerequisite" ||
			prerequisite.SourceKind != "runtime_api_error" ||
			prerequisite.SourceOperation != "ves.io.schema.registration.CustomAPI.GetImageDownloadUrl" ||
			!prerequisite.SourceImmutable ||
			!strings.Contains(strings.ToLower(err.Error()), mauriceConfigCardinalityError) {
			continue
		}
		return "External tenant prerequisite unavailable", fmt.Sprintf(
			"%s Contract prerequisite %q requires exactly one %s object. The provider cannot create or repair this tenant-owned object. Original server response: %s",
			prerequisite.Reason, prerequisite.ID, prerequisite.Resource, err), true
	}
	return "", "", false
}
