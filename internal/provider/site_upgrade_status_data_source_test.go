// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

const (
	baselineSoftware = "crt-20251002-0027"
	targetSoftware   = "crt-20260201-0179"
	baselineOS       = "9.2026.10"
	targetOS         = "9.2026.17"
)

func upgradeSiteObservation(software, operatingSystem, result string) client.SMSv2Observation {
	return client.SMSv2Observation{
		"spec": map[string]interface{}{"site_state": "ONLINE"},
		"status": []interface{}{
			map[string]interface{}{
				"operating_system_status": map[string]interface{}{
					"available_version": targetOS,
					"deployment_state": map[string]interface{}{
						"version": operatingSystem, "phase": "COMPLETED", "result": result,
					},
				},
				"volterra_software_status": map[string]interface{}{
					"last_installed_version": software, "available_version": targetSoftware,
					"deployment_state": map[string]interface{}{"phase": "COMPLETED", "result": result},
				},
			},
		},
	}
}

func passingPrecheck() client.SMSv2Observation {
	return client.SMSv2Observation{"checklist": []interface{}{
		map[string]interface{}{"item": "software compatibility", "status": "CHECKLIST_PASSED"},
		map[string]interface{}{"item": "capacity", "status": "CHECKLIST_WARNING"},
	}}
}

func TestExtractSiteUpgradeStatusIgnoresArrayOrderAndAcceptsDuplicates(t *testing.T) {
	first := upgradeSiteObservation(baselineSoftware, baselineOS, "COMPLETED")
	duplicate := first["status"].([]interface{})[0]
	first["status"] = append(first["status"].([]interface{}), duplicate)
	want, err := extractSiteUpgradeStatus(first)
	if err != nil {
		t.Fatal(err)
	}
	status := first["status"].([]interface{})
	first["status"] = []interface{}{status[1], status[0]}
	got, err := extractSiteUpgradeStatus(first)
	if err != nil {
		t.Fatal(err)
	}
	if got != want || got.SoftwareInstalledVersion != baselineSoftware || got.OSAvailableVersion != targetOS {
		t.Fatalf("status = %#v, want %#v", got, want)
	}
}

func TestExtractSiteUpgradeStatusRejectsConflictsAndMissingFields(t *testing.T) {
	conflict := upgradeSiteObservation(baselineSoftware, baselineOS, "COMPLETED")
	other := upgradeSiteObservation("crt-conflict", baselineOS, "COMPLETED")["status"].([]interface{})[0]
	conflict["status"] = append(conflict["status"].([]interface{}), other)
	if _, err := extractSiteUpgradeStatus(conflict); err == nil || !strings.Contains(err.Error(), "conflicting") {
		t.Fatalf("conflicting status error = %v", err)
	}
	missing := upgradeSiteObservation(baselineSoftware, baselineOS, "COMPLETED")
	delete(missing["spec"].(map[string]interface{}), "site_state")
	if _, err := extractSiteUpgradeStatus(missing); err == nil || !strings.Contains(err.Error(), "site_state") {
		t.Fatalf("missing status error = %v", err)
	}
	incomplete := upgradeSiteObservation(baselineSoftware, baselineOS, "COMPLETED")
	incompleteEntry := upgradeSiteObservation(baselineSoftware, baselineOS, "COMPLETED")["status"].([]interface{})[0].(map[string]interface{})
	delete(incompleteEntry["volterra_software_status"].(map[string]interface{}), "last_installed_version")
	incomplete["status"] = append(incomplete["status"].([]interface{}), incompleteEntry)
	if _, err := extractSiteUpgradeStatus(incomplete); err == nil || !strings.Contains(err.Error(), "software_installed_version") {
		t.Fatalf("partially missing status error = %v", err)
	}
}

func TestUpgradePrecheckFailuresAreSortedAndConflictsRejected(t *testing.T) {
	if _, _, err := extractFailedUpgradePrechecks(client.SMSv2Observation{"checklist": []interface{}{}}); err == nil {
		t.Fatal("empty precheck response was accepted as passing")
	}
	observation := client.SMSv2Observation{"checklist": []interface{}{
		map[string]interface{}{"item": "zeta", "status": "CHECKLIST_FAILED"},
		map[string]interface{}{"item": "alpha", "status": "CHECKLIST_UNKNOWN"},
		map[string]interface{}{"item": "capacity", "status": "CHECKLIST_WARNING"},
	}}
	failed, passing, err := extractFailedUpgradePrechecks(observation)
	if err != nil || passing || strings.Join(failed, ",") != "alpha,zeta" {
		t.Fatalf("failed=%v passing=%v err=%v", failed, passing, err)
	}
	observation["checklist"] = append(observation["checklist"].([]interface{}), map[string]interface{}{"item": "alpha", "status": "CHECKLIST_PASSED"})
	if _, _, err := extractFailedUpgradePrechecks(observation); err == nil || !strings.Contains(err.Error(), "conflicting") {
		t.Fatalf("conflicting precheck error = %v", err)
	}
}

func TestEvaluateSiteUpgradeTargetsRequiresAdvertisedTargetsAndConvergence(t *testing.T) {
	status, err := extractSiteUpgradeStatus(upgradeSiteObservation(baselineSoftware, baselineOS, "COMPLETED"))
	if err != nil {
		t.Fatal(err)
	}
	eligible, ready, converged := evaluateSiteUpgradeTargets(status, []string{targetSoftware}, true, targetSoftware, targetOS)
	if !eligible || !ready || converged {
		t.Fatalf("eligible=%v ready=%v converged=%v", eligible, ready, converged)
	}
	eligible, ready, _ = evaluateSiteUpgradeTargets(status, []string{targetSoftware}, true, targetSoftware, "9.9999.99")
	if eligible || !ready {
		t.Fatal("target eligibility and operational readiness were not evaluated independently")
	}
	convergedStatus, err := extractSiteUpgradeStatus(upgradeSiteObservation(targetSoftware, targetOS, "COMPLETED"))
	if err != nil {
		t.Fatal(err)
	}
	_, _, converged = evaluateSiteUpgradeTargets(convergedStatus, []string{targetSoftware}, true, targetSoftware, targetOS)
	if !converged {
		t.Fatal("installed targets did not converge")
	}
}

func TestEvaluateSiteUpgradeTargetsRequiresOnlineForReadiness(t *testing.T) {
	status, _ := extractSiteUpgradeStatus(upgradeSiteObservation(baselineSoftware, baselineOS, "COMPLETED"))
	status.SiteState = "OFFLINE"
	if eligible, ready, converged := evaluateSiteUpgradeTargets(status, []string{targetSoftware}, true, targetSoftware, targetOS); eligible || ready || converged {
		t.Fatalf("offline status unexpectedly passed: eligible=%v ready=%v converged=%v", eligible, ready, converged)
	}
}

func siteUpgradeReadRequest(t *testing.T, schemaResponse *datasource.SchemaResponse, wait bool, timeout int64) datasource.ReadRequest {
	t.Helper()
	model := SiteUpgradeStatusDataSourceModel{
		ID: types.StringNull(), Namespace: types.StringValue("system"), Site: types.StringValue("lab-site"),
		ExpectedSoftwareVersion: types.StringValue(targetSoftware), ExpectedOSVersion: types.StringValue(targetOS),
		Wait: types.BoolValue(wait), TimeoutSeconds: types.Int64Value(timeout), PollIntervalSeconds: types.Int64Value(1),
		SiteState: types.StringNull(), SoftwareInstalledVersion: types.StringNull(), SoftwareAvailableVersion: types.StringNull(),
		SoftwareDeploymentPhase: types.StringNull(), SoftwareDeploymentResult: types.StringNull(),
		OSInstalledVersion: types.StringNull(), OSAvailableVersion: types.StringNull(), OSDeploymentPhase: types.StringNull(), OSDeploymentResult: types.StringNull(),
		UpgradableSoftwareVersions: types.ListNull(types.StringType), FailedPrecheckNames: types.ListNull(types.StringType),
		Eligible: types.BoolNull(), Ready: types.BoolNull(), TargetConverged: types.BoolNull(),
	}
	return datasource.ReadRequest{Config: tfsdk.Config{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type())}}
}

func TestSiteUpgradeStatusPollsTransientFailedToCompleted(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"site_upgrade": "available"})
	var siteReads int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/config/namespaces/system/sites/lab-site":
			siteReads++
			software, osVersion, result := baselineSoftware, baselineOS, "Failed"
			if siteReads > 1 {
				software, osVersion, result = targetSoftware, targetOS, "Completed"
			}
			_ = json.NewEncoder(w).Encode(upgradeSiteObservation(software, osVersion, result))
		case "/api/maurice/upgradable_sw_versions":
			if request.URL.Query().Get("current_os_version") == "" || request.URL.Query().Get("current_sw_version") == "" {
				t.Error("target discovery query omitted current versions")
			}
			_, _ = w.Write([]byte(`{"sw_versions":["` + targetSoftware + `"]}`))
		case "/api/maurice/namespaces/system/sites/lab-site/pre_upgrade_check":
			if request.URL.Query().Get("sw_version") != targetSoftware {
				t.Errorf("precheck target = %q", request.URL.Query().Get("sw_version"))
			}
			_ = json.NewEncoder(w).Encode(passingPrecheck())
		case "/api/maurice/namespaces/system/sites/lab-site/upgrade_status":
			_, _ = w.Write([]byte(`{"upgrade_status":{"sw_upgrade_progress":{"status":"Completed"}}}`))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()
	source := &SiteUpgradeStatusDataSource{client: client.NewClient(server.URL, "token", client.WithMaxRetries(0)), wait: func(context.Context, time.Duration) error { return nil }}
	schemaResponse := &datasource.SchemaResponse{}
	source.Schema(context.Background(), datasource.SchemaRequest{}, schemaResponse)
	response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	source.Read(context.Background(), siteUpgradeReadRequest(t, schemaResponse, true, 5), &response)
	if response.Diagnostics.HasError() || siteReads != 2 {
		t.Fatalf("reads=%d diagnostics=%v", siteReads, response.Diagnostics)
	}
	var state SiteUpgradeStatusDataSourceModel
	response.Diagnostics.Append(response.State.Get(context.Background(), &state)...)
	if response.Diagnostics.HasError() || !state.TargetConverged.ValueBool() || state.SoftwareInstalledVersion.ValueString() != targetSoftware {
		t.Fatalf("state=%#v diagnostics=%v", state, response.Diagnostics)
	}
}

func TestSiteUpgradeStatusReportsFailedPrechecksWithoutWaiting(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"site_upgrade": "available"})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/config/namespaces/system/sites/lab-site":
			_ = json.NewEncoder(w).Encode(upgradeSiteObservation(baselineSoftware, baselineOS, "Completed"))
		case "/api/maurice/upgradable_sw_versions":
			_, _ = w.Write([]byte(`{"sw_versions":["` + targetSoftware + `"]}`))
		case "/api/maurice/namespaces/system/sites/lab-site/pre_upgrade_check":
			_, _ = w.Write([]byte(`{"checklist":[{"item":"capacity","status":"CHECKLIST_FAILED"}]}`))
		case "/api/maurice/namespaces/system/sites/lab-site/upgrade_status":
			_, _ = w.Write([]byte(`{"upgrade_status":{"sw_upgrade_progress":{"status":"Completed"}}}`))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()
	source := &SiteUpgradeStatusDataSource{client: client.NewClient(server.URL, "token", client.WithMaxRetries(0))}
	schemaResponse := &datasource.SchemaResponse{}
	source.Schema(context.Background(), datasource.SchemaRequest{}, schemaResponse)
	response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	source.Read(context.Background(), siteUpgradeReadRequest(t, schemaResponse, false, 5), &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("read diagnostics: %v", response.Diagnostics)
	}
	var state SiteUpgradeStatusDataSourceModel
	response.Diagnostics.Append(response.State.Get(context.Background(), &state)...)
	var failed []string
	response.Diagnostics.Append(state.FailedPrecheckNames.ElementsAs(context.Background(), &failed, false)...)
	if response.Diagnostics.HasError() || state.Eligible.ValueBool() || !state.Ready.ValueBool() || state.TargetConverged.ValueBool() || strings.Join(failed, ",") != "capacity" {
		t.Fatalf("state=%#v failed=%v diagnostics=%v", state, failed, response.Diagnostics)
	}
}

func TestSiteUpgradeStatusTimeoutIsSanitized(t *testing.T) {
	withSMSv2Capabilities(t, map[string]string{"site_upgrade": "available"})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/config/namespaces/system/sites/lab-site":
			_ = json.NewEncoder(w).Encode(upgradeSiteObservation(baselineSoftware, baselineOS, "Failed"))
		case "/api/maurice/upgradable_sw_versions":
			_, _ = w.Write([]byte(`{"sw_versions":["` + targetSoftware + `"]}`))
		case "/api/maurice/namespaces/system/sites/lab-site/pre_upgrade_check":
			_ = json.NewEncoder(w).Encode(passingPrecheck())
		case "/api/maurice/namespaces/system/sites/lab-site/upgrade_status":
			_, _ = w.Write([]byte(`{"upgrade_status":{"sw_upgrade_progress":{"status":"Failed","message":"secret-node https://private.invalid"}}}`))
		default:
			http.NotFound(w, request)
		}
	}))
	defer server.Close()
	nowCalls := 0
	source := &SiteUpgradeStatusDataSource{
		client: client.NewClient(server.URL, "token", client.WithMaxRetries(0)),
		wait:   func(context.Context, time.Duration) error { return nil },
		now:    func() time.Time { nowCalls++; return time.Unix(int64(nowCalls), 0) },
	}
	schemaResponse := &datasource.SchemaResponse{}
	source.Schema(context.Background(), datasource.SchemaRequest{}, schemaResponse)
	keys := make([]string, 0, len(schemaResponse.Schema.Attributes))
	for key := range schemaResponse.Schema.Attributes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	joined := strings.Join(keys, " ")
	for _, forbidden := range []string{"message", "node", "url", "raw"} {
		if strings.Contains(joined, forbidden) {
			t.Fatalf("schema exports forbidden field %q: %s", forbidden, joined)
		}
	}
	response := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	source.Read(context.Background(), siteUpgradeReadRequest(t, schemaResponse, true, 2), &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("timeout did not fail")
	}
	if text := response.Diagnostics.Errors()[0].Detail(); strings.Contains(text, "secret-node") || strings.Contains(text, "private.invalid") {
		t.Fatalf("diagnostic leaked raw API data: %q", text)
	}
}
