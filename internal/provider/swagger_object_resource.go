package provider

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                   = &SwaggerObjectResource{}
	_ resource.ResourceWithConfigure      = &SwaggerObjectResource{}
	_ resource.ResourceWithImportState    = &SwaggerObjectResource{}
	_ resource.ResourceWithValidateConfig = &SwaggerObjectResource{}
	_ resource.ResourceWithModifyPlan     = &SwaggerObjectResource{}
)

type SwaggerObjectResource struct{ client *client.Client }
type SwaggerObjectResourceModel struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
	Content   types.String `tfsdk:"content"`
	Version   types.String `tfsdk:"version"`
	ID        types.String `tfsdk:"id"`
	Path      types.String `tfsdk:"path"`
	SHA256    types.String `tfsdk:"sha256"`
}

func NewSwaggerObjectResource() resource.Resource { return &SwaggerObjectResource{} }
func (r *SwaggerObjectResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_swagger_object"
}
func (r *SwaggerObjectResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	replace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{MarkdownDescription: "Owns one immutable, content-verified Swagger object version. Use a content-addressed name: content changes require a new name and replace only the owned version. XC may reuse a deleted version label, so same-name content replacement is rejected. Import uses namespace/name/version; latest and external presigned URLs are prohibited. State contains the complete document.", Attributes: map[string]schema.Attribute{
		"name":      schema.StringAttribute{Required: true, MarkdownDescription: "Object DNS label. Include a content digest so changed content has a distinct immutable path.", PlanModifiers: replace},
		"namespace": schema.StringAttribute{Required: true, MarkdownDescription: "Owning namespace DNS label.", PlanModifiers: replace},
		"content":   schema.StringAttribute{Required: true, MarkdownDescription: "Exact UTF-8 OpenAPI 3.0/3.1 or Swagger 2.0 JSON bytes. Use file(). Changed bytes require a new object name and a reviewed replacement.", PlanModifiers: replace},
		"version":   schema.StringAttribute{Computed: true, MarkdownDescription: "Exact server-issued immutable version."},
		"id":        schema.StringAttribute{Computed: true, MarkdownDescription: "Import identity: namespace/name/version."},
		"path":      schema.StringAttribute{Computed: true, MarkdownDescription: "Immutable object-store path for API definition swagger_specs."},
		"sha256":    schema.StringAttribute{Computed: true, MarkdownDescription: "SHA-256 of the exact verified content bytes."},
	}}
}
func (r *SwaggerObjectResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	r.client, ok = req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *client.Client.")
	}
}
func (r *SwaggerObjectResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var d SwaggerObjectResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &d)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !d.Name.IsUnknown() && !d.Namespace.IsUnknown() {
		if _, err := client.SwaggerIdentity(d.Namespace.ValueString(), d.Name.ValueString(), ""); err != nil {
			resp.Diagnostics.AddError("Invalid Swagger Identity", err.Error())
		}
	}
	if !d.Content.IsUnknown() && !d.Content.IsNull() {
		if err := client.ValidateSwaggerContent(d.Content.ValueString()); err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("content"), "Invalid Swagger Content", err.Error())
		}
	}
}

// ModifyPlan prevents XC from reusing a deleted version label for different bytes.
func (r *SwaggerObjectResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	var before, after SwaggerObjectResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &before)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &after)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if before.Content.IsUnknown() || before.Content.IsNull() || after.Content.IsUnknown() || after.Content.IsNull() || before.Name.IsUnknown() || after.Name.IsUnknown() || before.Namespace.IsUnknown() || after.Namespace.IsUnknown() {
		return
	}
	if before.Content.ValueString() != after.Content.ValueString() && before.Name.Equal(after.Name) && before.Namespace.Equal(after.Namespace) {
		resp.Diagnostics.AddAttributeError(path.Root("name"), "New Swagger Name Required", "XC may reuse a deleted version label. Changed content must use a new content-addressed object name so its immutable path cannot identify different bytes. Review the replacement before applying it.")
	}
}

func swaggerModel(d *SwaggerObjectResourceModel, version, content string) {
	d.Version = types.StringValue(version)
	d.Content = types.StringValue(content)
	d.SHA256 = types.StringValue(fmt.Sprintf("%x", sha256.Sum256([]byte(content))))
	d.ID = types.StringValue(d.Namespace.ValueString() + "/" + d.Name.ValueString() + "/" + version)
	p, _ := client.SwaggerIdentity(d.Namespace.ValueString(), d.Name.ValueString(), version)
	d.Path = types.StringValue(p)
}
func swaggerMissing(err error) bool {
	var e *xcsherrors.XCSHError
	return errors.As(err, &e) && e.IsNotFound()
}
func (r *SwaggerObjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	var d SwaggerObjectResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &d)...)
	if resp.Diagnostics.HasError() {
		return
	}
	ns, n, c := d.Namespace.ValueString(), d.Name.ValueString(), d.Content.ValueString()
	versions, err := r.client.SwaggerVersions(ctx, ns, n)
	if err != nil {
		resp.Diagnostics.AddError("Swagger Conflict Check Failed", "Unable to verify exact-name version inventory.")
		return
	}
	if len(versions) > 0 {
		resp.Diagnostics.AddError("Swagger Adoption Required", "An out-of-state object exists. Review ownership and explicitly import namespace/name/version before deployment.")
		return
	}
	obj, err := r.client.CreateSwaggerObject(ctx, ns, n, c)
	if err != nil {
		resp.Diagnostics.AddError("Swagger Upload Failed", "Upload outcome may be ambiguous. No automatic replay occurred. Inspect exact-name versions and import a verified owned version before retrying.")
		return
	}
	// Record returned identity even if verification later fails, so an issued version is not orphaned.
	swaggerModel(&d, obj.Metadata.Version, c)
	resp.Diagnostics.Append(resp.State.Set(ctx, &d)...)
	exact, err := r.client.GetSwaggerObject(ctx, ns, n, obj.Metadata.Version)
	if err != nil {
		resp.Diagnostics.AddError("Swagger Readback Failed", "The issued version is recorded in state; exact-version readback failed.")
		return
	}
	actual, err := exact.Content(ns, n, obj.Metadata.Version)
	if err != nil || actual != c {
		resp.Diagnostics.AddError("Swagger Content Verification Failed", "Exact-version readback differs from configured bytes.")
		return
	}
}
func (r *SwaggerObjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	var d SwaggerObjectResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &d)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := r.client.GetSwaggerObject(ctx, d.Namespace.ValueString(), d.Name.ValueString(), d.Version.ValueString())
	if swaggerMissing(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Swagger Read Failed", "Exact-version identity or content could not be verified.")
		return
	}
	content, err := obj.Content(d.Namespace.ValueString(), d.Name.ValueString(), d.Version.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Swagger Content Failed", err.Error())
		return
	}
	if !d.SHA256.IsNull() && !d.SHA256.IsUnknown() && d.SHA256.ValueString() != fmt.Sprintf("%x", sha256.Sum256([]byte(content))) {
		resp.Diagnostics.AddError("Immutable Swagger Drift", "Owned version content changed outside Terraform. Review ownership before replacement.")
		return
	}
	swaggerModel(&d, d.Version.ValueString(), content)
	resp.Diagnostics.Append(resp.State.Set(ctx, &d)...)
}
func (r *SwaggerObjectResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Immutable Swagger Object", "All configured changes require replacement.")
}
func (r *SwaggerObjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	var d SwaggerObjectResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &d)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := r.client.GetSwaggerObject(ctx, d.Namespace.ValueString(), d.Name.ValueString(), d.Version.ValueString())
	if swaggerMissing(err) {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Swagger Delete Verification Failed", "Exact owned version could not be verified; deletion refused.")
		return
	}
	content, err := obj.Content(d.Namespace.ValueString(), d.Name.ValueString(), d.Version.ValueString())
	if err != nil || d.SHA256.IsNull() || d.SHA256.ValueString() != fmt.Sprintf("%x", sha256.Sum256([]byte(content))) {
		resp.Diagnostics.AddError("Swagger Delete Ownership Failed", "Owned content digest mismatch; deletion refused.")
		return
	}
	err = r.client.DeleteSwaggerObject(ctx, d.Namespace.ValueString(), d.Name.ValueString(), d.Version.ValueString())
	if err != nil && !swaggerMissing(err) {
		resp.Diagnostics.AddError("Swagger Delete Failed", "Deletion of the exact owned version failed.")
	}
}
func (r *SwaggerObjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import ID format: namespace/name/version
	parts := strings.Split(req.ID, "/")
	if len(parts) != 3 || parts[2] == "" {
		resp.Diagnostics.AddError("Invalid Swagger Import", "Use namespace/name/version with an exact version.")
		return
	}
	if _, err := client.SwaggerIdentity(parts[0], parts[1], parts[2]); err != nil {
		resp.Diagnostics.AddError("Invalid Swagger Import", err.Error())
		return
	}
	for i, key := range []string{"namespace", "name", "version"} {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(key), parts[i])...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
