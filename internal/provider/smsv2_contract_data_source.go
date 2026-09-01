// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &Smsv2ContractDataSource{}

func NewSmsv2ContractDataSource() datasource.DataSource { return &Smsv2ContractDataSource{} }

type Smsv2ContractDataSource struct{}

type Smsv2ContractDataSourceModel struct {
	ID               types.String `tfsdk:"id"`
	ContractID       types.String `tfsdk:"contract_id"`
	ContractVersion  types.String `tfsdk:"contract_version"`
	APIReleaseTag    types.String `tfsdk:"api_release_tag"`
	APIReleaseCommit types.String `tfsdk:"api_release_commit"`
	TelemetrySchema  types.String `tfsdk:"telemetry_schema_id"`
	Capabilities     types.Map    `tfsdk:"capabilities"`
	F5XCAuthorities  types.List   `tfsdk:"f5xc_authorities"`
	AWSAuthorities   types.List   `tfsdk:"aws_authorities"`
}

func (d *Smsv2ContractDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smsv2_contract"
}

func (d *Smsv2ContractDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Publishes the immutable clean-break SMSv2 AWS TGW Connect contract compiled into this provider release.",
		Attributes: map[string]schema.Attribute{
			"id":                  schema.StringAttribute{Computed: true},
			"contract_id":         schema.StringAttribute{Computed: true},
			"contract_version":    schema.StringAttribute{Computed: true},
			"api_release_tag":     schema.StringAttribute{Computed: true},
			"api_release_commit":  schema.StringAttribute{Computed: true},
			"telemetry_schema_id": schema.StringAttribute{Computed: true},
			"capabilities":        schema.MapAttribute{Computed: true, ElementType: types.StringType},
			"f5xc_authorities":    schema.ListAttribute{Computed: true, ElementType: types.StringType},
			"aws_authorities":     schema.ListAttribute{Computed: true, ElementType: types.StringType},
		},
	}
}

func (d *Smsv2ContractDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	capabilities, diags := types.MapValueFrom(ctx, types.StringType, smsv2ContractCapabilities)
	resp.Diagnostics.Append(diags...)
	f5xc, diags := types.ListValueFrom(ctx, types.StringType, smsv2ContractF5XCAuthorities)
	resp.Diagnostics.Append(diags...)
	aws, diags := types.ListValueFrom(ctx, types.StringType, smsv2ContractAWSAuthorities)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := Smsv2ContractDataSourceModel{
		ID: types.StringValue(smsv2ContractID + "@" + smsv2ContractVersion), ContractID: types.StringValue(smsv2ContractID),
		ContractVersion: types.StringValue(smsv2ContractVersion), APIReleaseTag: types.StringValue(smsv2APIReleaseTag),
		APIReleaseCommit: types.StringValue(smsv2APIReleaseCommit), TelemetrySchema: types.StringValue(smsv2TelemetrySchemaID),
		Capabilities: capabilities, F5XCAuthorities: f5xc, AWSAuthorities: aws,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
