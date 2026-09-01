// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var (
	_ datasource.DataSource              = &Smsv2AWSRuntimeDataSource{}
	_ datasource.DataSourceWithConfigure = &Smsv2AWSRuntimeDataSource{}
)

type Smsv2AWSRuntimeDataSource struct{ client *client.Client }

func NewSmsv2AWSRuntimeDataSource() datasource.DataSource { return &Smsv2AWSRuntimeDataSource{} }

type smsv2BindingModel struct {
	Node types.String `tfsdk:"node"`
	Role types.String `tfsdk:"role"`
	MAC  types.String `tfsdk:"mac"`
}

type smsv2RuntimeInterfaceModel struct {
	Node          types.String `tfsdk:"node"`
	Role          types.String `tfsdk:"role"`
	MAC           types.String `tfsdk:"mac"`
	InterfaceName types.String `tfsdk:"interface_name"`
	MTU           types.Int64  `tfsdk:"mtu"`
	Healthy       types.Bool   `tfsdk:"healthy"`
}

var smsv2RuntimeInterfaceAttrTypes = map[string]attr.Type{
	"node": types.StringType, "role": types.StringType, "mac": types.StringType,
	"interface_name": types.StringType, "mtu": types.Int64Type, "healthy": types.BoolType,
}

type Smsv2AWSRuntimeDataSourceModel struct {
	ID         types.String `tfsdk:"id"`
	Namespace  types.String `tfsdk:"namespace"`
	Site       types.String `tfsdk:"site"`
	Nodes      types.Map    `tfsdk:"nodes"`
	Interfaces types.Map    `tfsdk:"interfaces"`
	Healthy    types.Bool   `tfsdk:"healthy"`
}

type smsv2ConfiguredInterface struct {
	Node string
	Role string
	MAC  string
	Name string
	MTU  int64
}

func (d *Smsv2AWSRuntimeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smsv2_aws_runtime"
}

func (d *Smsv2AWSRuntimeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Correlates AWS ENI MAC identities with authoritative SMSv2 configuration and node health.", Attributes: map[string]schema.Attribute{
		"id":        schema.StringAttribute{Computed: true},
		"namespace": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("system")}},
		"site":      schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
		"nodes": schema.MapNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
			"node": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
			"role": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("slo", "sli")}},
			"mac":  schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
		}}},
		"interfaces": schema.MapNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
			"node": schema.StringAttribute{Computed: true}, "role": schema.StringAttribute{Computed: true},
			"mac": schema.StringAttribute{Computed: true}, "interface_name": schema.StringAttribute{Computed: true},
			"mtu": schema.Int64Attribute{Computed: true}, "healthy": schema.BoolAttribute{Computed: true},
		}}},
		"healthy": schema.BoolAttribute{Computed: true},
	}}
}

func (d *Smsv2AWSRuntimeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func normalizeSMSv2MAC(value string) (string, error) {
	parsed, err := net.ParseMAC(strings.TrimSpace(value))
	if err != nil || len(parsed) != 6 {
		return "", fmt.Errorf("%q is not a six-octet IEEE 802 MAC address", value)
	}
	return strings.ToLower(parsed.String()), nil
}

func nestedMap(value interface{}, keys ...string) (map[string]interface{}, bool) {
	current, ok := value.(map[string]interface{})
	if !ok {
		return nil, false
	}
	for _, key := range keys {
		next, found := current[key]
		if !found {
			return nil, false
		}
		current, ok = next.(map[string]interface{})
		if !ok {
			return nil, false
		}
	}
	return current, true
}

func stringField(value map[string]interface{}, key string) string {
	text, _ := value[key].(string)
	return strings.TrimSpace(text)
}

func int64Field(value map[string]interface{}, key string) int64 {
	switch number := value[key].(type) {
	case float64:
		return int64(number)
	case int64:
		return number
	case int:
		return int64(number)
	default:
		return 0
	}
}

func extractSMSv2ConfiguredInterfaces(configuration client.SMSv2Observation) ([]smsv2ConfiguredInterface, error) {
	spec, ok := nestedMap(map[string]interface{}(configuration), "spec")
	if !ok {
		return nil, fmt.Errorf("SMSv2 configuration response has no spec object")
	}
	notManaged, ok := nestedMap(spec, "aws", "not_managed")
	if !ok {
		return nil, fmt.Errorf("SMSv2 configuration is not aws.not_managed")
	}
	rawNodes, ok := notManaged["node_list"].([]interface{})
	if !ok || len(rawNodes) == 0 {
		return nil, fmt.Errorf("SMSv2 configuration has no AWS nodes")
	}
	var result []smsv2ConfiguredInterface
	seen := map[string]string{}
	for _, rawNode := range rawNodes {
		node, ok := rawNode.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("SMSv2 node observation is malformed")
		}
		hostname := stringField(node, "hostname")
		if hostname == "" {
			return nil, fmt.Errorf("SMSv2 node observation has no hostname")
		}
		rawInterfaces, ok := node["interface_list"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("SMSv2 node %q has no interface observations", hostname)
		}
		for _, rawInterface := range rawInterfaces {
			iface, ok := rawInterface.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("SMSv2 node %q has a malformed interface", hostname)
			}
			ethernet, ok := nestedMap(iface, "ethernet_interface")
			if !ok {
				return nil, fmt.Errorf("SMSv2 node %q has an interface without ethernet identity", hostname)
			}
			mac, err := normalizeSMSv2MAC(stringField(ethernet, "mac"))
			if err != nil {
				return nil, fmt.Errorf("SMSv2 node %q: %w", hostname, err)
			}
			role := ""
			network, _ := nestedMap(iface, "network_option")
			if network != nil {
				if _, exists := network["site_local_network"]; exists {
					role = "slo"
				}
				if _, exists := network["site_local_inside_network"]; exists {
					if role != "" {
						return nil, fmt.Errorf("SMSv2 MAC %s has ambiguous roles", mac)
					}
					role = "sli"
				}
			}
			if role == "" {
				return nil, fmt.Errorf("SMSv2 MAC %s has no SLO/SLI role", mac)
			}
			if previous, exists := seen[mac]; exists {
				return nil, fmt.Errorf("SMSv2 MAC %s is duplicated on nodes %q and %q", mac, previous, hostname)
			}
			seen[mac] = hostname
			name := stringField(iface, "name")
			if name == "" {
				return nil, fmt.Errorf("SMSv2 MAC %s has no authoritative interface name", mac)
			}
			mtu := int64Field(iface, "mtu")
			if mtu <= 0 {
				return nil, fmt.Errorf("SMSv2 MAC %s has invalid MTU %d", mac, mtu)
			}
			result = append(result, smsv2ConfiguredInterface{Node: hostname, Role: role, MAC: mac, Name: name, MTU: mtu})
		}
	}
	return result, nil
}

func validateSMSv2Bindings(bindings map[string]smsv2BindingModel) (map[string]smsv2BindingModel, error) {
	if len(bindings) == 0 {
		return nil, fmt.Errorf("nodes must contain at least one MAC-bound interface")
	}
	seen := map[string]string{}
	for key, binding := range bindings {
		if strings.TrimSpace(key) == "" || binding.Node.IsNull() || binding.Role.IsNull() || binding.MAC.IsNull() {
			return nil, fmt.Errorf("binding %q is incomplete", key)
		}
		role := binding.Role.ValueString()
		if role != "slo" && role != "sli" {
			return nil, fmt.Errorf("binding %q has unsupported role %q", key, role)
		}
		mac, err := normalizeSMSv2MAC(binding.MAC.ValueString())
		if err != nil {
			return nil, fmt.Errorf("binding %q: %w", key, err)
		}
		if previous, exists := seen[mac]; exists {
			return nil, fmt.Errorf("bindings %q and %q use duplicate MAC %s", previous, key, mac)
		}
		seen[mac] = key
		binding.MAC = types.StringValue(mac)
		bindings[key] = binding
	}
	return bindings, nil
}

func correlateSMSv2Runtime(bindings map[string]smsv2BindingModel, configured []smsv2ConfiguredInterface, health client.SMSv2Observation) (map[string]smsv2RuntimeInterfaceModel, error) {
	byMAC := map[string][]smsv2ConfiguredInterface{}
	configuredNodes := map[string]struct{}{}
	for _, iface := range configured {
		byMAC[iface.MAC] = append(byMAC[iface.MAC], iface)
		configuredNodes[iface.Node] = struct{}{}
	}
	healthNode, healthState := stringField(health, "hostname"), strings.ToUpper(stringField(health, "state"))
	if healthNode == "" || healthState == "" {
		return nil, fmt.Errorf("SMSv2 node health observation is incomplete")
	}
	if _, exists := configuredNodes[healthNode]; !exists {
		return nil, fmt.Errorf("SMSv2 node health node %q is not configured for this site", healthNode)
	}
	siteHealthy := healthState == "PROVISIONED"
	result := make(map[string]smsv2RuntimeInterfaceModel, len(bindings))
	for key, expected := range bindings {
		matches := byMAC[expected.MAC.ValueString()]
		if len(matches) != 1 {
			return nil, fmt.Errorf("binding %q MAC %s resolved to %d configured interfaces", key, expected.MAC.ValueString(), len(matches))
		}
		got := matches[0]
		if got.Node != expected.Node.ValueString() || got.Role != expected.Role.ValueString() {
			return nil, fmt.Errorf("binding %q disagrees with F5 XC configuration: got node=%q role=%q", key, got.Node, got.Role)
		}
		result[key] = smsv2RuntimeInterfaceModel{Node: types.StringValue(got.Node), Role: types.StringValue(got.Role), MAC: types.StringValue(got.MAC), InterfaceName: types.StringValue(got.Name), MTU: types.Int64Value(got.MTU), Healthy: types.BoolValue(siteHealthy)}
	}
	return result, nil
}

func (d *Smsv2AWSRuntimeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Smsv2AWSRuntimeDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	bindings := map[string]smsv2BindingModel{}
	resp.Diagnostics.Append(data.Nodes.ElementsAs(ctx, &bindings, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	bindings, err := validateSMSv2Bindings(bindings)
	if err != nil {
		resp.Diagnostics.AddError("Invalid SMSv2 Runtime Identity", err.Error())
		return
	}
	configuration, err := d.client.GetSMSv2Configuration(ctx, data.Namespace.ValueString(), data.Site.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("SMSv2 Configuration Read Failed", err.Error())
		return
	}
	health, err := d.client.GetSMSv2Health(ctx, data.Site.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("SMSv2 Health Read Failed", err.Error())
		return
	}
	configured, err := extractSMSv2ConfiguredInterfaces(configuration)
	if err != nil {
		resp.Diagnostics.AddError("Invalid SMSv2 Configuration Observation", err.Error())
		return
	}
	interfaces, err := correlateSMSv2Runtime(bindings, configured, health)
	if err != nil {
		resp.Diagnostics.AddError("Inconsistent SMSv2 Runtime Observation", err.Error())
		return
	}
	healthy := true
	for _, iface := range interfaces {
		healthy = healthy && iface.Healthy.ValueBool()
	}
	value, diags := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: smsv2RuntimeInterfaceAttrTypes}, interfaces)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ID = types.StringValue(data.Namespace.ValueString() + "/" + data.Site.ValueString())
	data.Interfaces = value
	data.Healthy = types.BoolValue(healthy)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
