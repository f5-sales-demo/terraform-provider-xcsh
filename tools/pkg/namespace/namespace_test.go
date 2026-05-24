// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package namespace

import (
	"testing"
)

func TestProfileDefaults(t *testing.T) {
	ClearProfiles()
	typ, ns := ForResource("unknown_resource")
	if typ != Application {
		t.Errorf("expected Application, got %v", typ)
	}
	if ns != "staging" {
		t.Errorf("expected staging, got %s", ns)
	}
}

func TestSetProfileSystem(t *testing.T) {
	ClearProfiles()
	SetProfile("aws_vpc_site", Profile{
		Allowed:     []NamespaceType{System},
		Enforced:    true,
		Recommended: System,
	})
	typ, ns := ForResource("aws_vpc_site")
	if typ != SystemType {
		t.Errorf("expected SystemType, got %v", typ)
	}
	if ns != "system" {
		t.Errorf("expected system, got %s", ns)
	}
}

func TestSetProfileShared(t *testing.T) {
	ClearProfiles()
	SetProfile("namespace_role_binding", Profile{
		Allowed:     []NamespaceType{Shared},
		Enforced:    true,
		Recommended: Shared,
	})
	typ, ns := ForResource("namespace_role_binding")
	if typ != SharedType {
		t.Errorf("expected SharedType, got %v", typ)
	}
	if ns != "shared" {
		t.Errorf("expected shared, got %s", ns)
	}
}

func TestSetProfileTenant(t *testing.T) {
	ClearProfiles()
	SetProfile("http_loadbalancer", Profile{
		Allowed:     []NamespaceType{Custom, Default, Shared},
		Enforced:    true,
		Recommended: Custom,
	})
	typ, ns := ForResource("http_loadbalancer")
	if typ != Application {
		t.Errorf("expected Application, got %v", typ)
	}
	if ns != "staging" {
		t.Errorf("expected staging, got %s", ns)
	}
}

func TestGetProfile(t *testing.T) {
	ClearProfiles()
	p := Profile{
		Allowed:            []NamespaceType{Shared, Custom},
		Enforced:           true,
		Recommended:        Shared,
		Category:           "security",
		MultiTenantPattern: "shared-ref",
	}
	SetProfile("app_firewall", p)
	got, ok := GetProfile("app_firewall")
	if !ok {
		t.Fatal("expected profile to exist")
	}
	if got.Recommended != Shared {
		t.Errorf("expected Shared, got %v", got.Recommended)
	}
	if got.Category != "security" {
		t.Errorf("expected security, got %s", got.Category)
	}
}

func TestIsAllowed(t *testing.T) {
	ClearProfiles()
	SetProfile("aws_vpc_site", Profile{
		Allowed: []NamespaceType{System},
	})
	p, _ := GetProfile("aws_vpc_site")
	if !p.IsAllowed(System) {
		t.Error("expected system to be allowed")
	}
	if p.IsAllowed(Custom) {
		t.Error("expected custom to not be allowed")
	}
}

func TestClearProfiles(t *testing.T) {
	SetProfile("test_resource", Profile{Allowed: []NamespaceType{System}})
	ClearProfiles()
	_, ok := GetProfile("test_resource")
	if ok {
		t.Error("expected profile to be cleared")
	}
}

func TestTypeString(t *testing.T) {
	tests := []struct {
		t    Type
		want string
	}{
		{SystemType, "system"},
		{SharedType, "shared"},
		{Application, "staging"},
	}
	for _, tt := range tests {
		if got := tt.t.String(); got != tt.want {
			t.Errorf("Type(%d).String() = %s, want %s", tt.t, got, tt.want)
		}
	}
}
