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

type SiteBGPStatusDataSource struct {
	client *client.Client
	now    func() time.Time
	wait   func(context.Context, time.Duration) error
}

func NewSiteBGPStatusDataSource() datasource.DataSource {
	return &SiteBGPStatusDataSource{now: time.Now, wait: waitForSMSv2Poll}
}

func waitForSMSv2Poll(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (d *SiteBGPStatusDataSource) nowTime() time.Time {
	if d.now != nil {
		return d.now()
	}
	return time.Now()
}

func (d *SiteBGPStatusDataSource) waitFor(ctx context.Context, delay time.Duration) error {
	if d.wait != nil {
		return d.wait(ctx, delay)
	}
	return waitForSMSv2Poll(ctx, delay)
}

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
	StateChangedAt     types.String `tfsdk:"state_changed_at"`
	ReceivedPrefixes   types.Int64  `tfsdk:"received_prefix_count"`
	AdvertisedPrefixes types.Int64  `tfsdk:"advertised_prefix_count"`
	Established        types.Bool   `tfsdk:"established"`
}

var smsv2PeerStatusAttrTypes = map[string]attr.Type{
	"node": types.StringType, "role": types.StringType, "mac": types.StringType,
	"interface_name": types.StringType, "peer_address": types.StringType, "state": types.StringType,
	"state_changed_at": types.StringType, "received_prefix_count": types.Int64Type,
	"advertised_prefix_count": types.Int64Type, "established": types.BoolType,
}

type SiteBGPStatusDataSourceModel struct {
	ID                  types.String `tfsdk:"id"`
	Namespace           types.String `tfsdk:"namespace"`
	Site                types.String `tfsdk:"site"`
	ExpectedPeers       types.Map    `tfsdk:"expected_peers"`
	TimeoutSeconds      types.Int64  `tfsdk:"timeout_seconds"`
	PollIntervalSeconds types.Int64  `tfsdk:"poll_interval_seconds"`
	Peers               types.Map    `tfsdk:"peers"`
	BGPRoutesJSON       types.String `tfsdk:"bgp_routes_json"`
	SLORoutesJSON       types.String `tfsdk:"slo_routes_json"`
	SLIRoutesJSON       types.String `tfsdk:"sli_routes_json"`
	Converged           types.Bool   `tfsdk:"converged"`
}

type observedBGPPeer struct {
	Node, InterfaceName, PeerAddress, State, StateChangedAt string
	ReceivedPrefixes, AdvertisedPrefixes                    int64
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
		"timeout_seconds":       schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 1800)}},
		"poll_interval_seconds": schema.Int64Attribute{Optional: true, Computed: true, Validators: []validator.Int64{int64validator.Between(1, 60)}},
		"peers": schema.MapNestedAttribute{Computed: true, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
			"node": schema.StringAttribute{Computed: true}, "role": schema.StringAttribute{Computed: true}, "mac": schema.StringAttribute{Computed: true},
			"interface_name": schema.StringAttribute{Computed: true}, "peer_address": schema.StringAttribute{Computed: true},
			"state": schema.StringAttribute{Computed: true}, "state_changed_at": schema.StringAttribute{Computed: true},
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
		for _, family := range []string{"ipv4", "ipv6"} {
			value, exists := object[family]
			if !exists {
				continue
			}
			if text, ok := value.(string); ok {
				if address := strings.TrimSpace(text); address != "" {
					return address
				}
			}
			if address, ok := value.(map[string]interface{}); ok {
				if text := stringField(address, "addr"); text != "" {
					return text
				}
			}
		}
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
			result = append(result, observedBGPPeer{Node: name, InterfaceName: stringField(peer, "interface_name"), PeerAddress: peerAddress(peer["peer_address"]), State: stringField(peer, "protocol_status"), StateChangedAt: stringField(peer, "up_down_timestamp"), ReceivedPrefixes: int64Field(peer, "received_prefix_count"), AdvertisedPrefixes: int64Field(peer, "advertised_prefix_count")})
		}
	}
	return result, nil
}

func routeArray(value map[string]interface{}, key string) ([]interface{}, error) {
	routes, ok := value[key].([]interface{})
	if !ok {
		return nil, fmt.Errorf("route response field %q is missing or malformed", key)
	}
	return routes, nil
}

func addRouteSubnets(routes []interface{}, node string, result map[string]map[string]struct{}) error {
	if result[node] == nil {
		result[node] = map[string]struct{}{}
	}
	for _, rawRoute := range routes {
		route, ok := rawRoute.(map[string]interface{})
		if !ok {
			return fmt.Errorf("route observation for node %q is malformed", node)
		}
		prefix := stringField(route, "subnet")
		if _, _, err := net.ParseCIDR(prefix); err != nil {
			return fmt.Errorf("route observation for node %q has invalid subnet %q", node, prefix)
		}
		result[node][prefix] = struct{}{}
	}
	return nil
}

func extractBGPRoutePrefixesForDirections(observation client.SMSv2Observation, directions ...string) (map[string]map[string]struct{}, error) {
	result := map[string]map[string]struct{}{}
	rawNodes, ok := observation["ver"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("BGP route response has no ver observations")
	}
	for _, rawNode := range rawNodes {
		node, ok := rawNode.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("BGP route node observation is malformed")
		}
		name := stringField(node, "name")
		if name == "" {
			return nil, fmt.Errorf("BGP route node observation has no name")
		}
		routingInstances, err := routeArray(node, "ri_table")
		if err != nil {
			return nil, err
		}
		for _, rawInstance := range routingInstances {
			instance, ok := rawInstance.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("BGP routing instance for node %q is malformed", name)
			}
			tables, err := routeArray(instance, "rt_table")
			if err != nil {
				return nil, err
			}
			for _, rawTable := range tables {
				table, ok := rawTable.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("BGP route table for node %q is malformed", name)
				}
				for _, key := range directions {
					routes, err := routeArray(table, key)
					if err != nil {
						return nil, err
					}
					if err := addRouteSubnets(routes, name, result); err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return result, nil
}

func extractBGPRoutePrefixes(observation client.SMSv2Observation) (map[string]map[string]struct{}, error) {
	return extractBGPRoutePrefixesForDirections(observation, "imported", "exported")
}

func extractSimplifiedRoutePrefixes(observation client.SMSv2Observation) (map[string]map[string]struct{}, error) {
	result := map[string]map[string]struct{}{}
	rawNodes, ok := observation["ver_routes"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("simplified route response has no ver_routes observations")
	}
	for _, rawNode := range rawNodes {
		node, ok := rawNode.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("simplified route node observation is malformed")
		}
		name := stringField(node, "node")
		if name == "" {
			return nil, fmt.Errorf("simplified route observation has no node")
		}
		routes, err := routeArray(node, "route")
		if err != nil {
			return nil, err
		}
		if result[name] == nil {
			result[name] = map[string]struct{}{}
		}
		for _, rawRoute := range routes {
			route, ok := rawRoute.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("simplified route for node %q is malformed", name)
			}
			prefix := stringField(route, "prefix")
			if _, _, err := net.ParseCIDR(prefix); err != nil {
				return nil, fmt.Errorf("simplified route for node %q has invalid prefix %q", name, prefix)
			}
			result[name][prefix] = struct{}{}
		}
	}
	return result, nil
}

func canonicalJSON(value interface{}) string { bytes, _ := json.Marshal(value); return string(bytes) }

func routePrefixesForNode(prefixes map[string]map[string]struct{}, configuredNode, observation string) (map[string]struct{}, error) {
	matches := make([]map[string]struct{}, 0, 1)
	for observedNode, routes := range prefixes {
		if smsv2NodeMatches(configuredNode, observedNode) {
			matches = append(matches, routes)
		}
	}
	if len(matches) != 1 {
		return nil, fmt.Errorf("%s node %q resolved to %d observations", observation, configuredNode, len(matches))
	}
	return matches[0], nil
}

func convergeSMSv2BGP(expected map[string]smsv2ExpectedPeerModel, configured []smsv2ConfiguredInterface, peers []observedBGPPeer, bgpRoutes, sloRoutes, sliRoutes client.SMSv2Observation) (map[string]smsv2PeerStatusModel, bool, string) {
	if len(expected) == 0 {
		return nil, false, "expected_peers must not be empty"
	}
	configuredByIdentity := map[string][]smsv2ConfiguredInterface{}
	for _, iface := range configured {
		configuredByIdentity[iface.Node+"\x00"+iface.MAC] = append(configuredByIdentity[iface.Node+"\x00"+iface.MAC], iface)
	}
	bgpPrefixes, err := extractBGPRoutePrefixes(bgpRoutes)
	if err != nil {
		return nil, false, err.Error()
	}
	bgpExportedPrefixes, err := extractBGPRoutePrefixesForDirections(bgpRoutes, "exported")
	if err != nil {
		return nil, false, err.Error()
	}
	sloPrefixes, err := extractSimplifiedRoutePrefixes(sloRoutes)
	if err != nil {
		return nil, false, err.Error()
	}
	sliPrefixes, err := extractSimplifiedRoutePrefixes(sliRoutes)
	if err != nil {
		return nil, false, err.Error()
	}
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
		interfaces := configuredByIdentity[strings.TrimSpace(want.Node.ValueString())+"\x00"+mac]
		if len(interfaces) != 1 {
			return nil, false, fmt.Sprintf("peer %q MAC %s resolved to %d interfaces", key, mac, len(interfaces))
		}
		iface := interfaces[0]
		if iface.Node != want.Node.ValueString() {
			return nil, false, fmt.Sprintf("peer %q MAC identity disagrees with configured node", key)
		}
		var matches []observedBGPPeer
		for _, peer := range peers {
			// API v6 defines BGP observation identity as node plus peer address.
			// interface_name is retained as observed status, but it is not the
			// SMSv2 network-interface object name and must not be used as identity.
			if smsv2NodeMatches(iface.Node, peer.Node) && peer.PeerAddress == want.PeerAddress.ValueString() {
				matches = append(matches, peer)
			}
		}
		if len(matches) != 1 {
			return nil, false, fmt.Sprintf("peer %q resolved to %d BGP observations", key, len(matches))
		}
		got := matches[0]
		if got.State != "Established" {
			return nil, false, fmt.Sprintf("peer %q state is %q", key, got.State)
		}
		nodeBGPPrefixes, err := routePrefixesForNode(bgpPrefixes, iface.Node, "BGP route")
		if err != nil {
			return nil, false, err.Error()
		}
		nodeExportedPrefixes, err := routePrefixesForNode(bgpExportedPrefixes, iface.Node, "exported BGP route")
		if err != nil {
			return nil, false, err.Error()
		}
		if len(nodeExportedPrefixes) == 0 {
			return nil, false, fmt.Sprintf("peer %q has no exported BGP routes on node %q", key, iface.Node)
		}
		routeRole := want.Role.ValueString()
		rolePrefixSets := sloPrefixes
		if routeRole == "sli" {
			rolePrefixSets = sliPrefixes
		}
		rolePrefixes, err := routePrefixesForNode(rolePrefixSets, iface.Node, routeRole+" route")
		if err != nil {
			return nil, false, err.Error()
		}
		var routes []string
		diags := want.ExpectedRoutes.ElementsAs(context.Background(), &routes, false)
		if diags.HasError() {
			return nil, false, fmt.Sprintf("peer %q expected routes are invalid", key)
		}
		for _, route := range routes {
			if _, ok := nodeBGPPrefixes[route]; !ok {
				return nil, false, fmt.Sprintf("peer %q BGP route %s is missing on node %q", key, route, iface.Node)
			}
		}
		for route := range nodeExportedPrefixes {
			if _, ok := rolePrefixes[route]; !ok {
				return nil, false, fmt.Sprintf("peer %q exported BGP route %s is missing from the %s route view on node %q", key, route, routeRole, iface.Node)
			}
		}
		result[key] = smsv2PeerStatusModel{Node: types.StringValue(iface.Node), Role: types.StringValue(routeRole), MAC: types.StringValue(mac), InterfaceName: types.StringValue(got.InterfaceName), PeerAddress: types.StringValue(got.PeerAddress), State: types.StringValue(got.State), StateChangedAt: types.StringValue(got.StateChangedAt), ReceivedPrefixes: types.Int64Value(got.ReceivedPrefixes), AdvertisedPrefixes: types.Int64Value(got.AdvertisedPrefixes), Established: types.BoolValue(true)}
	}
	return result, true, ""
}

func (d *SiteBGPStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SiteBGPStatusDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Namespace.IsUnknown() || data.Site.IsUnknown() || data.ExpectedPeers.IsUnknown() ||
		data.TimeoutSeconds.IsUnknown() || data.PollIntervalSeconds.IsUnknown() {
		if req.ClientCapabilities.DeferralAllowed {
			resp.Deferred = &datasource.Deferred{Reason: datasource.DeferredReasonDataSourceConfigUnknown}
		}
		return
	}
	if err := requireSMSv2Capabilities("runtime_status", "tgw_connect"); err != nil {
		resp.Diagnostics.AddError("SMSv2 BGP Status Unavailable", err.Error())
		return
	}
	if data.TimeoutSeconds.IsNull() || data.TimeoutSeconds.IsUnknown() {
		data.TimeoutSeconds = types.Int64Value(300)
	}
	if data.PollIntervalSeconds.IsNull() || data.PollIntervalSeconds.IsUnknown() {
		data.PollIntervalSeconds = types.Int64Value(10)
	}
	expected := map[string]smsv2ExpectedPeerModel{}
	resp.Diagnostics.Append(data.ExpectedPeers.ElementsAs(ctx, &expected, true)...)
	if resp.Diagnostics.HasError() {
		return
	}
	seenMAC := map[string]string{}
	for key, peer := range expected {
		if peer.Node.IsUnknown() || peer.Role.IsUnknown() || peer.MAC.IsUnknown() ||
			peer.PeerAddress.IsUnknown() || peer.ExpectedRoutes.IsUnknown() {
			if req.ClientCapabilities.DeferralAllowed {
				resp.Deferred = &datasource.Deferred{Reason: datasource.DeferredReasonDataSourceConfigUnknown}
			}
			return
		}
		if peer.Node.IsNull() || peer.Role.IsNull() || peer.MAC.IsNull() ||
			peer.PeerAddress.IsNull() || peer.ExpectedRoutes.IsNull() {
			resp.Diagnostics.AddError("Invalid BGP Peer Identity", fmt.Sprintf("peer %q is incomplete", key))
			return
		}
		node := strings.TrimSpace(peer.Node.ValueString())
		role := strings.TrimSpace(peer.Role.ValueString())
		if node == "" || (role != "slo" && role != "sli") {
			resp.Diagnostics.AddError("Invalid BGP Peer Identity", fmt.Sprintf("peer %q has invalid node or role", key))
			return
		}
		mac, err := normalizeSMSv2MAC(peer.MAC.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Invalid BGP Peer Identity", fmt.Sprintf("peer %q: %v", key, err))
			return
		}
		identity := node + "\x00" + mac
		if old, exists := seenMAC[identity]; exists {
			resp.Diagnostics.AddError("Duplicate BGP Peer MAC", fmt.Sprintf("peers %q and %q use MAC %s within node %q", old, key, mac, node))
			return
		}
		seenMAC[identity] = key
		peer.Node = types.StringValue(node)
		peer.Role = types.StringValue(role)
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
	deadline := d.nowTime().Add(time.Duration(data.TimeoutSeconds.ValueInt64()) * time.Second)
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
				peerStatus, converged, reason = convergeSMSv2BGP(expected, configured, observed, bgpRoutes, sloRoutes, sliRoutes)
				if converged {
					break
				}
			}
		}
		if err != nil {
			reason = err.Error()
		}
		remaining := deadline.Sub(d.nowTime())
		if remaining <= 0 {
			resp.Diagnostics.AddError("SMSv2 BGP Convergence Timed Out", reason)
			return
		}
		delay := time.Duration(data.PollIntervalSeconds.ValueInt64()) * time.Second
		if delay > remaining {
			delay = remaining
		}
		if err := d.waitFor(ctx, delay); err != nil {
			resp.Diagnostics.AddError("SMSv2 BGP Convergence Canceled", err.Error())
			return
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
