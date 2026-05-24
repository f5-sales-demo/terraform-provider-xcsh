// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package namespace provides consistent namespace classification for F5XC resources.
// Resources are classified via Profile structs set at code-generation time from enriched
// API specs. No hardcoded resource maps — all classification is data-driven.
package namespace

import (
	"sync"
)

// Type represents the namespace type enum.
type Type int

const (
	// SystemType is for infrastructure objects (sites, networks, fleet, cluster, cloud credentials).
	SystemType Type = iota
	// SharedType is for cross-app security policies (app_firewall, certificates, rate_limiters).
	SharedType
	// Application is for app-specific workloads (load balancers, origin pools, healthchecks).
	Application
)

// String returns the string representation of the namespace type.
func (t Type) String() string {
	switch t {
	case SystemType:
		return "system"
	case SharedType:
		return "shared"
	default:
		return "staging"
	}
}

// NamespaceType is a string-typed namespace classifier used in Profile definitions.
type NamespaceType string

const (
	// System namespace — infrastructure-level objects.
	System NamespaceType = "system"
	// Shared namespace — cross-app reusable policies.
	Shared NamespaceType = "shared"
	// Default namespace — the tenant's default namespace.
	Default NamespaceType = "default"
	// Custom namespace — user-created application namespaces.
	Custom NamespaceType = "custom"
)

// Profile describes the namespace behaviour for a single resource type.
type Profile struct {
	Allowed            []NamespaceType
	Enforced           bool
	Recommended        NamespaceType
	Category           string
	MultiTenantPattern string
}

// IsAllowed returns true if the given NamespaceType is in the Allowed list.
func (p Profile) IsAllowed(nsType NamespaceType) bool {
	for _, a := range p.Allowed {
		if a == nsType {
			return true
		}
	}
	return false
}

var (
	profiles   = make(map[string]Profile)
	profilesMu sync.RWMutex
)

// SetProfile registers a namespace profile for the named resource.
func SetProfile(resourceName string, profile Profile) {
	profilesMu.Lock()
	defer profilesMu.Unlock()
	profiles[resourceName] = profile
}

// GetProfile retrieves the profile for a resource. The bool is false when no
// profile has been registered.
func GetProfile(resourceName string) (Profile, bool) {
	profilesMu.RLock()
	defer profilesMu.RUnlock()
	p, ok := profiles[resourceName]
	return p, ok
}

// ClearProfiles removes all registered profiles (useful in tests).
func ClearProfiles() {
	profilesMu.Lock()
	defer profilesMu.Unlock()
	profiles = make(map[string]Profile)
}

// ForResource returns the Type and namespace string for a resource.
// It consults the profiles map first. If a profile exists with exactly one
// Allowed namespace of System or Shared, the corresponding type is returned.
// Everything else (multi-namespace, custom, default, or no profile) defaults
// to Application / "staging".
func ForResource(name string) (Type, string) {
	profilesMu.RLock()
	p, ok := profiles[name]
	profilesMu.RUnlock()

	if ok {
		if len(p.Allowed) == 1 {
			switch p.Allowed[0] {
			case System:
				return SystemType, "system"
			case Shared:
				return SharedType, "shared"
			}
		}
		return Application, "staging"
	}

	return Application, "staging"
}

// ForReference returns the namespace string for a referenced resource.
// Callers pass the resource type name (e.g. "app_firewall") and get back
// the namespace string ("system", "shared", or "staging").
func ForReference(referencedResourceType string) string {
	_, ns := ForResource(referencedResourceType)
	return ns
}

// IsSystem returns true if the resource belongs in the system namespace.
func IsSystem(name string) bool {
	t, _ := ForResource(name)
	return t == SystemType
}

// IsShared returns true if the resource belongs in the shared namespace.
func IsShared(name string) bool {
	t, _ := ForResource(name)
	return t == SharedType
}

// IsApplication returns true if the resource belongs in an application namespace.
func IsApplication(name string) bool {
	return !IsSystem(name) && !IsShared(name)
}

