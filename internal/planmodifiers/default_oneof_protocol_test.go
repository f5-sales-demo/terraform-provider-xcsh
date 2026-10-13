package planmodifiers_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	ps "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/planmodifiers"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/validators"
)

type markerProvider struct{}

func (markerProvider) Metadata(_ context.Context, _ provider.MetadataRequest, r *provider.MetadataResponse) {
	r.TypeName = "marker"
}
func (markerProvider) Schema(_ context.Context, _ provider.SchemaRequest, r *provider.SchemaResponse) {
	r.Schema = ps.Schema{}
}
func (markerProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}
func (markerProvider) DataSources(context.Context) []func() datasource.DataSource { return nil }
func (markerProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{func() resource.Resource { return markerResource{} }}
}

// DataSources uses the distinct framework interface below.

type markerResource struct{}

func (markerResource) Metadata(_ context.Context, _ resource.MetadataRequest, r *resource.MetadataResponse) {
	r.TypeName = "marker_probe"
}
func (markerResource) Schema(_ context.Context, _ resource.SchemaRequest, r *resource.SchemaResponse) {
	r.Schema = schema.Schema{Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{Required: true}}, Blocks: map[string]schema.Block{
		"settings": schema.SingleNestedBlock{
			Validators: []validator.Object{validators.ConflictingObjectAttributes("default", "custom")},
			Attributes: map[string]schema.Attribute{
				"default": schema.ObjectAttribute{Optional: true, Computed: true, AttributeTypes: map[string]attr.Type{}, PlanModifiers: []planmodifier.Object{planmodifiers.DefaultOneOfMarker("default", []string{"default", "custom"}), objectplanmodifier.UseStateForUnknown()}},
				"custom":  schema.ObjectAttribute{Optional: true, AttributeTypes: map[string]attr.Type{}},
			},
		},
	}}
}
func (markerResource) Create(context.Context, resource.CreateRequest, *resource.CreateResponse) {}
func (markerResource) Read(context.Context, resource.ReadRequest, *resource.ReadResponse)       {}
func (markerResource) Update(context.Context, resource.UpdateRequest, *resource.UpdateResponse) {}
func (markerResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {}

func TestNestedDefaultProtocolPlanAndConflictValidation(t *testing.T) {
	ctx := context.Background()
	server := providerserver.NewProtocol6(markerProvider{})()
	empty := tftypes.Object{AttributeTypes: map[string]tftypes.Type{}}
	settings := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"default": empty, "custom": empty}}
	root := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"name": tftypes.String, "settings": settings}}
	dynamic := func(def, custom interface{}) *tfprotov6.DynamicValue {
		value := tftypes.NewValue(root, map[string]tftypes.Value{"name": tftypes.NewValue(tftypes.String, "identity-preserved"), "settings": tftypes.NewValue(settings, map[string]tftypes.Value{"default": tftypes.NewValue(empty, def), "custom": tftypes.NewValue(empty, custom)})})
		d, err := tfprotov6.NewDynamicValue(root, value)
		if err != nil {
			t.Fatal(err)
		}
		return &d
	}
	marker := map[string]tftypes.Value{}
	config := dynamic(nil, marker)
	resp, err := server.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{TypeName: "marker_probe", PriorState: dynamic(marker, nil), ProposedNewState: dynamic(marker, marker), Config: config})
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range resp.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatal(d)
		}
	}
	planned, err := resp.PlannedState.Unmarshal(root)
	if err != nil {
		t.Fatal(err)
	}
	value, _, err := tftypes.WalkAttributePath(planned, tftypes.NewAttributePath().WithAttributeName("settings").WithAttributeName("default"))
	if err != nil {
		t.Fatal(err)
	}
	if !value.(tftypes.Value).IsNull() {
		t.Fatal("full protocol plan retained omitted default alongside explicit custom")
	}
	validation, err := server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{TypeName: "marker_probe", Config: dynamic(marker, marker)})
	if err != nil {
		t.Fatal(err)
	}
	failed := false
	for _, d := range validation.Diagnostics {
		failed = failed || d.Severity == tfprotov6.DiagnosticSeverityError
	}
	if !failed {
		t.Fatal("explicit contradictory configuration was silently accepted")
	}
}
