// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &Smsv2ContractDataSource{}
var _ datasource.DataSourceWithValidateConfig = &Smsv2ContractDataSource{}

func NewSmsv2ContractDataSource() datasource.DataSource { return &Smsv2ContractDataSource{} }

type Smsv2ContractDataSource struct{}

type Smsv2ContractDataSourceModel struct {
	ID                           types.String `tfsdk:"id"`
	RequiredCapabilities         types.Set    `tfsdk:"required_capabilities"`
	ContractID                   types.String `tfsdk:"contract_id"`
	ContractVersion              types.String `tfsdk:"contract_version"`
	APIReleaseTag                types.String `tfsdk:"api_release_tag"`
	APIReleaseCommit             types.String `tfsdk:"api_release_commit"`
	TelemetrySchema              types.String `tfsdk:"telemetry_schema_id"`
	Capabilities                 types.Map    `tfsdk:"capabilities"`
	F5XCAuthorities              types.List   `tfsdk:"f5xc_authorities"`
	AWSAuthorities               types.List   `tfsdk:"aws_authorities"`
	AzureRouteServerEBGPMultihop types.Object `tfsdk:"azure_route_server_ebgp_multihop"`
}

type smsv2CapabilitySourceContract struct {
	Repository  string
	Commit      string
	AssetPath   string
	AssetSHA256 string
	SchemaPaths []string
}

type smsv2CapabilityBoundaryContract struct {
	Availability string
	Enforcement  string
	Reason       string
	Source       smsv2CapabilitySourceContract
}

var smsv2CapabilitySourceAttrTypes = map[string]attr.Type{
	"repository":   types.StringType,
	"commit":       types.StringType,
	"asset_path":   types.StringType,
	"asset_sha256": types.StringType,
	"schema_paths": types.ListType{ElemType: types.StringType},
}

var smsv2CapabilityBoundaryAttrTypes = map[string]attr.Type{
	"availability": types.StringType,
	"enforcement":  types.StringType,
	"reason":       types.StringType,
	"source":       types.ObjectType{AttrTypes: smsv2CapabilitySourceAttrTypes},
}

func (d *Smsv2ContractDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smsv2_contract"
}

func (d *Smsv2ContractDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Publishes the immutable clean-break SMSv2 AWS and Azure capability contract compiled into this provider release.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true},
			"required_capabilities": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Capabilities that must be available. A known unavailable capability produces a planning diagnostic before any F5 API request.",
			},
			"contract_id":         schema.StringAttribute{Computed: true},
			"contract_version":    schema.StringAttribute{Computed: true},
			"api_release_tag":     schema.StringAttribute{Computed: true},
			"api_release_commit":  schema.StringAttribute{Computed: true},
			"telemetry_schema_id": schema.StringAttribute{Computed: true},
			"capabilities":        schema.MapAttribute{Computed: true, ElementType: types.StringType},
			"f5xc_authorities":    schema.ListAttribute{Computed: true, ElementType: types.StringType},
			"aws_authorities":     schema.ListAttribute{Computed: true, ElementType: types.StringType},
			"azure_route_server_ebgp_multihop": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Authoritative Azure Route Server eBGP multihop availability and immutable source provenance.",
				Attributes: map[string]schema.Attribute{
					"availability": schema.StringAttribute{Computed: true},
					"enforcement":  schema.StringAttribute{Computed: true},
					"reason":       schema.StringAttribute{Computed: true},
					"source": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"repository":   schema.StringAttribute{Computed: true},
							"commit":       schema.StringAttribute{Computed: true},
							"asset_path":   schema.StringAttribute{Computed: true},
							"asset_sha256": schema.StringAttribute{Computed: true},
							"schema_paths": schema.ListAttribute{Computed: true, ElementType: types.StringType},
						},
					},
				},
			},
		},
	}
}

func (d *Smsv2ContractDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
	var config Smsv2ContractDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() || config.RequiredCapabilities.IsNull() || config.RequiredCapabilities.IsUnknown() {
		return
	}
	var names []string
	for _, value := range config.RequiredCapabilities.Elements() {
		name, ok := value.(types.String)
		if !ok || name.IsUnknown() {
			return
		}
		if name.IsNull() {
			resp.Diagnostics.AddError("Invalid Required SMSv2 Capability", "required_capabilities cannot contain null values")
			return
		}
		names = append(names, name.ValueString())
	}
	sort.Strings(names)
	if err := validateRequiredSMSv2Capabilities(smsv2ContractCapabilities, smsv2AzureRouteServerEBGPMultihop, smsv2APIReleaseTag, names...); err != nil {
		resp.Diagnostics.AddError("Required SMSv2 Capability Unavailable", err.Error())
	}
}

func (d *Smsv2ContractDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	state := Smsv2ContractDataSourceModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	capabilities, diags := types.MapValueFrom(ctx, types.StringType, smsv2ContractCapabilities)
	resp.Diagnostics.Append(diags...)
	f5xc, diags := types.ListValueFrom(ctx, types.StringType, smsv2ContractF5XCAuthorities)
	resp.Diagnostics.Append(diags...)
	aws, diags := types.ListValueFrom(ctx, types.StringType, smsv2ContractAWSAuthorities)
	resp.Diagnostics.Append(diags...)
	sourcePaths, diags := types.ListValueFrom(ctx, types.StringType, smsv2AzureRouteServerEBGPMultihop.Source.SchemaPaths)
	resp.Diagnostics.Append(diags...)
	source, diags := types.ObjectValue(smsv2CapabilitySourceAttrTypes, map[string]attr.Value{
		"repository":   types.StringValue(smsv2AzureRouteServerEBGPMultihop.Source.Repository),
		"commit":       types.StringValue(smsv2AzureRouteServerEBGPMultihop.Source.Commit),
		"asset_path":   types.StringValue(smsv2AzureRouteServerEBGPMultihop.Source.AssetPath),
		"asset_sha256": types.StringValue(smsv2AzureRouteServerEBGPMultihop.Source.AssetSHA256),
		"schema_paths": sourcePaths,
	})
	resp.Diagnostics.Append(diags...)
	azureMultihop, diags := types.ObjectValue(smsv2CapabilityBoundaryAttrTypes, map[string]attr.Value{
		"availability": types.StringValue(smsv2AzureRouteServerEBGPMultihop.Availability),
		"enforcement":  types.StringValue(smsv2AzureRouteServerEBGPMultihop.Enforcement),
		"reason":       types.StringValue(smsv2AzureRouteServerEBGPMultihop.Reason),
		"source":       source,
	})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state.ID = types.StringValue(smsv2ContractID + "@" + smsv2ContractVersion)
	state.ContractID = types.StringValue(smsv2ContractID)
	state.ContractVersion = types.StringValue(smsv2ContractVersion)
	state.APIReleaseTag = types.StringValue(smsv2APIReleaseTag)
	state.APIReleaseCommit = types.StringValue(smsv2SourceCommit)
	state.TelemetrySchema = types.StringValue(smsv2TelemetrySchemaID)
	state.Capabilities = capabilities
	state.F5XCAuthorities = f5xc
	state.AWSAuthorities = aws
	state.AzureRouteServerEBGPMultihop = azureMultihop
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func requireSMSv2Capabilities(names ...string) error {
	return validateSMSv2Capabilities(smsv2ContractCapabilities, smsv2APIReleaseTag, names...)
}

func validateSMSv2Capabilities(capabilities map[string]string, release string, names ...string) error {
	var unavailable []string
	for _, name := range names {
		if capabilities[name] != "available" {
			unavailable = append(unavailable, name)
		}
	}
	if len(unavailable) != 0 {
		return fmt.Errorf("API release %s does not make SMSv2 capabilities available: %s", release, strings.Join(unavailable, ", "))
	}
	return nil
}

func validateRequiredSMSv2Capabilities(capabilities map[string]string, azure smsv2CapabilityBoundaryContract, release string, names ...string) error {
	for _, name := range names {
		if name == "azure_route_server_ebgp_multihop" {
			if azure.Availability != "available" {
				return fmt.Errorf(
					"API release %s declares required capability %q %s with enforcement %q: %s (source %s@%s %s %s)",
					release, name, azure.Availability, azure.Enforcement, azure.Reason,
					azure.Source.Repository, azure.Source.Commit, azure.Source.AssetPath, azure.Source.AssetSHA256,
				)
			}
			continue
		}
		availability, known := capabilities[name]
		if !known {
			return fmt.Errorf("API release %s does not declare required SMSv2 capability %q", release, name)
		}
		if availability != "available" {
			return fmt.Errorf("API release %s declares required SMSv2 capability %q %s", release, name, availability)
		}
	}
	return nil
}
