// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/netip"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource                   = &networkAllowlistDataSource{}
	_ datasource.DataSourceWithValidateConfig = &networkAllowlistDataSource{}
)

type networkAllowlistManifest struct {
	APIReleaseTag string                                  `json:"api_release_tag"`
	SourceURL     string                                  `json:"source_url"`
	GeneratedAt   string                                  `json:"generated_at"`
	SourceSHA256  string                                  `json:"source_sha256"`
	Services      map[string]networkAllowlistServiceGroup `json:"services"`
	CustomerEdge  networkAllowlistCustomerEdge            `json:"customer_edge"`
}

type networkAllowlistServiceGroup struct {
	SourceEntries []string                                `json:"source_entries,omitempty"`
	Domains       []string                                `json:"domains,omitempty"`
	Regions       map[string]networkAllowlistServiceGroup `json:"regions,omitempty"`
}

type networkAllowlistCustomerEdge struct {
	Defaults     networkAllowlistCustomerEdgeDefaults `json:"defaults"`
	SecureMeshV2 networkAllowlistCustomerEdgeEgress   `json:"secure_mesh_v2"`
}

type networkAllowlistCustomerEdgeDefaults struct {
	DNSServers []string `json:"dns_servers"`
	NTPServers []string `json:"ntp_servers"`
}

type networkAllowlistCustomerEdgeEgress struct {
	RegistrationAddresses []string `json:"registration_addresses"`
	Domains               []string `json:"domains"`
}

var bundledNetworkAllowlist = mustDecodeNetworkAllowlist(bundledNetworkAllowlistJSON)

func mustDecodeNetworkAllowlist(raw string) networkAllowlistManifest {
	var manifest networkAllowlistManifest
	if err := json.Unmarshal([]byte(raw), &manifest); err != nil {
		panic("decode generated network allowlist: " + err.Error())
	}
	return manifest
}

type networkAllowlistKind int

const (
	networkAllowlistIP networkAllowlistKind = iota
	networkAllowlistRegionedIP
	networkAllowlistDomains
	networkAllowlistCEDefaults
	networkAllowlistCEEgress
)

type networkAllowlistDataSource struct {
	name        string
	group       string
	kind        networkAllowlistKind
	description string
}

func NewNetworkRegionalEdgesDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_regional_edges", "regional_edges", networkAllowlistRegionedIP, "Regional Edge IPv4 networks for origin ingress allowlists.")
}
func NewNetworkCDNDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_cdn", "cdn", networkAllowlistIP, "CDN IPv4 networks for origin or network-firewall ingress allowlists.")
}
func NewNetworkSecondaryDNSZoneTransferDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_secondary_dns_zone_transfer", "secondary_dns_zone_transfer", networkAllowlistIP, "Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their purposes.")
}
func NewNetworkGlobalLogReceiverDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_global_log_receiver", "global_log_receiver", networkAllowlistIP, "Global Log Receiver destinations. Published source entries mix CIDRs and individual IPv4 addresses.")
}
func NewNetworkDnslbHealthChecksDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_dnslb_health_checks", "dnslb_health_checks", networkAllowlistIP, "DNS Load Balancer health-check probe IPv4 addresses.")
}
func NewNetworkGlobalControllerSSOEgressDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_global_controller_sso_egress", "global_controller_sso_egress", networkAllowlistIP, "Global Controller SSO egress IPv4 addresses.")
}
func NewNetworkBotDefenseDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_bot_defense", "bot_defense", networkAllowlistDomains, "Bot Defense domains for an FQDN-aware firewall or proxy.")
}
func NewNetworkDataIntelligenceDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_data_intelligence", "data_intelligence", networkAllowlistRegionedIP, "Regional Data Intelligence IPv4 destinations.")
}
func NewNetworkCustomerEdgeDefaultsDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_customer_edge_defaults", "", networkAllowlistCEDefaults, "Default DNS and NTP destinations for Customer Edge firewall rules.")
}
func NewNetworkCustomerEdgeEgressDataSource() datasource.DataSource {
	return newNetworkAllowlistDataSource("network_customer_edge_egress", "", networkAllowlistCEEgress, "Secure Mesh v2 registration IPv4 addresses and egress domains. Legacy Customer Edge values are intentionally excluded.")
}

func newNetworkAllowlistDataSource(name, group string, kind networkAllowlistKind, description string) datasource.DataSource {
	return &networkAllowlistDataSource{name: name, group: group, kind: kind, description: description}
}

func (d *networkAllowlistDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.name
}

func (d *networkAllowlistDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attributes := map[string]schema.Attribute{
		"id":                    schema.StringAttribute{Computed: true, MarkdownDescription: "Stable identifier derived from the pinned source digest, data-source group, and selected regions."},
		"api_release_tag":       schema.StringAttribute{Computed: true, MarkdownDescription: "Pinned api-specs-enriched release tag compiled into this provider."},
		"source_url":            schema.StringAttribute{Computed: true, MarkdownDescription: "Published F5 network allowlist source URL."},
		"manifest_generated_at": schema.StringAttribute{Computed: true, MarkdownDescription: "Generation timestamp reported by the published network allowlist manifest."},
		"source_sha256":         schema.StringAttribute{Computed: true, MarkdownDescription: "SHA-256 digest of the canonical published network allowlist manifest."},
	}
	switch d.kind {
	case networkAllowlistIP:
		attributes["source_entries"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "IPv4 entries exactly as represented by the published group, preserving source order."}
		attributes["cidr_blocks"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Sorted unique IPv4 CIDRs. Individual IPv4 addresses are normalized to /32."}
	case networkAllowlistRegionedIP:
		attributes["regions"] = schema.SetAttribute{Optional: true, Computed: true, ElementType: types.StringType, MarkdownDescription: "Published regions to include. Omit to select every region."}
		attributes["source_entries"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Selected IPv4 entries exactly as represented by the published regions, ordered by region name and source order."}
		attributes["cidr_blocks"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Sorted unique IPv4 CIDRs across selected regions. Individual IPv4 addresses are normalized to /32."}
		attributes["cidr_blocks_by_region"] = schema.MapAttribute{Computed: true, ElementType: types.ListType{ElemType: types.StringType}, MarkdownDescription: "Sorted unique IPv4 CIDRs keyed by selected published region."}
	case networkAllowlistDomains:
		attributes["domains"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Published domains in source order. Leading dots are preserved and denote suffix matching where supported by the consuming firewall or proxy."}
	case networkAllowlistCEDefaults:
		attributes["dns_servers"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Published default DNS server IPv4 addresses."}
		attributes["ntp_servers"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Published default NTP server IPv4 addresses."}
	case networkAllowlistCEEgress:
		attributes["registration_addresses"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Published Secure Mesh v2 registration and update IPv4 addresses."}
		attributes["domains"] = schema.ListAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Published Secure Mesh v2 egress domains in source order. Leading dots are preserved."}
	}
	resp.Schema = schema.Schema{
		MarkdownDescription: d.description + " Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.",
		Attributes:          attributes,
	}
}

func (d *networkAllowlistDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	if d.kind != networkAllowlistRegionedIP {
		return
	}
	var configured types.Set
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("regions"), &configured)...)
	if resp.Diagnostics.HasError() || configured.IsNull() || configured.IsUnknown() {
		return
	}
	_, err := selectedNetworkAllowlistRegions(bundledNetworkAllowlist.Services[d.group].Regions, configured)
	if err != nil {
		resp.Diagnostics.AddAttributeError(path.Root("regions"), "Unknown Network Allowlist Region", err.Error())
	}
}

func (d *networkAllowlistDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	selectedRegions := []string(nil)
	if d.kind == networkAllowlistRegionedIP {
		var configured types.Set
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("regions"), &configured)...)
		if resp.Diagnostics.HasError() {
			return
		}
		var err error
		selectedRegions, err = selectedNetworkAllowlistRegions(bundledNetworkAllowlist.Services[d.group].Regions, configured)
		if err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("regions"), "Unknown Network Allowlist Region", err.Error())
			return
		}
		setValue, diags := types.SetValueFrom(ctx, types.StringType, selectedRegions)
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("regions"), setValue)...)
	}
	d.setCommonState(ctx, resp, selectedRegions)
	if resp.Diagnostics.HasError() {
		return
	}
	switch d.kind {
	case networkAllowlistIP:
		d.setIPState(ctx, resp, bundledNetworkAllowlist.Services[d.group])
	case networkAllowlistRegionedIP:
		d.setRegionedIPState(ctx, resp, selectedRegions)
	case networkAllowlistDomains:
		d.setStringList(ctx, resp, "domains", bundledNetworkAllowlist.Services[d.group].Domains)
	case networkAllowlistCEDefaults:
		d.setStringList(ctx, resp, "dns_servers", bundledNetworkAllowlist.CustomerEdge.Defaults.DNSServers)
		d.setStringList(ctx, resp, "ntp_servers", bundledNetworkAllowlist.CustomerEdge.Defaults.NTPServers)
	case networkAllowlistCEEgress:
		d.setStringList(ctx, resp, "registration_addresses", bundledNetworkAllowlist.CustomerEdge.SecureMeshV2.RegistrationAddresses)
		d.setStringList(ctx, resp, "domains", bundledNetworkAllowlist.CustomerEdge.SecureMeshV2.Domains)
	}
}

func (d *networkAllowlistDataSource) setCommonState(ctx context.Context, resp *datasource.ReadResponse, regions []string) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), stableNetworkAllowlistID(d.groupOrName(), regions))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("api_release_tag"), bundledNetworkAllowlist.APIReleaseTag)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_url"), bundledNetworkAllowlist.SourceURL)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("manifest_generated_at"), bundledNetworkAllowlist.GeneratedAt)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_sha256"), bundledNetworkAllowlist.SourceSHA256)...)
}

func (d *networkAllowlistDataSource) groupOrName() string {
	if d.group != "" {
		return d.group
	}
	return d.name
}

func (d *networkAllowlistDataSource) setIPState(ctx context.Context, resp *datasource.ReadResponse, group networkAllowlistServiceGroup) {
	cidrs, err := networkAllowlistCIDRs(group.SourceEntries)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Bundled Network Allowlist", err.Error())
		return
	}
	d.setStringList(ctx, resp, "source_entries", group.SourceEntries)
	d.setStringList(ctx, resp, "cidr_blocks", cidrs)
}

func (d *networkAllowlistDataSource) setRegionedIPState(ctx context.Context, resp *datasource.ReadResponse, regions []string) {
	group := bundledNetworkAllowlist.Services[d.group]
	var sourceEntries []string
	byRegion := make(map[string]attr.Value, len(regions))
	for _, region := range regions {
		regionGroup := group.Regions[region]
		sourceEntries = append(sourceEntries, regionGroup.SourceEntries...)
		cidrs, err := networkAllowlistCIDRs(regionGroup.SourceEntries)
		if err != nil {
			resp.Diagnostics.AddError("Invalid Bundled Network Allowlist", err.Error())
			return
		}
		value, diags := types.ListValueFrom(ctx, types.StringType, cidrs)
		resp.Diagnostics.Append(diags...)
		byRegion[region] = value
	}
	cidrs, err := networkAllowlistCIDRs(sourceEntries)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Bundled Network Allowlist", err.Error())
		return
	}
	d.setStringList(ctx, resp, "source_entries", sourceEntries)
	d.setStringList(ctx, resp, "cidr_blocks", cidrs)
	value, diags := types.MapValue(types.ListType{ElemType: types.StringType}, byRegion)
	resp.Diagnostics.Append(diags...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cidr_blocks_by_region"), value)...)
}

func (d *networkAllowlistDataSource) setStringList(ctx context.Context, resp *datasource.ReadResponse, name string, values []string) {
	value, diags := types.ListValueFrom(ctx, types.StringType, values)
	resp.Diagnostics.Append(diags...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(name), value)...)
}

func selectedNetworkAllowlistRegions(available map[string]networkAllowlistServiceGroup, configured types.Set) ([]string, error) {
	if configured.IsNull() {
		regions := make([]string, 0, len(available))
		for region := range available {
			regions = append(regions, region)
		}
		sort.Strings(regions)
		return regions, nil
	}
	if configured.IsUnknown() {
		return nil, fmt.Errorf("regions must be known before reading the bundled allowlist")
	}
	regions := make([]string, 0, len(configured.Elements()))
	for _, element := range configured.Elements() {
		value, ok := element.(types.String)
		if !ok || value.IsNull() || value.IsUnknown() {
			return nil, fmt.Errorf("regions must contain known, non-null strings")
		}
		region := value.ValueString()
		if _, ok := available[region]; !ok {
			names := make([]string, 0, len(available))
			for name := range available {
				names = append(names, name)
			}
			sort.Strings(names)
			return nil, fmt.Errorf("region %q is not published; expected one of: %s", region, strings.Join(names, ", "))
		}
		regions = append(regions, region)
	}
	sort.Strings(regions)
	return regions, nil
}

func networkAllowlistCIDRs(entries []string) ([]string, error) {
	unique := make(map[string]bool, len(entries))
	for _, entry := range entries {
		if prefix, err := netip.ParsePrefix(entry); err == nil && prefix.Addr().Is4() {
			unique[prefix.String()] = true
			continue
		}
		address, err := netip.ParseAddr(entry)
		if err != nil || !address.Is4() {
			return nil, fmt.Errorf("invalid bundled IPv4 address or CIDR %q", entry)
		}
		unique[address.String()+"/32"] = true
	}
	result := make([]string, 0, len(unique))
	for entry := range unique {
		result = append(result, entry)
	}
	sort.Strings(result)
	return result, nil
}

func stableNetworkAllowlistID(group string, regions []string) string {
	sortedRegions := append([]string(nil), regions...)
	sort.Strings(sortedRegions)
	sum := sha256.Sum256([]byte(bundledNetworkAllowlist.SourceSHA256 + "\n" + group + "\n" + strings.Join(sortedRegions, ",")))
	return "network-allowlist:" + hex.EncodeToString(sum[:])
}
