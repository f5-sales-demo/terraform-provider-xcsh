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

func TestRegistrationInstanceIDDistinguishesRebuiltIdentity(t *testing.T) {
	first := item("r-first", "rebuilt-site", "master-0", "ONLINE")
	first.GetSpec.Infra.InstanceID = "instance-old"
	second := item("r-second", "rebuilt-site", "master-0", "ONLINE")
	second.GetSpec.Infra.InstanceID = "instance-new"

	var firstState, secondState SiteRegistrationDataSourceModel
	applyRegistrationMatch(&firstState, &first)
	applyRegistrationMatch(&secondState, &second)
	if firstState.InstanceID.ValueString() != "instance-old" || secondState.InstanceID.ValueString() != "instance-new" {
		t.Fatalf("instance identities were not preserved: first=%q second=%q", firstState.InstanceID.ValueString(), secondState.InstanceID.ValueString())
	}
	if first.GetSpec.Passport.ClusterName != second.GetSpec.Passport.ClusterName || first.GetSpec.Infra.Hostname != second.GetSpec.Infra.Hostname {
		t.Fatal("fixture must hold hostname and cluster identity constant")
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

// Rebuilding a CE under the same site name leaves the previous registration
// behind in F5 XC, and registrations_by_site keeps returning it: the tenant
// really does carry FAILED_INACTIVE registrations whose cluster_name AND
// hostname are identical to the live node's. Neither filter can separate them,
// so without dropping terminal states this resolved to an ambiguity error and
// failed the consumer's whole plan — strictly worse than found = false.
func TestSelectRegistration_StaleRegistrationIsSkipped(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-stale", "appstack-demo", "master-0", "FAILED_INACTIVE"),
		item("r-live", "appstack-demo", "master-0", "ONLINE"),
	}

	got, err := selectRegistration(items, "appstack-demo", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil: a stale registration must not make the lookup ambiguous", err)
	}
	if got == nil {
		t.Fatal("selectRegistration() = nil, want the live ONLINE registration")
	}
	if got.Name != "r-live" {
		t.Errorf("Name = %q, want %q", got.Name, "r-live")
	}
	if got.Object.Status.CurrentState != "ONLINE" {
		t.Errorf("CurrentState = %q, want %q", got.Object.Status.CurrentState, "ONLINE")
	}
}

// Order must not matter: the live registration is selected whether it precedes
// or follows the stale one in the API response.
func TestSelectRegistration_StaleRegistrationIsSkippedRegardlessOfOrder(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-live", "appstack-demo", "master-0", "PENDING"),
		item("r-stale", "appstack-demo", "master-0", "FAILED_INACTIVE"),
	}

	got, err := selectRegistration(items, "appstack-demo", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil", err)
	}
	if got == nil || got.Name != "r-live" {
		t.Fatalf("selectRegistration() = %+v, want the r-live registration", got)
	}
}

// Every terminal registrationObjectState is skipped in favour of a live one.
func TestSelectRegistration_EveryTerminalStateIsSkipped(t *testing.T) {
	for _, state := range []string{"RETIRED", "FAILED", "FAILED_INACTIVE", "DONE"} {
		t.Run(state, func(t *testing.T) {
			items := []client.RegistrationListItem{
				item("r-stale", "ha-site", "master-0", state),
				item("r-live", "ha-site", "master-0", "ONLINE"),
			}

			got, err := selectRegistration(items, "ha-site", "")
			if err != nil {
				t.Fatalf("selectRegistration() error = %v, want nil for a stale %s registration", err, state)
			}
			if got == nil || got.Name != "r-live" {
				t.Fatalf("selectRegistration() = %+v, want the r-live registration", got)
			}
		})
	}
}

// A non-terminal state is never skipped, so a site awaiting approval still
// resolves.
func TestSelectRegistration_NonTerminalStatesAreKept(t *testing.T) {
	for _, state := range []string{"NOTSET", "NEW", "APPROVED", "ADMITTED", "PENDING", "ONLINE", "UPGRADING", "MAINTENANCE"} {
		t.Run(state, func(t *testing.T) {
			items := []client.RegistrationListItem{
				item("r-aaa", "ha-site", "master-0", state),
			}

			got, err := selectRegistration(items, "ha-site", "")
			if err != nil {
				t.Fatalf("selectRegistration() error = %v, want nil", err)
			}
			if got == nil || got.Name != "r-aaa" {
				t.Fatalf("selectRegistration() = %+v, want the r-aaa registration for state %s", got, state)
			}
		})
	}
}

// When every candidate is terminal there is nothing live to approve, which is
// the benign found = false outcome — NOT an error. A site whose CE was
// destroyed must not fail its consumer's plan.
func TestSelectRegistration_AllTerminalIsNotAnError(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-aaa", "ha-site", "master-0", "FAILED_INACTIVE"),
		item("r-bbb", "ha-site", "master-1", "RETIRED"),
	}

	got, err := selectRegistration(items, "ha-site", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil when every registration is terminal", err)
	}
	if got != nil {
		t.Fatalf("selectRegistration() = %+v, want nil", got)
	}
}

// Dropping terminal candidates must not weaken the ambiguity guard: a genuine
// multi-node site (cluster_size 3) with no hostname set is still an error, and
// the message names only the live candidates.
func TestSelectRegistration_AmbiguityAmongLiveCandidatesSurvivesTerminalFiltering(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-stale", "ha-site", "master-9", "RETIRED"),
		item("r-aaa", "ha-site", "master-0", "ONLINE"),
		item("r-bbb", "ha-site", "master-1", "PENDING"),
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
			t.Errorf("error %q does not mention live candidate hostname %q", err.Error(), want)
		}
	}
	if strings.Contains(err.Error(), "master-9") {
		t.Errorf("error %q names the retired candidate hostname master-9", err.Error())
	}
	if !strings.Contains(err.Error(), "2 matching registrations") {
		t.Errorf("error %q does not count exactly the 2 live candidates", err.Error())
	}
}

// A registration reporting no state at all is kept, for the same defensive
// reason a registration reporting no cluster_name is: a missing field must
// never turn a real match into a miss.
func TestSelectRegistration_MissingStateIsKept(t *testing.T) {
	items := []client.RegistrationListItem{
		item("r-aaa", "ha-site", "master-0", ""),
	}

	got, err := selectRegistration(items, "ha-site", "")
	if err != nil {
		t.Fatalf("selectRegistration() error = %v, want nil", err)
	}
	if got == nil || got.Name != "r-aaa" {
		t.Fatalf("selectRegistration() = %+v, want the r-aaa registration", got)
	}
}
