// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var (
	_ datasource.DataSource              = &SiteBGPStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &SiteBGPStatusDataSource{}
)

type SiteBGPStatusDataSource struct{ client *client.Client }

func NewSiteBGPStatusDataSource() datasource.DataSource { return &SiteBGPStatusDataSource{} }

type smsv2ExpectedPeerModel struct {
	Node           types.String `tfsdk:"node"`
	Role           types.String `tfsdk:"role"`
	MAC            types.String `tfsdk:"mac"`
	PeerAddress    types.String `tfsdk:"peer_address"`
	ExpectedRoutes types.Set    `tfsdk:"expected_routes"`
}

type smsv2PeerStatusModel struct {
	Node               types.String `tfsdk:"node"`
	Role               types.String `tfsdk:"role"`
	MAC                types.String `tfsdk:"mac"`
	InterfaceName      types.String `tfsdk:"interface_name"`
	PeerAddress        types.String `tfsdk:"peer_address"`
	State              types.String `tfsdk:"state"`
	ObservedAt         types.String `tfsdk:"observed_at"`
	ReceivedPrefixes   types.Int64  `tfsdk:"received_prefix_count"`
	AdvertisedPrefixes types.Int64  `tfsdk:"advertised_prefix_count"`
	Established        types.Bool   `tfsdk:"established"`
}

var smsv2ExpectedPeerAttrTypes = map[string]attr.Type{
	"node": types.StringType, "role": types.StringType, "mac": types.StringType,
	"peer_address": types.StringType, "expected_routes": types.SetType{ElemType: types.StringType},
}
var smsv2PeerStatusAttrTypes = map[string]attr.Type{
	"node": types.StringType, "role": types.StringType, "mac": types.StringType,
	"interface_name": types.StringType, "peer_address": types.StringType, "state": types.StringType,
	"observed_at": types.StringType, "received_prefix_count": types.Int64Type,
	"advertised_prefix_count": types.Int64Type, "established": types.BoolType,
}

type SiteBGPStatusDataSourceModel struct {
	ID                       types.String `tfsdk:"id"`
	Namespace                types.String `tfsdk:"namespace"`
	Site                     types.String `tfsdk:"site"`
	ExpectedPeers            types.Map    `tfsdk:"expected_peers"`
	TimeoutSeconds           types.Int64  `tfsdk:"timeout_seconds"`
	PollIntervalSeconds      types.Int64  `tfsdk:"poll_interval_seconds"`
	MaxObservationAgeSeconds types.Int64  `tfsdk:"max_observation_age_seconds"`
	Peers                    types.Map    `tfsdk:"peers"`
	BGPRoutesJSON            types.String `tfsdk:"bgp_routes_json"`
	SLORoutesJSON            types.String `tfsdk:"slo_routes_json"`
	SLIRoutesJSON            types.String `tfsdk:"sli_routes_json"`
	Converged                types.Bool   `tfsdk:"converged"`
}

type observedBGPPeer struct {
	Node, InterfaceName, PeerAddress, State, ObservedAt string
	ReceivedPrefixes, AdvertisedPrefixes                int64
}

func (d *SiteBGPStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_bgp_status"
}

func (d *SiteBGPStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Polls authoritative F5 XC BGP and route observations until they agree with MAC-bound AWS expectations.", Attributes: map[string]schema.Attribute{
		"id":        schema.StringAttribute{Computed: true},
		"namespace": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("system")}},
		"site":      schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.LengthAtLeast(1)}},
		"expected_peers": schema.MapNestedAttribute{Required: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
			"node": schema.StringAttribute{Required: true}, "role": schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("slo", "sli")}},
			"mac": schema.StringAttribute{Required: true}, "peer_address": schema.StringAttribute{Required: true},
			"expected_routes": schema.SetAttribute{Required: true, ElementType: types.StringType},
		}}},
		"timeout_seconds":             schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 1800)}},
		"poll_interval_seconds":       schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 60)}},
		"max_observation_age_seconds": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 3600)}},
		"peers": schema.MapNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
			"node": schema.StringAttribute{Computed: true}, "role": schema.StringAttribute{Computed: true}, "mac": schema.StringAttribute{Computed: true},
			"interface_name": schema.StringAttribute{Computed: true}, "peer_address": schema.StringAttribute{Computed: true},
			"state": schema.StringAttribute{Computed: true}, "observed_at": schema.StringAttribute{Computed: true},
			"received_prefix_count": schema.Int64Attribute{Computed: true}, "advertised_prefix_count": schema.Int64Attribute{Computed: true},
			"established": schema.BoolAttribute{Computed: true},
		}}},
		"bgp_routes_json": schema.StringAttribute{Computed: true}, "slo_routes_json": schema.StringAttribute{Computed: true},
		"sli_routes_json": schema.StringAttribute{Computed: true}, "converged": schema.BoolAttribute{Computed: true},
	}}
}

func (d *SiteBGPStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func peerAddress(value interface{}) string {
	if text, ok := value.(string); ok {
		return strings.TrimSpace(text)
	}
	if object, ok := value.(map[string]interface{}); ok {
		if value := stringField(object, "ipv4"); value != "" {
			return value
		}
		return stringField(object, "ipv6")
	}
	return ""
}

func extractSMSv2BGPPeers(observation client.SMSv2Observation) ([]observedBGPPeer, error) {
	rawVER, ok := observation["ver"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("BGP peer response has no ver observations")
	}
	var result []observedBGPPeer
	for _, rawNode := range rawVER {
		node, ok := rawNode.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("BGP node observation is malformed")
		}
		name := stringField(node, "name")
		rawPeers, ok := node["peer"].([]interface{})
		if !ok {
			continue
		}
		for _, rawPeer := range rawPeers {
			peer, ok := rawPeer.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("BGP peer observation is malformed")
			}
			result = append(result, observedBGPPeer{Node: name, InterfaceName: stringField(peer, "interface_name"), PeerAddress: peerAddress(peer["peer_address"]), State: stringField(peer, "protocol_status"), ObservedAt: stringField(peer, "observed_at"), ReceivedPrefixes: int64Field(peer, "received_prefix_count"), AdvertisedPrefixes: int64Field(peer, "advertised_prefix_count")})
		}
	}
	return result, nil
}

func collectRoutePrefixes(value interface{}, result map[string]struct{}) {
	switch item := value.(type) {
	case map[string]interface{}:
		for key, child := range item {
			if key == "prefix" || key == "subnet" {
				if text, ok := child.(string); ok {
					if _, _, err := net.ParseCIDR(strings.TrimSpace(text)); err == nil {
						result[strings.TrimSpace(text)] = struct{}{}
					}
				}
			}
			collectRoutePrefixes(child, result)
		}
	case []interface{}:
		for _, child := range item {
			collectRoutePrefixes(child, result)
		}
	}
}

func canonicalJSON(value interface{}) string { bytes, _ := json.Marshal(value); return string(bytes) }

func validateObservationAge(observedAt string, maximum time.Duration, now time.Time) error {
	if observedAt == "" {
		return nil
	}
	timestamp, err := time.Parse(time.RFC3339, observedAt)
	if err != nil {
		return fmt.Errorf("malformed observation timestamp %q", observedAt)
	}
	if now.Sub(timestamp) > maximum {
		return fmt.Errorf("observation timestamp %q is stale", observedAt)
	}
	return nil
}

func convergeSMSv2BGP(expected map[string]smsv2ExpectedPeerModel, configured []smsv2ConfiguredInterface, peers []observedBGPPeer, bgpRoutes, sloRoutes, sliRoutes client.SMSv2Observation, maximumAge time.Duration, now time.Time) (map[string]smsv2PeerStatusModel, bool, string) {
	if len(expected) == 0 {
		return nil, false, "expected_peers must not be empty"
	}
	configuredByMAC := map[string][]smsv2ConfiguredInterface{}
	for _, iface := range configured {
		configuredByMAC[iface.MAC] = append(configuredByMAC[iface.MAC], iface)
	}
	bgpPrefixes, sloPrefixes, sliPrefixes := map[string]struct{}{}, map[string]struct{}{}, map[string]struct{}{}
	collectRoutePrefixes(map[string]interface{}(bgpRoutes), bgpPrefixes)
	collectRoutePrefixes(map[string]interface{}(sloRoutes), sloPrefixes)
	collectRoutePrefixes(map[string]interface{}(sliRoutes), sliPrefixes)
	result := make(map[string]smsv2PeerStatusModel, len(expected))
	keys := make([]string, 0, len(expected))
	for key := range expected {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		want := expected[key]
		mac, err := normalizeSMSv2MAC(want.MAC.ValueString())
		if err != nil {
			return nil, false, fmt.Sprintf("peer %q: %v", key, err)
		}
		interfaces := configuredByMAC[mac]
		if len(interfaces) != 1 {
			return nil, false, fmt.Sprintf("peer %q MAC %s resolved to %d interfaces", key, mac, len(interfaces))
		}
		iface := interfaces[0]
		if iface.Node != want.Node.ValueString() || iface.Role != want.Role.ValueString() {
			return nil, false, fmt.Sprintf("peer %q MAC identity disagrees with configured node/role", key)
		}
		var matches []observedBGPPeer
		for _, peer := range peers {
			if peer.Node == iface.Node && peer.InterfaceName == iface.Name && peer.PeerAddress == want.PeerAddress.ValueString() {
				matches = append(matches, peer)
			}
		}
		if len(matches) != 1 {
			return nil, false, fmt.Sprintf("peer %q resolved to %d BGP observations", key, len(matches))
		}
		got := matches[0]
		if err := validateObservationAge(got.ObservedAt, maximumAge, now); err != nil {
			return nil, false, fmt.Sprintf("peer %q: %v", key, err)
		}
		if got.State != "Established" {
			return nil, false, fmt.Sprintf("peer %q state is %q", key, got.State)
		}
		var routes []string
		diags := want.ExpectedRoutes.ElementsAs(context.Background(), &routes, false)
		if diags.HasError() {
			return nil, false, fmt.Sprintf("peer %q expected routes are invalid", key)
		}
		rolePrefixes := sloPrefixes
		if iface.Role == "sli" {
			rolePrefixes = sliPrefixes
		}
		for _, route := range routes {
			if _, ok := bgpPrefixes[route]; !ok {
				return nil, false, fmt.Sprintf("peer %q BGP route %s is missing", key, route)
			}
			if _, ok := rolePrefixes[route]; !ok {
				return nil, false, fmt.Sprintf("peer %q %s route %s is missing", key, iface.Role, route)
			}
		}
		result[key] = smsv2PeerStatusModel{Node: types.StringValue(iface.Node), Role: types.StringValue(iface.Role), MAC: types.StringValue(mac), InterfaceName: types.StringValue(iface.Name), PeerAddress: types.StringValue(got.PeerAddress), State: types.StringValue(got.State), ObservedAt: types.StringValue(got.ObservedAt), ReceivedPrefixes: types.Int64Value(got.ReceivedPrefixes), AdvertisedPrefixes: types.Int64Value(got.AdvertisedPrefixes), Established: types.BoolValue(true)}
	}
	return result, true, ""
}

func (d *SiteBGPStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SiteBGPStatusDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.TimeoutSeconds.IsNull() || data.TimeoutSeconds.IsUnknown() {
		data.TimeoutSeconds = types.Int64Value(300)
	}
	if data.PollIntervalSeconds.IsNull() || data.PollIntervalSeconds.IsUnknown() {
		data.PollIntervalSeconds = types.Int64Value(10)
	}
	if data.MaxObservationAgeSeconds.IsNull() || data.MaxObservationAgeSeconds.IsUnknown() {
		data.MaxObservationAgeSeconds = types.Int64Value(120)
	}
	expected := map[string]smsv2ExpectedPeerModel{}
	resp.Diagnostics.Append(data.ExpectedPeers.ElementsAs(ctx, &expected, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	seenMAC := map[string]string{}
	for key, peer := range expected {
		mac, err := normalizeSMSv2MAC(peer.MAC.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Invalid BGP Peer Identity", fmt.Sprintf("peer %q: %v", key, err))
			return
		}
		if old, exists := seenMAC[mac]; exists {
			resp.Diagnostics.AddError("Duplicate BGP Peer MAC", fmt.Sprintf("peers %q and %q use MAC %s", old, key, mac))
			return
		}
		seenMAC[mac] = key
		peer.MAC = types.StringValue(mac)
		expected[key] = peer
	}
	configuration, err := d.client.GetSMSv2Configuration(ctx, data.Namespace.ValueString(), data.Site.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("SMSv2 Configuration Read Failed", err.Error())
		return
	}
	configured, err := extractSMSv2ConfiguredInterfaces(configuration)
	if err != nil {
		resp.Diagnostics.AddError("Invalid SMSv2 Configuration Observation", err.Error())
		return
	}
	deadline := time.Now().Add(time.Duration(data.TimeoutSeconds.ValueInt64()) * time.Second)
	var peerStatus map[string]smsv2PeerStatusModel
	var peers, bgpRoutes, sloRoutes, sliRoutes client.SMSv2Observation
	reason := "no observations"
	for {
		peers, err = d.client.GetSMSv2BGPPeers(ctx, data.Namespace.ValueString(), data.Site.ValueString())
		if err == nil {
			bgpRoutes, err = d.client.GetSMSv2BGPRoutes(ctx, data.Namespace.ValueString(), data.Site.ValueString())
		}
		if err == nil {
			sloRoutes, err = d.client.GetSMSv2SimplifiedRoutes(ctx, data.Namespace.ValueString(), data.Site.ValueString(), "slo")
		}
		if err == nil {
			sliRoutes, err = d.client.GetSMSv2SimplifiedRoutes(ctx, data.Namespace.ValueString(), data.Site.ValueString(), "sli")
		}
		if err == nil {
			var observed []observedBGPPeer
			observed, err = extractSMSv2BGPPeers(peers)
			if err == nil {
				var converged bool
				peerStatus, converged, reason = convergeSMSv2BGP(expected, configured, observed, bgpRoutes, sloRoutes, sliRoutes, time.Duration(data.MaxObservationAgeSeconds.ValueInt64())*time.Second, time.Now())
				if converged {
					break
				}
			}
		}
		if err != nil {
			reason = err.Error()
		}
		if !time.Now().Before(deadline) {
			resp.Diagnostics.AddError("SMSv2 BGP Convergence Timed Out", reason)
			return
		}
		select {
		case <-ctx.Done():
			resp.Diagnostics.AddError("SMSv2 BGP Convergence Canceled", ctx.Err().Error())
			return
		case <-time.After(time.Duration(data.PollIntervalSeconds.ValueInt64()) * time.Second):
		}
	}
	value, diags := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: smsv2PeerStatusAttrTypes}, peerStatus)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ID = types.StringValue(data.Namespace.ValueString() + "/" + data.Site.ValueString())
	data.Peers = value
	data.Converged = types.BoolValue(true)
	data.BGPRoutesJSON = types.StringValue(canonicalJSON(bgpRoutes))
	data.SLORoutesJSON = types.StringValue(canonicalJSON(sloRoutes))
	data.SLIRoutesJSON = types.StringValue(canonicalJSON(sliRoutes))
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
