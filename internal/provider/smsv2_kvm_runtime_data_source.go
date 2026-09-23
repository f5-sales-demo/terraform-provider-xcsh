// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"errors"
	"fmt"
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
	_ datasource.DataSource              = &Smsv2KvmRuntimeDataSource{}
	_ datasource.DataSourceWithConfigure = &Smsv2KvmRuntimeDataSource{}
)

type Smsv2KvmRuntimeDataSource struct {
	client *client.Client
	now    func() time.Time
	wait   func(context.Context, time.Duration) error
}

type Smsv2KvmRuntimeDataSourceModel struct {
	ID                  types.String `tfsdk:"id"`
	Namespace           types.String `tfsdk:"namespace"`
	Site                types.String `tfsdk:"site"`
	ExpectedMAC         types.String `tfsdk:"expected_mac"`
	TimeoutSeconds      types.Int64  `tfsdk:"timeout_seconds"`
	PollIntervalSeconds types.Int64  `tfsdk:"poll_interval_seconds"`
	InterfaceName       types.String `tfsdk:"interface_name"`
	Hostname            types.String `tfsdk:"hostname"`
	Device              types.String `tfsdk:"device"`
	MAC                 types.String `tfsdk:"mac"`
	RegistrationState   types.String `tfsdk:"registration_state"`
	Online              types.Bool   `tfsdk:"online"`
}

type smsv2KvmRuntimeIdentity struct {
	InterfaceName     string
	Hostname          string
	Device            string
	MAC               string
	RegistrationState string
}

type kvmRuntimePendingError struct{ reason string }

func (e *kvmRuntimePendingError) Error() string { return e.reason }

func NewSmsv2KvmRuntimeDataSource() datasource.DataSource {
	return &Smsv2KvmRuntimeDataSource{now: time.Now, wait: waitForSMSv2Poll}
}

func (d *Smsv2KvmRuntimeDataSource) nowTime() time.Time {
	if d.now != nil {
		return d.now()
	}
	return time.Now()
}

func (d *Smsv2KvmRuntimeDataSource) waitFor(ctx context.Context, delay time.Duration) error {
	if d.wait != nil {
		return d.wait(ctx, delay)
	}
	return waitForSMSv2Poll(ctx, delay)
}

func (d *Smsv2KvmRuntimeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smsv2_kvm_runtime"
}

func (d *Smsv2KvmRuntimeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resolves one realized KVM Secure Mesh Site v2 SLO network interface through site UID ownership, the live registration hostname and device, and an expected MAC address. The name is observed, never guessed.",
		Attributes: map[string]schema.Attribute{
			"id":                    schema.StringAttribute{Computed: true, MarkdownDescription: "Stable lookup identity composed from the site and normalized expected MAC."},
			"namespace":             schema.StringAttribute{Optional: true, Computed: true, Validators: []validator.String{stringvalidator.OneOf("system")}, MarkdownDescription: "Namespace containing the site and realized interface. Defaults to `system`."},
			"site":                  schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}, MarkdownDescription: "Secure Mesh Site v2 configuration name."},
			"expected_mac":          schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}, MarkdownDescription: "Terraform-owned KVM CE MAC used to select one registration device."},
			"timeout_seconds":       schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 7200)}, MarkdownDescription: "Bounded runtime discovery timeout. Defaults to 7200 seconds."},
			"poll_interval_seconds": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 60)}, MarkdownDescription: "Polling interval. Defaults to 10 seconds."},
			"interface_name":        schema.StringAttribute{Computed: true, MarkdownDescription: "Exact owned XC network_interface object name."},
			"hostname":              schema.StringAttribute{Computed: true, MarkdownDescription: "Hostname reported by the matching live KVM registration."},
			"device":                schema.StringAttribute{Computed: true, MarkdownDescription: "Guest device reported by the matching live KVM registration."},
			"mac":                   schema.StringAttribute{Computed: true, MarkdownDescription: "Normalized six-octet MAC address."},
			"registration_state":    schema.StringAttribute{Computed: true, MarkdownDescription: "Current state of the matching registration."},
			"online":                schema.BoolAttribute{Computed: true, MarkdownDescription: "Whether the matching registration currently reports ONLINE."},
		},
	}
}

func (d *Smsv2KvmRuntimeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected *client.Client")
		return
	}
	d.client = c
}

func kvmRegistrationProvider(infra client.RegistrationInfra) string {
	if strings.TrimSpace(infra.ProviderRef) != "" {
		return strings.ToUpper(strings.TrimSpace(infra.ProviderRef))
	}
	return strings.ToUpper(strings.TrimSpace(infra.Provider))
}

func resolveSMSv2KVMRuntime(
	configuration client.SMSv2Observation,
	registrations *client.RegistrationListResponse,
	interfaces client.SMSv2Observation,
	expectedMAC string,
) (smsv2KvmRuntimeIdentity, error) {
	mac, err := normalizeSMSv2MAC(expectedMAC)
	if err != nil {
		return smsv2KvmRuntimeIdentity{}, err
	}
	metadata, _ := nestedMap(map[string]interface{}(configuration), "metadata")
	system, _ := nestedMap(map[string]interface{}(configuration), "system_metadata")
	site, namespace, uid := stringField(metadata, "name"), stringField(metadata, "namespace"), stringField(system, "uid")
	if site == "" || namespace == "" || uid == "" {
		return smsv2KvmRuntimeIdentity{}, &kvmRuntimePendingError{"site name, namespace, or UID is not yet available"}
	}
	if registrations == nil {
		return smsv2KvmRuntimeIdentity{}, fmt.Errorf("registration discovery returned no response")
	}
	if len(registrations.Errors) != 0 {
		return smsv2KvmRuntimeIdentity{}, fmt.Errorf("registration discovery returned partial errors")
	}
	type registrationMatch struct {
		hostname string
		device   string
		state    string
	}
	matches := []registrationMatch{}
	for _, registration := range registrations.Items {
		if isTerminalRegistrationState(registration.Object.Status.CurrentState) {
			continue
		}
		if registration.GetSpec.Passport.ClusterName != site || kvmRegistrationProvider(registration.GetSpec.Infra) != "KVM" {
			continue
		}
		hostname := strings.TrimSpace(registration.GetSpec.Infra.Hostname)
		for _, network := range registration.GetSpec.Infra.HWInfo.Network {
			observed, normalizeErr := normalizeSMSv2MAC(network.MACAddress)
			if normalizeErr != nil || observed != mac {
				continue
			}
			device := strings.TrimSpace(network.Name)
			if hostname == "" || device == "" {
				return smsv2KvmRuntimeIdentity{}, fmt.Errorf("matching KVM registration has incomplete hostname or device identity")
			}
			matches = append(matches, registrationMatch{hostname: hostname, device: device, state: registration.Object.Status.CurrentState})
		}
	}
	if len(matches) == 0 {
		return smsv2KvmRuntimeIdentity{}, &kvmRuntimePendingError{fmt.Sprintf("expected one live KVM registration interface for MAC %s; observed 0 live matches", mac)}
	}
	if len(matches) != 1 {
		return smsv2KvmRuntimeIdentity{}, fmt.Errorf("expected one live KVM registration interface for MAC %s; observed %d", mac, len(matches))
	}
	if rawErrors, exists := interfaces["errors"]; exists && rawErrors != nil {
		entries, ok := rawErrors.([]interface{})
		if !ok || len(entries) != 0 {
			return smsv2KvmRuntimeIdentity{}, fmt.Errorf("network_interface discovery returned partial errors")
		}
	}
	items, ok := interfaces["items"].([]interface{})
	if !ok {
		return smsv2KvmRuntimeIdentity{}, fmt.Errorf("network_interface discovery has no items array")
	}
	objectNames := []string{}
	for _, raw := range items {
		item, itemOK := raw.(map[string]interface{})
		if !itemOK {
			return smsv2KvmRuntimeIdentity{}, fmt.Errorf("network_interface discovery contains a malformed item")
		}
		owner, _ := nestedMap(item, "owner_view")
		if len(owner) == 0 {
			owner, _ = nestedMap(item, "system_metadata", "owner_view")
		}
		if stringField(owner, "kind") != "securemesh_site_v2" || stringField(owner, "name") != site || stringField(owner, "namespace") != namespace || stringField(owner, "uid") != uid {
			continue
		}
		ethernet, _ := nestedMap(item, "get_spec", "ethernet_interface")
		if stringField(ethernet, "node") != matches[0].hostname || stringField(ethernet, "device") != matches[0].device {
			continue
		}
		_, slo := ethernet["site_local_network"]
		_, sli := ethernet["site_local_inside_network"]
		name := stringField(item, "name")
		if name == "" || stringField(item, "namespace") != namespace || !slo || sli {
			return smsv2KvmRuntimeIdentity{}, fmt.Errorf("owned KVM network_interface has invalid identity or role")
		}
		objectNames = append(objectNames, name)
	}
	if len(objectNames) == 0 {
		return smsv2KvmRuntimeIdentity{}, &kvmRuntimePendingError{fmt.Sprintf("owned KVM network_interface for node %q device %q is not yet available", matches[0].hostname, matches[0].device)}
	}
	if len(objectNames) != 1 {
		return smsv2KvmRuntimeIdentity{}, fmt.Errorf("node %q device %q resolved to %d owned network_interface objects", matches[0].hostname, matches[0].device, len(objectNames))
	}
	return smsv2KvmRuntimeIdentity{InterfaceName: objectNames[0], Hostname: matches[0].hostname, Device: matches[0].device, MAC: mac, RegistrationState: matches[0].state}, nil
}

func (d *Smsv2KvmRuntimeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Smsv2KvmRuntimeDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Site.IsUnknown() || data.ExpectedMAC.IsUnknown() || data.Namespace.IsUnknown() {
		if req.ClientCapabilities.DeferralAllowed {
			resp.Deferred = &datasource.Deferred{Reason: datasource.DeferredReasonDataSourceConfigUnknown}
		}
		return
	}
	namespace := data.Namespace.ValueString()
	if data.Namespace.IsNull() || namespace == "" {
		namespace = "system"
		data.Namespace = types.StringValue(namespace)
	}
	mac, err := normalizeSMSv2MAC(data.ExpectedMAC.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid KVM Runtime Identity", err.Error())
		return
	}
	data.ExpectedMAC = types.StringValue(mac)
	if data.TimeoutSeconds.IsNull() || data.TimeoutSeconds.IsUnknown() {
		data.TimeoutSeconds = types.Int64Value(7200)
	}
	if data.PollIntervalSeconds.IsNull() || data.PollIntervalSeconds.IsUnknown() {
		data.PollIntervalSeconds = types.Int64Value(10)
	}
	deadline := d.nowTime().Add(time.Duration(data.TimeoutSeconds.ValueInt64()) * time.Second)
	var lastReason string
	for {
		configuration, readErr := d.client.GetSMSv2Configuration(ctx, namespace, data.Site.ValueString())
		var registrations *client.RegistrationListResponse
		var interfaces client.SMSv2Observation
		if readErr == nil {
			registrations, readErr = d.client.ListRegistrationsBySite(ctx, namespace, data.Site.ValueString())
		}
		if readErr == nil {
			interfaces, readErr = d.client.ListSMSv2NetworkInterfaces(ctx, namespace)
		}
		if readErr != nil {
			lastReason = "KVM runtime observations are not yet available"
		} else if identity, resolveErr := resolveSMSv2KVMRuntime(configuration, registrations, interfaces, mac); resolveErr == nil {
			data.ID = types.StringValue(data.Site.ValueString() + "/" + identity.MAC)
			data.InterfaceName = types.StringValue(identity.InterfaceName)
			data.Hostname = types.StringValue(identity.Hostname)
			data.Device = types.StringValue(identity.Device)
			data.MAC = types.StringValue(identity.MAC)
			data.RegistrationState = stringOrNull(identity.RegistrationState)
			data.Online = types.BoolValue(identity.RegistrationState == "ONLINE")
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			var pending *kvmRuntimePendingError
			if !errors.As(resolveErr, &pending) {
				resp.Diagnostics.AddError("KVM Runtime Discovery Failed", resolveErr.Error())
				return
			}
			lastReason = pending.Error()
		}
		remaining := deadline.Sub(d.nowTime())
		if remaining <= 0 {
			resp.Diagnostics.AddError("KVM Runtime Discovery Timed Out", lastReason)
			return
		}
		if waitErr := d.waitFor(ctx, min(time.Duration(data.PollIntervalSeconds.ValueInt64())*time.Second, remaining)); waitErr != nil {
			resp.Diagnostics.AddError("KVM Runtime Discovery Interrupted", waitErr.Error())
			return
		}
	}
}
