// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

// item builds a RegistrationListItem carrying only the fields selectRegistration
// looks at.
func item(name, clusterName, hostname, state string) client.RegistrationListItem {
	return client.RegistrationListItem{
		Name: name,
		Object: client.RegistrationObject{
			Status: client.RegistrationStatus{CurrentState: state},
		},
		GetSpec: client.RegistrationGetSpec{
			Infra:    client.RegistrationInfra{Hostname: hostname},
			Passport: client.RegistrationPassport{ClusterName: clusterName},
		},
	}
}

// A registered single-node site resolves to its one registration.
func TestSelectRegistration_SingleMatch(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-dcec2400-52d5-4154-9fd0-4b042d3fe18d", "ar-bgp-eastus01", "f5-xc-ce-vm-01", "ONLINE"),
	}

	got, err := selectRegistration(items, "ar-bgp-eastus01", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("selectRegistration() = nil, want the single registration")
	}
	if got.Name != "r-dcec2400-52d5-4154-9fd0-4b042d3fe18d" {
		t.Errorf("Name = %q, want %q", got.Name, "r-dcec2400-52d5-4154-9fd0-4b042d3fe18d")
	}
	if got.Object.Status.CurrentState != "ONLINE" {
		t.Errorf("CurrentState = %q, want %q", got.Object.Status.CurrentState, "ONLINE")
	}
}

// The normal not-yet-registered case: no items, no match, and crucially NO
// error — a consumer gates a count on this and an error would break early
// plans.
func TestSelectRegistration_EmptyItemsIsNotAnError(t *testing.T) {
	got, err := selectRegistration(nil, "ar-bgp-eastus01", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil for an unregistered site", err)
	}
	if got != nil {
		t.Fatalf("selectRegistration() = %+v, want nil", got)
	}

	got, err = selectRegistration([]client.RegistrationListItem{}, "ar-bgp-eastus01", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil for an empty items slice", err)
	}
	if got != nil {
		t.Fatalf("selectRegistration() = %+v, want nil", got)
	}
}

// A multi-node site is disambiguated by hostname.
func TestSelectRegistration_HostnameDiscriminator(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-aaa", "ha-site", "master-0", "ONLINE"),
		item("r-bbb", "ha-site", "master-1", "PENDING"),
	}

	got, err := selectRegistration(items, "ha-site", "master-1")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil", err)
	}
	if got == nil {
		t.Fatal("selectRegistration() = nil, want the master-1 registration")
	}
	if got.Name != "r-bbb" {
		t.Errorf("Name = %q, want %q", got.Name, "r-bbb")
	}
}

// Several candidates and no way to choose is a configuration error, and the
// message must name the candidate hostnames so the operator can set one.
func TestSelectRegistration_AmbiguousIsAnError(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-aaa", "ha-site", "master-0", "ONLINE"),
		item("r-bbb", "ha-site", "master-1", "ONLINE"),
	}

	got, err := selectRegistration(items, "ha-site", "")
	if err == nil {
		t.Fatalf("selectRegistration() = %+v, want an ambiguity error", got)
	}
	if got != nil {
		t.Errorf("selectRegistration() = %+v, want nil alongside the error", got)
	}
	for _, want := range []string{"master-0", "master-1"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention candidate hostname %q", err.Error(), want)
		}
	}
}

// A hostname that matches nothing is the same benign "not registered yet"
// outcome as an empty list, not an error.
func TestSelectRegistration_UnmatchedHostnameIsNotAnError(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-aaa", "ha-site", "master-0", "ONLINE"),
	}

	got, err := selectRegistration(items, "ha-site", "master-9")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil", err)
	}
	if got != nil {
		t.Fatalf("selectRegistration() = %+v, want nil", got)
	}
}

// The endpoint is already site-scoped, but a registration reporting a
// different cluster_name belongs to another site and must not be selected.
func TestSelectRegistration_ForeignClusterNameIsFilteredOut(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-aaa", "some-other-site", "master-0", "ONLINE"),
	}

	got, err := selectRegistration(items, "ar-bgp-eastus01", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil", err)
	}
	if got != nil {
		t.Fatalf("selectRegistration() = %+v, want nil", got)
	}
}

// A registration that reports no cluster_name at all is kept: the endpoint
// already scoped the list to the site, so dropping it would be a false
// negative.
func TestSelectRegistration_MissingClusterNameIsKept(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-aaa", "", "master-0", "NEW"),
	}

	got, err := selectRegistration(items, "ar-bgp-eastus01", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil", err)
	}
	if got == nil || got.Name != "r-aaa" {
		t.Fatalf("selectRegistration() = %+v, want the r-aaa registration", got)
	}
}
