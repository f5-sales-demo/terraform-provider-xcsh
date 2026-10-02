// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/netip"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

type PublicIPBindingResource struct{ client *client.Client }
type publicIPBindingModel struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Namespace            types.String `tfsdk:"namespace"`
	ExpectedIP           types.String `tfsdk:"expected_ip"`
	VirtualSite          types.String `tfsdk:"virtual_site"`
	VirtualSiteNamespace types.String `tfsdk:"virtual_site_namespace"`
	OriginalBindings     types.String `tfsdk:"original_bindings"`
	ManagedBindings      types.String `tfsdk:"managed_bindings"`
}
type publicIPBindingObject struct {
	Document            map[string]interface{}
	Name, Namespace, IP string
	Bindings            []interface{}
}

var _ resource.Resource = &PublicIPBindingResource{}
var _ resource.ResourceWithConfigure = &PublicIPBindingResource{}

func NewPublicIPBindingResource() resource.Resource { return &PublicIPBindingResource{} }
func (r *PublicIPBindingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_public_ip_binding"
}
func (r *PublicIPBindingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	immutable := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage the regional virtual-site binding of an already allocated public IP. Creation adopts only the binding; deletion restores its captured original binding. This resource never allocates, deallocates or deletes the public IP.",
		Attributes: map[string]schema.Attribute{
			"id":                     schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":                   schema.StringAttribute{Required: true, PlanModifiers: immutable, MarkdownDescription: "Existing public IP object name."},
			"namespace":              schema.StringAttribute{Required: true, PlanModifiers: immutable, MarkdownDescription: "Existing public IP namespace, usually shared."},
			"expected_ip":            schema.StringAttribute{Required: true, PlanModifiers: immutable, MarkdownDescription: "Exact IPv4/IPv6 address expected in the allocated object."},
			"virtual_site":           schema.StringAttribute{Required: true, MarkdownDescription: "Desired REGIONAL_EDGE virtual site. Refreshed from the actual binding so drift is repairable."},
			"virtual_site_namespace": schema.StringAttribute{Required: true, MarkdownDescription: "Namespace of the regional virtual site."},
			"managed_bindings":       schema.StringAttribute{Computed: true, MarkdownDescription: "Last applied binding retained to reject foreign changes at teardown."},
			"original_bindings":      schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}, MarkdownDescription: "Original virtual-site bindings restored on delete; kept in Terraform state."},
		},
	}
}
func (r *PublicIPBindingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Invalid client", "Expected XC API client")
		return
	}
	r.client = c
}
func publicIPBindingPath(namespace, name string) string {
	return "/api/config/namespaces/" + url.PathEscape(namespace) + "/public_ips/" + url.PathEscape(name)
}
func readPublicIPBinding(ctx context.Context, c *client.Client, namespace, name, expected string) (*publicIPBindingObject, error) {
	if namespace == "" || name == "" {
		return nil, errors.New("public IP name and namespace are required")
	}
	if _, err := netip.ParseAddr(expected); err != nil {
		return nil, errors.New("expected_ip must identify a valid allocated IP")
	}
	var doc map[string]interface{}
	if err := c.Get(ctx, publicIPBindingPath(namespace, name), &doc); err != nil {
		return nil, err
	}
	metadata, ok := doc["metadata"].(map[string]interface{})
	if !ok || metadata["name"] != name || metadata["namespace"] != namespace {
		return nil, errors.New("public IP readback identity differs")
	}
	spec, ok := doc["spec"].(map[string]interface{})
	if !ok {
		return nil, errors.New("public IP spec is missing")
	}
	ip, _ := spec["ip"].(string)
	ipv6, _ := spec["ipv6"].(string)
	if ip != expected && ipv6 != expected {
		return nil, errors.New("allocated public IP address differs from expected_ip")
	}
	ip = expected
	bindings, ok := spec["virtual_sites"].([]interface{})
	if !ok {
		bindings = []interface{}{}
	}
	return &publicIPBindingObject{Document: doc, Name: name, Namespace: namespace, IP: ip, Bindings: bindings}, nil
}
func normalizedPublicIPBindings(values []interface{}) string {
	refs := make([][2]string, 0, len(values))
	for _, v := range values {
		m, ok := v.(map[string]interface{})
		if !ok {
			return "!malformed"
		}
		name, _ := m["name"].(string)
		namespace, _ := m["namespace"].(string)
		refs = append(refs, [2]string{namespace, name})
	}
	b, _ := json.Marshal(refs)
	return string(b)
}
func samePublicIPBindings(a, b []interface{}) bool {
	return normalizedPublicIPBindings(a) == normalizedPublicIPBindings(b)
}
func writePublicIPBinding(ctx context.Context, c *client.Client, current *publicIPBindingObject, desired []interface{}) error {
	if samePublicIPBindings(current.Bindings, desired) {
		return nil
	}
	// Use the latest full object identity and spec while changing only bindings.
	spec := current.Document["spec"].(map[string]interface{})
	revision, ok := current.Document["resource_version"].(string)
	if !ok || revision == "" {
		return errors.New("current public IP resource_version is missing")
	}
	next := make(map[string]interface{}, len(spec))
	for k, v := range spec {
		next[k] = v
	}
	next["virtual_sites"] = desired
	body := map[string]interface{}{"metadata": current.Document["metadata"], "spec": next, "resource_version": revision}
	writeErr := c.PutOnce(ctx, publicIPBindingPath(current.Namespace, current.Name), body, nil)
	observed, readErr := readPublicIPBinding(ctx, c, current.Namespace, current.Name, current.IP)
	if readErr != nil {
		return fmt.Errorf("public IP write outcome cannot be reconciled: %w", readErr)
	}
	if samePublicIPBindings(observed.Bindings, desired) {
		return nil
	}
	if writeErr != nil {
		return fmt.Errorf("public IP binding write failed and readback did not converge: %w", writeErr)
	}
	return errors.New("public IP binding readback did not converge")
}
func desiredPublicIPBinding(m publicIPBindingModel) []interface{} {
	return []interface{}{map[string]interface{}{"name": m.VirtualSite.ValueString(), "namespace": m.VirtualSiteNamespace.ValueString()}}
}
func (r *PublicIPBindingResource) validateSite(ctx context.Context, m publicIPBindingModel) error {
	if strings.TrimSpace(m.VirtualSite.ValueString()) == "" || strings.TrimSpace(m.VirtualSiteNamespace.ValueString()) == "" {
		return errors.New("regional virtual-site name and namespace are required")
	}
	site, err := r.client.GetVirtualSite(ctx, m.VirtualSiteNamespace.ValueString(), m.VirtualSite.ValueString())
	if err != nil {
		return err
	}
	if site.Spec["site_type"] != "REGIONAL_EDGE" {
		return errors.New("public IP binding requires a REGIONAL_EDGE virtual site")
	}
	return nil
}
func (r *PublicIPBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var m publicIPBindingModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := r.validateSite(ctx, m); err != nil {
		resp.Diagnostics.AddError("Regional virtual-site validation failed", err.Error())
		return
	}
	current, err := readPublicIPBinding(ctx, r.client, m.Namespace.ValueString(), m.Name.ValueString(), m.ExpectedIP.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Public IP adoption failed", err.Error())
		return
	}
	original, err := json.Marshal(current.Bindings)
	if err != nil {
		resp.Diagnostics.AddError("Public IP baseline capture failed", err.Error())
		return
	}
	m.ID = types.StringValue(m.Namespace.ValueString() + "/" + m.Name.ValueString())
	m.OriginalBindings = types.StringValue(string(original))
	managed, _ := json.Marshal(desiredPublicIPBinding(m))
	m.ManagedBindings = types.StringValue(string(managed))
	// Persist the restore baseline before a write whose transport outcome can be uncertain.
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err = writePublicIPBinding(ctx, r.client, current, desiredPublicIPBinding(m)); err != nil {
		resp.Diagnostics.AddError("Public IP binding failed", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
func (r *PublicIPBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var m publicIPBindingModel
	resp.Diagnostics.Append(req.State.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}
	current, err := readPublicIPBinding(ctx, r.client, m.Namespace.ValueString(), m.Name.ValueString(), m.ExpectedIP.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Public IP binding read failed", err.Error())
		return
	}
	if len(current.Bindings) == 1 {
		ref, ok := current.Bindings[0].(map[string]interface{})
		if !ok {
			resp.Diagnostics.AddError("Public IP binding read failed", "Malformed virtual-site reference")
			return
		}
		name, _ := ref["name"].(string)
		namespace, _ := ref["namespace"].(string)
		m.VirtualSite = types.StringValue(name)
		m.VirtualSiteNamespace = types.StringValue(namespace)
	} else {
		m.VirtualSite = types.StringValue("")
		m.VirtualSiteNamespace = types.StringValue("")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
func (r *PublicIPBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var m, prior publicIPBindingModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &m)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}
	m.OriginalBindings = prior.OriginalBindings
	m.ManagedBindings = prior.ManagedBindings
	if err := r.validateSite(ctx, m); err != nil {
		resp.Diagnostics.AddError("Regional virtual-site validation failed", err.Error())
		return
	}
	current, err := readPublicIPBinding(ctx, r.client, m.Namespace.ValueString(), m.Name.ValueString(), m.ExpectedIP.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Public IP binding read failed", err.Error())
		return
	}
	if err = writePublicIPBinding(ctx, r.client, current, desiredPublicIPBinding(m)); err != nil {
		resp.Diagnostics.AddError("Public IP binding failed", err.Error())
		return
	}
	managed, _ := json.Marshal(desiredPublicIPBinding(m))
	m.ManagedBindings = types.StringValue(string(managed))
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
func (r *PublicIPBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var m publicIPBindingModel
	resp.Diagnostics.Append(req.State.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var original []interface{}
	if err := json.Unmarshal([]byte(m.OriginalBindings.ValueString()), &original); err != nil {
		resp.Diagnostics.AddError("Public IP restore baseline invalid", err.Error())
		return
	}
	current, err := readPublicIPBinding(ctx, r.client, m.Namespace.ValueString(), m.Name.ValueString(), m.ExpectedIP.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Public IP binding restore read failed", err.Error())
		return
	}
	var managed []interface{}
	if err := json.Unmarshal([]byte(m.ManagedBindings.ValueString()), &managed); err != nil {
		resp.Diagnostics.AddError("Public IP ownership baseline invalid", err.Error())
		return
	}
	if !samePublicIPBindings(current.Bindings, managed) && !samePublicIPBindings(current.Bindings, original) {
		resp.Diagnostics.AddError("Public IP binding changed externally", "Refusing to overwrite an unrelated public IP binding during restoration")
		return
	}
	if err = writePublicIPBinding(ctx, r.client, current, original); err != nil {
		resp.Diagnostics.AddError("Public IP original binding restoration failed", err.Error())
	}
}
