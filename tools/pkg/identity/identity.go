// Package identity defines the fleet's canonical synthetic organization values.
package identity

import "strings"

const (
	Namespace    = "demo-app"
	Organization = "example-corp"
	NumericID    = "123456789012"
)

var canonicalValues = map[string]bool{
	NumericID:    true,
	"default":    true,
	"demo":       true,
	Namespace:    true,
	Organization: true,
	"library":    true,
	"production": true,
	"shared":     true,
	"staging":    true,
	"system":     true,
}

// Canonical returns the canonical synthetic value for a recognized identity
// field. Already-safe values and unrecognized fields are preserved.
func Canonical(field, value string) string {
	field = strings.ToLower(strings.TrimSpace(field))
	value = strings.TrimSpace(value)
	if canonicalValues[strings.ToLower(value)] {
		return value
	}

	switch field {
	case "namespace":
		return Namespace
	case "tenant", "tenant_name", "customer", "customer_name", "account_name",
		"subscription_name", "project_name":
		return Organization
	case "tenant_id", "customer_id", "account", "account_id", "subscription",
		"subscription_id", "project", "project_id":
		return NumericID
	default:
		return value
	}
}
