// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var (
	_ datasource.DataSource              = &SiteUpgradeStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &SiteUpgradeStatusDataSource{}
)

type SiteUpgradeStatusDataSource struct {
	client *client.Client
	now    func() time.Time
	wait   func(context.Context, time.Duration) error
}

func NewSiteUpgradeStatusDataSource() datasource.DataSource {
	return &SiteUpgradeStatusDataSource{now: time.Now, wait: waitForSMSv2Poll}
}

type SiteUpgradeStatusDataSourceModel struct {
	ID                         types.String `tfsdk:"id"`
	Namespace                  types.String `tfsdk:"namespace"`
	Site                       types.String `tfsdk:"site"`
	ExpectedSoftwareVersion    types.String `tfsdk:"expected_software_version"`
	ExpectedOSVersion          types.String `tfsdk:"expected_os_version"`
	Wait                       types.Bool   `tfsdk:"wait"`
	TimeoutSeconds             types.Int64  `tfsdk:"timeout_seconds"`
	PollIntervalSeconds        types.Int64  `tfsdk:"poll_interval_seconds"`
	SiteState                  types.String `tfsdk:"site_state"`
	SoftwareInstalledVersion   types.String `tfsdk:"software_installed_version"`
	SoftwareAvailableVersion   types.String `tfsdk:"software_available_version"`
	SoftwareDeploymentPhase    types.String `tfsdk:"software_deployment_phase"`
	SoftwareDeploymentResult   types.String `tfsdk:"software_deployment_result"`
	OSInstalledVersion         types.String `tfsdk:"os_installed_version"`
	OSAvailableVersion         types.String `tfsdk:"os_available_version"`
	OSDeploymentPhase          types.String `tfsdk:"os_deployment_phase"`
	OSDeploymentResult         types.String `tfsdk:"os_deployment_result"`
	UpgradableSoftwareVersions types.List   `tfsdk:"upgradable_software_versions"`
	FailedPrecheckNames        types.List   `tfsdk:"failed_precheck_names"`
	Eligible                   types.Bool   `tfsdk:"eligible"`
	Ready                      types.Bool   `tfsdk:"ready"`
	TargetConverged            types.Bool   `tfsdk:"target_converged"`
}

type siteUpgradeStatus struct {
	SiteState                string
	SoftwareInstalledVersion string
	SoftwareAvailableVersion string
	SoftwareDeploymentPhase  string
	SoftwareDeploymentResult string
	OSInstalledVersion       string
	OSAvailableVersion       string
	OSDeploymentPhase        string
	OSDeploymentResult       string
}

type siteUpgradeSnapshot struct {
	status             siteUpgradeStatus
	upgradableSoftware []string
	failedPrechecks    []string
	prechecksPassing   bool
	eligible           bool
	ready              bool
	converged          bool
}

func (d *SiteUpgradeStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_upgrade_status"
}

func (d *SiteUpgradeStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Observes SMSv2 site upgrade eligibility and waits for explicitly supplied software and operating-system targets to converge.",
		Attributes: map[string]schema.Attribute{
			"id":                           schema.StringAttribute{Computed: true},
			"namespace":                    schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("system")}},
			"site":                         schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
			"expected_software_version":    schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
			"expected_os_version":          schema.StringAttribute{Optional: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
			"wait":                         schema.BoolAttribute{Optional: true, Computed: true},
			"timeout_seconds":              schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 7200)}},
			"poll_interval_seconds":        schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 300)}},
			"site_state":                   schema.StringAttribute{Computed: true},
			"software_installed_version":   schema.StringAttribute{Computed: true},
			"software_available_version":   schema.StringAttribute{Computed: true},
			"software_deployment_phase":    schema.StringAttribute{Computed: true},
			"software_deployment_result":   schema.StringAttribute{Computed: true},
			"os_installed_version":         schema.StringAttribute{Computed: true},
			"os_available_version":         schema.StringAttribute{Computed: true},
			"os_deployment_phase":          schema.StringAttribute{Computed: true},
			"os_deployment_result":         schema.StringAttribute{Computed: true},
			"upgradable_software_versions": schema.ListAttribute{Computed: true, ElementType: types.StringType},
			"failed_precheck_names":        schema.ListAttribute{Computed: true, ElementType: types.StringType},
			"eligible":                     schema.BoolAttribute{Computed: true},
			"ready":                        schema.BoolAttribute{Computed: true, MarkdownDescription: "Whether the site is operationally ready (`ONLINE`), independent of target eligibility."},
			"target_converged":             schema.BoolAttribute{Computed: true},
		},
	}
}

func (d *SiteUpgradeStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	configured, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected *client.Client")
		return
	}
	d.client = configured
}

func (d *SiteUpgradeStatusDataSource) nowTime() time.Time {
	if d.now != nil {
		return d.now()
	}
	return time.Now()
}

func (d *SiteUpgradeStatusDataSource) waitFor(ctx context.Context, delay time.Duration) error {
	if d.wait != nil {
		return d.wait(ctx, delay)
	}
	return waitForSMSv2Poll(ctx, delay)
}

func statusConsensus(observation client.SMSv2Observation, label string, path ...string) (string, error) {
	rawStatuses, ok := observation["status"].([]interface{})
	if !ok || len(rawStatuses) == 0 {
		return "", fmt.Errorf("site status array is missing")
	}
	values := map[string]struct{}{}
	for _, rawStatus := range rawStatuses {
		current, ok := rawStatus.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("site status entry is malformed")
		}
		for _, key := range path[:len(path)-1] {
			next, exists := current[key].(map[string]interface{})
			if !exists {
				current = nil
				break
			}
			current = next
		}
		if current == nil {
			return "", fmt.Errorf("site status field %s is missing", label)
		}
		value := strings.TrimSpace(stringField(current, path[len(path)-1]))
		if value == "" {
			return "", fmt.Errorf("site status field %s is missing", label)
		}
		values[value] = struct{}{}
	}
	if len(values) == 0 {
		return "", fmt.Errorf("site status field %s is missing", label)
	}
	if len(values) != 1 {
		return "", fmt.Errorf("site status field %s has conflicting observations", label)
	}
	for value := range values {
		return value, nil
	}
	return "", fmt.Errorf("site status field %s is missing", label)
}

func extractSiteUpgradeStatus(observation client.SMSv2Observation) (siteUpgradeStatus, error) {
	spec, ok := observation["spec"].(map[string]interface{})
	if !ok || stringField(spec, "site_state") == "" {
		return siteUpgradeStatus{}, fmt.Errorf("site status field site_state is missing")
	}
	result := siteUpgradeStatus{SiteState: stringField(spec, "site_state")}
	fields := []struct {
		label string
		path  []string
		set   func(string)
	}{
		{"software_installed_version", []string{"volterra_software_status", "last_installed_version"}, func(value string) { result.SoftwareInstalledVersion = value }},
		{"software_available_version", []string{"volterra_software_status", "available_version"}, func(value string) { result.SoftwareAvailableVersion = value }},
		{"software_deployment_phase", []string{"volterra_software_status", "deployment_state", "phase"}, func(value string) { result.SoftwareDeploymentPhase = value }},
		{"software_deployment_result", []string{"volterra_software_status", "deployment_state", "result"}, func(value string) { result.SoftwareDeploymentResult = value }},
		{"os_installed_version", []string{"operating_system_status", "deployment_state", "version"}, func(value string) { result.OSInstalledVersion = value }},
		{"os_available_version", []string{"operating_system_status", "available_version"}, func(value string) { result.OSAvailableVersion = value }},
		{"os_deployment_phase", []string{"operating_system_status", "deployment_state", "phase"}, func(value string) { result.OSDeploymentPhase = value }},
		{"os_deployment_result", []string{"operating_system_status", "deployment_state", "result"}, func(value string) { result.OSDeploymentResult = value }},
	}
	for _, field := range fields {
		value, err := statusConsensus(observation, field.label, field.path...)
		if err != nil {
			return siteUpgradeStatus{}, err
		}
		field.set(value)
	}
	return result, nil
}

func extractUpgradableSoftwareVersions(observation client.SMSv2Observation) ([]string, error) {
	rawVersions, ok := observation["sw_versions"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("software target list is missing")
	}
	unique := map[string]struct{}{}
	for _, rawVersion := range rawVersions {
		version, ok := rawVersion.(string)
		version = strings.TrimSpace(version)
		if !ok || version == "" {
			return nil, fmt.Errorf("software target list is malformed")
		}
		unique[version] = struct{}{}
	}
	versions := make([]string, 0, len(unique))
	for version := range unique {
		versions = append(versions, version)
	}
	sort.Strings(versions)
	return versions, nil
}

func extractFailedUpgradePrechecks(observation client.SMSv2Observation) ([]string, bool, error) {
	rawChecks, ok := observation["checklist"].([]interface{})
	if !ok {
		return nil, false, fmt.Errorf("software precheck list is missing")
	}
	if len(rawChecks) == 0 {
		return nil, false, fmt.Errorf("software precheck list is empty")
	}
	statuses := map[string]string{}
	for _, rawCheck := range rawChecks {
		check, ok := rawCheck.(map[string]interface{})
		if !ok {
			return nil, false, fmt.Errorf("software precheck entry is malformed")
		}
		name, status := stringField(check, "item"), stringField(check, "status")
		if name == "" || status == "" {
			return nil, false, fmt.Errorf("software precheck entry is incomplete")
		}
		if previous, exists := statuses[name]; exists && previous != status {
			return nil, false, fmt.Errorf("software precheck %q has conflicting observations", name)
		}
		statuses[name] = status
	}
	failed := make([]string, 0)
	for name, status := range statuses {
		if status != "CHECKLIST_PASSED" && status != "CHECKLIST_WARNING" {
			failed = append(failed, name)
		}
	}
	sort.Strings(failed)
	return failed, len(failed) == 0, nil
}

func validateUpgradeProgress(observation client.SMSv2Observation) error {
	root, ok := nestedMap(map[string]interface{}(observation), "upgrade_status", "sw_upgrade_progress")
	if !ok || stringField(root, "status") == "" {
		return fmt.Errorf("upgrade progress status is missing")
	}
	return nil
}

func containsString(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}

func evaluateSiteUpgradeTargets(status siteUpgradeStatus, upgradable []string, prechecksPassing bool, expectedSoftware, expectedOS string) (bool, bool, bool) {
	softwareTarget := strings.TrimSpace(expectedSoftware)
	if softwareTarget == "" {
		softwareTarget = status.SoftwareAvailableVersion
	}
	osTarget := strings.TrimSpace(expectedOS)
	if osTarget == "" {
		osTarget = status.OSAvailableVersion
	}
	eligible := status.SiteState == "ONLINE" && softwareTarget != "" && osTarget != "" &&
		containsString(upgradable, softwareTarget) && status.OSAvailableVersion == osTarget && prechecksPassing
	converged := status.SiteState == "ONLINE" &&
		(expectedSoftware == "" || status.SoftwareInstalledVersion == expectedSoftware) &&
		(expectedOS == "" || status.OSInstalledVersion == expectedOS)
	return eligible, status.SiteState == "ONLINE", converged
}

func (d *SiteUpgradeStatusDataSource) observe(ctx context.Context, namespace, site, expectedSoftware, expectedOS string) (siteUpgradeSnapshot, error) {
	rawStatus, err := d.client.GetSMSv2SiteUpgradeStatus(ctx, namespace, site)
	if err != nil {
		return siteUpgradeSnapshot{}, fmt.Errorf("site upgrade status request failed")
	}
	status, err := extractSiteUpgradeStatus(rawStatus)
	if err != nil {
		return siteUpgradeSnapshot{}, err
	}
	rawTargets, err := d.client.GetSMSv2UpgradableSoftwareVersions(ctx, status.OSInstalledVersion, status.SoftwareInstalledVersion)
	if err != nil {
		return siteUpgradeSnapshot{}, fmt.Errorf("software target discovery request failed")
	}
	targets, err := extractUpgradableSoftwareVersions(rawTargets)
	if err != nil {
		return siteUpgradeSnapshot{}, err
	}
	precheckTarget := strings.TrimSpace(expectedSoftware)
	if precheckTarget == "" {
		precheckTarget = status.SoftwareAvailableVersion
	}
	rawPrechecks, err := d.client.GetSMSv2PreUpgradeCheck(ctx, namespace, site, precheckTarget)
	if err != nil {
		return siteUpgradeSnapshot{}, fmt.Errorf("software precheck request failed")
	}
	failed, passing, err := extractFailedUpgradePrechecks(rawPrechecks)
	if err != nil {
		return siteUpgradeSnapshot{}, err
	}
	progress, err := d.client.GetSMSv2UpgradeProgress(ctx, namespace, site)
	if err != nil {
		return siteUpgradeSnapshot{}, fmt.Errorf("upgrade progress request failed")
	}
	if err := validateUpgradeProgress(progress); err != nil {
		return siteUpgradeSnapshot{}, err
	}
	eligible, ready, converged := evaluateSiteUpgradeTargets(status, targets, passing, expectedSoftware, expectedOS)
	return siteUpgradeSnapshot{status: status, upgradableSoftware: targets, failedPrechecks: failed, prechecksPassing: passing, eligible: eligible, ready: ready, converged: converged}, nil
}

func (d *SiteUpgradeStatusDataSource) setState(ctx context.Context, data *SiteUpgradeStatusDataSourceModel, snapshot siteUpgradeSnapshot, resp *datasource.ReadResponse) {
	versions, diags := types.ListValueFrom(ctx, types.StringType, snapshot.upgradableSoftware)
	resp.Diagnostics.Append(diags...)
	failed, diags := types.ListValueFrom(ctx, types.StringType, snapshot.failedPrechecks)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ID = types.StringValue(data.Namespace.ValueString() + "/" + data.Site.ValueString())
	data.SiteState = types.StringValue(snapshot.status.SiteState)
	data.SoftwareInstalledVersion = types.StringValue(snapshot.status.SoftwareInstalledVersion)
	data.SoftwareAvailableVersion = types.StringValue(snapshot.status.SoftwareAvailableVersion)
	data.SoftwareDeploymentPhase = types.StringValue(snapshot.status.SoftwareDeploymentPhase)
	data.SoftwareDeploymentResult = types.StringValue(snapshot.status.SoftwareDeploymentResult)
	data.OSInstalledVersion = types.StringValue(snapshot.status.OSInstalledVersion)
	data.OSAvailableVersion = types.StringValue(snapshot.status.OSAvailableVersion)
	data.OSDeploymentPhase = types.StringValue(snapshot.status.OSDeploymentPhase)
	data.OSDeploymentResult = types.StringValue(snapshot.status.OSDeploymentResult)
	data.UpgradableSoftwareVersions = versions
	data.FailedPrecheckNames = failed
	data.Eligible = types.BoolValue(snapshot.eligible)
	data.Ready = types.BoolValue(snapshot.ready)
	data.TargetConverged = types.BoolValue(snapshot.converged)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (d *SiteUpgradeStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SiteUpgradeStatusDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Namespace.IsUnknown() || data.Site.IsUnknown() || data.ExpectedSoftwareVersion.IsUnknown() ||
		data.ExpectedOSVersion.IsUnknown() || data.Wait.IsUnknown() || data.TimeoutSeconds.IsUnknown() || data.PollIntervalSeconds.IsUnknown() {
		if req.ClientCapabilities.DeferralAllowed {
			resp.Deferred = &datasource.Deferred{Reason: datasource.DeferredReasonDataSourceConfigUnknown}
		}
		return
	}
	if err := requireSMSv2Capabilities("site_upgrade"); err != nil {
		resp.Diagnostics.AddError("SMSv2 Site Upgrade Status Unavailable", err.Error())
		return
	}
	if data.Wait.IsNull() {
		data.Wait = types.BoolValue(false)
	}
	if data.TimeoutSeconds.IsNull() {
		data.TimeoutSeconds = types.Int64Value(300)
	}
	if data.PollIntervalSeconds.IsNull() {
		data.PollIntervalSeconds = types.Int64Value(10)
	}
	expectedSoftware := strings.TrimSpace(data.ExpectedSoftwareVersion.ValueString())
	expectedOS := strings.TrimSpace(data.ExpectedOSVersion.ValueString())
	if data.Wait.ValueBool() && expectedSoftware == "" && expectedOS == "" {
		resp.Diagnostics.AddError("Missing Upgrade Target", "Waiting requires an expected software version, an expected operating-system version, or both.")
		return
	}
	deadline := d.nowTime().Add(time.Duration(data.TimeoutSeconds.ValueInt64()) * time.Second)
	for {
		snapshot, err := d.observe(ctx, data.Namespace.ValueString(), data.Site.ValueString(), expectedSoftware, expectedOS)
		if err == nil && (!data.Wait.ValueBool() || snapshot.converged) {
			d.setState(ctx, &data, snapshot, resp)
			return
		}
		if !data.Wait.ValueBool() {
			resp.Diagnostics.AddError("SMSv2 Site Upgrade Observation Failed", "The sanitized site upgrade observation was unavailable or malformed.")
			return
		}
		remaining := deadline.Sub(d.nowTime())
		if remaining <= 0 {
			resp.Diagnostics.AddError("SMSv2 Site Upgrade Convergence Timed Out", "The supplied upgrade targets did not converge before the configured timeout.")
			return
		}
		delay := time.Duration(data.PollIntervalSeconds.ValueInt64()) * time.Second
		if delay > remaining {
			delay = remaining
		}
		if err := d.waitFor(ctx, delay); err != nil {
			resp.Diagnostics.AddError("SMSv2 Site Upgrade Convergence Canceled", "Waiting for the supplied upgrade targets was canceled.")
			return
		}
	}
}
