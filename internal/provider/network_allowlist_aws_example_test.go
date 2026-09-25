// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNetworkAllowlistAWSOriginExamplePlansWithMockProvider(t *testing.T) {
	clearCredentialEnvironment(t)
	example, err := os.ReadFile("../../examples/guides/network-allowlist-aws-origin/main.tf")
	if err != nil {
		t.Fatal(err)
	}
	config := strings.ReplaceAll(string(example), "f5-sales-demo/xcsh", "hashicorp/xcsh")
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"xcsh": providerserver.NewProtocol6WithError(New("test")()),
			"aws":  providerserver.NewProtocol6WithError(&mockAWSProvider{}),
		},
		Steps: []resource.TestStep{{Config: config, PlanOnly: true, ExpectNonEmptyPlan: true}},
	})
}

type mockAWSProvider struct{}

var _ frameworkprovider.Provider = &mockAWSProvider{}

func (p *mockAWSProvider) Metadata(_ context.Context, _ frameworkprovider.MetadataRequest, resp *frameworkprovider.MetadataResponse) {
	resp.TypeName = "aws"
	resp.Version = "test"
}

func (p *mockAWSProvider) Schema(_ context.Context, _ frameworkprovider.SchemaRequest, resp *frameworkprovider.SchemaResponse) {
	resp.Schema = providerschema.Schema{}
}

func (p *mockAWSProvider) Configure(context.Context, frameworkprovider.ConfigureRequest, *frameworkprovider.ConfigureResponse) {
}

func (p *mockAWSProvider) Resources(context.Context) []func() frameworkresource.Resource {
	return []func() frameworkresource.Resource{func() frameworkresource.Resource {
		return &mockAWSSecurityGroupIngressRule{}
	}}
}

func (p *mockAWSProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

type mockAWSSecurityGroupIngressRule struct{}

var _ frameworkresource.Resource = &mockAWSSecurityGroupIngressRule{}

type mockAWSSecurityGroupIngressRuleModel struct {
	ID              types.String `tfsdk:"id"`
	SecurityGroupID types.String `tfsdk:"security_group_id"`
	CIDRIPv4        types.String `tfsdk:"cidr_ipv4"`
	FromPort        types.Int64  `tfsdk:"from_port"`
	ToPort          types.Int64  `tfsdk:"to_port"`
	IPProtocol      types.String `tfsdk:"ip_protocol"`
	Description     types.String `tfsdk:"description"`
}

func (r *mockAWSSecurityGroupIngressRule) Metadata(_ context.Context, req frameworkresource.MetadataRequest, resp *frameworkresource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_security_group_ingress_rule"
}

func (r *mockAWSSecurityGroupIngressRule) Schema(_ context.Context, _ frameworkresource.SchemaRequest, resp *frameworkresource.SchemaResponse) {
	resp.Schema = resourceschema.Schema{Attributes: map[string]resourceschema.Attribute{
		"id":                resourceschema.StringAttribute{Computed: true},
		"security_group_id": resourceschema.StringAttribute{Required: true},
		"cidr_ipv4":         resourceschema.StringAttribute{Required: true},
		"from_port":         resourceschema.Int64Attribute{Required: true},
		"to_port":           resourceschema.Int64Attribute{Required: true},
		"ip_protocol":       resourceschema.StringAttribute{Required: true},
		"description":       resourceschema.StringAttribute{Optional: true},
	}}
}

func (r *mockAWSSecurityGroupIngressRule) Create(ctx context.Context, req frameworkresource.CreateRequest, resp *frameworkresource.CreateResponse) {
	var state mockAWSSecurityGroupIngressRuleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	state.ID = types.StringValue(state.CIDRIPv4.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *mockAWSSecurityGroupIngressRule) Read(context.Context, frameworkresource.ReadRequest, *frameworkresource.ReadResponse) {
}

func (r *mockAWSSecurityGroupIngressRule) Update(ctx context.Context, req frameworkresource.UpdateRequest, resp *frameworkresource.UpdateResponse) {
	var state mockAWSSecurityGroupIngressRuleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *mockAWSSecurityGroupIngressRule) Delete(context.Context, frameworkresource.DeleteRequest, *frameworkresource.DeleteResponse) {
}
