// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Site registration client types for F5 XC
// Read-only access to the registrations of a site's CE nodes. A registration
// is named "r-<uuid>" — NOT after the site — so it can only be found by
// listing the registrations that belong to the site.

package client

import (
	"context"
	"fmt"
)

// RegistrationPassport carries the CE-supplied cluster identity. ClusterName is
// the F5 XC site name the node registered against, and is the field used to
// match a registration to a site.
type RegistrationPassport struct {
	ClusterName string `json:"cluster_name,omitempty"`
	ClusterSize int64  `json:"cluster_size,omitempty"`
}

// RegistrationInfra carries the node-level facts reported by the registering
// CE. Hostname distinguishes the nodes of a multi-node site; it is NOT unique
// tenant-wide (several sites report the same hostname), so it is only a
// within-site discriminator.
type RegistrationInfra struct {
	Provider   string `json:"provider,omitempty"`
	Hostname   string `json:"hostname,omitempty"`
	InstanceID string `json:"instance_id,omitempty"`
}

// RegistrationGetSpec is the registration's spec view.
type RegistrationGetSpec struct {
	Infra    RegistrationInfra    `json:"infra,omitempty"`
	Passport RegistrationPassport `json:"passport,omitempty"`
}

// RegistrationStatus holds the registration's lifecycle state. CurrentState is
// a registrationObjectState enum: NOTSET, NEW, APPROVED, ADMITTED, RETIRED,
// FAILED, DONE, PENDING, ONLINE, UPGRADING, MAINTENANCE, FAILED_INACTIVE.
type RegistrationStatus struct {
	CurrentState string `json:"current_state,omitempty"`
}

// RegistrationObject is the embedded full object; only its status is modelled.
type RegistrationObject struct {
	Status RegistrationStatus `json:"status,omitempty"`
}

// RegistrationSystemMetadata is the system-populated metadata; only the uid is
// modelled.
type RegistrationSystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

// RegistrationListItem is a single registration returned by
// ListRegistrationsBySite. Name is the approvable registration name
// ("r-" + uid).
type RegistrationListItem struct {
	Name           string                     `json:"name,omitempty"`
	UID            string                     `json:"uid,omitempty"`
	SystemMetadata RegistrationSystemMetadata `json:"system_metadata,omitempty"`
	Object         RegistrationObject         `json:"object,omitempty"`
	GetSpec        RegistrationGetSpec        `json:"get_spec,omitempty"`
}

// RegistrationListResponse is the registrationListResponse envelope. A site
// with no registered node yields HTTP 200 with an empty Items slice — not a
// 404 and not an error.
type RegistrationListResponse struct {
	Items  []RegistrationListItem `json:"items"`
	Errors []APIErrorDetail       `json:"errors,omitempty"`
}

// APIErrorDetail is a non-fatal error entry returned alongside items.
type APIErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ListRegistrationsBySite lists the CE registrations belonging to a site.
// Uses GetLenient because registration objects embed certificate and log
// strings that can contain raw control bytes, which the strict JSON decoder
// rejects. An unregistered or unknown site returns an empty Items slice with a
// nil error.
func (c *Client) ListRegistrationsBySite(ctx context.Context, namespace, siteName string) (*RegistrationListResponse, error) {
	var result RegistrationListResponse
	path := fmt.Sprintf("/api/register/namespaces/%s/registrations_by_site/%s", namespace, siteName)
	err := c.GetLenient(ctx, path, &result)
	return &result, err
}
