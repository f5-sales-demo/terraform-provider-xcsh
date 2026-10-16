package provider

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                   = &SecretPolicyResource{}
	_ resource.ResourceWithConfigure      = &SecretPolicyResource{}
	_ resource.ResourceWithImportState    = &SecretPolicyResource{}
	_ resource.ResourceWithValidateConfig = &SecretPolicyResource{}
)

type SecretPolicyResource struct{ client *client.Client }
type secretPolicyModel struct {
	Name            types.String `tfsdk:"name"`
	Namespace       types.String `tfsdk:"namespace"`
	Description     types.String `tfsdk:"description"`
	Labels          types.Map    `tfsdk:"labels"`
	Annotations     types.Map    `tfsdk:"annotations"`
	AllowF5XC       types.Bool   `tfsdk:"allow_f5xc"`
	Cache           types.String `tfsdk:"decrypt_cache_timeout"`
	Algorithm       types.String `tfsdk:"rule_combining_algorithm"`
	Rules           types.List   `tfsdk:"rules"`
	ID              types.String `tfsdk:"id"`
	UID             types.String `tfsdk:"uid"`
	MarkedForDelete types.Bool   `tfsdk:"marked_for_delete"`
	DeletionTime    types.String `tfsdk:"deletion_time"`
}
type secretPolicyRuleModel struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Action      types.String `tfsdk:"action"`
	Selector    types.List   `tfsdk:"client_selector"`
}

var secretPolicyIdentifier = regexp.MustCompile(`^[a-z0-9](?:[-a-z0-9]{0,61}[a-z0-9])?$`)
var secretPolicyDuration = regexp.MustCompile(`^[1-9][0-9]*[smh]$`)
var secretPolicyRuleTypes = map[string]attr.Type{"name": types.StringType, "description": types.StringType, "action": types.StringType, "client_selector": types.ListType{ElemType: types.StringType}}

func NewSecretPolicyResource() resource.Resource { return &SecretPolicyResource{} }
func (r *SecretPolicyResource) Metadata(_ context.Context, q resource.MetadataRequest, s *resource.MetadataResponse) {
	s.TypeName = q.ProviderTypeName + "_secret_policy"
}
func (r *SecretPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, s *resource.SchemaResponse) {
	replace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	s.Schema = schema.Schema{MarkdownDescription: "Owns a native Secret Policy with ordered inline ALLOW/DENY rules. Evaluation is FIRST_MATCH (top to bottom); no invented API algorithm field is sent. Deletion is soft: destroy remains pending while a tombstone exists and never repeats DELETE for a marked policy. Import namespace/name; legacy child references are normalized into inline rules. No secret or credential belongs in this resource.", Attributes: map[string]schema.Attribute{
		"name":                     schema.StringAttribute{Required: true, PlanModifiers: replace, Validators: []validator.String{validators.NameValidator()}},
		"namespace":                schema.StringAttribute{Required: true, PlanModifiers: replace, Validators: []validator.String{validators.NamespaceValidator()}},
		"description":              schema.StringAttribute{Optional: true},
		"labels":                   schema.MapAttribute{Optional: true, ElementType: types.StringType},
		"annotations":              schema.MapAttribute{Optional: true, ElementType: types.StringType},
		"allow_f5xc":               schema.BoolAttribute{Optional: true, Computed: true, Default: booldefault.StaticBool(false)},
		"decrypt_cache_timeout":    schema.StringAttribute{Optional: true, Computed: true, Default: stringdefault.StaticString("30s"), Validators: []validator.String{stringvalidator.RegexMatches(secretPolicyDuration, "Use a positive duration ending in s, m, or h.")}},
		"rule_combining_algorithm": schema.StringAttribute{Optional: true, Computed: true, Default: stringdefault.StaticString("FIRST_MATCH"), Validators: []validator.String{stringvalidator.OneOf("FIRST_MATCH")}},
		"rules": schema.ListNestedAttribute{Required: true, Validators: []validator.List{listvalidator.SizeBetween(1, 256)}, NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
			"name":            schema.StringAttribute{Required: true, Validators: []validator.String{validators.NameValidator()}},
			"description":     schema.StringAttribute{Optional: true},
			"action":          schema.StringAttribute{Required: true, Validators: []validator.String{stringvalidator.OneOf("ALLOW", "DENY")}},
			"client_selector": schema.ListAttribute{Optional: true, ElementType: types.StringType, Validators: []validator.List{listvalidator.SizeAtMost(1)}, MarkdownDescription: "Kubernetes selector expression. Omitted selector matches every client; use only for a final DENY rule."},
		}}},
		"id":                schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		"uid":               schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		"marked_for_delete": schema.BoolAttribute{Computed: true},
		"deletion_time":     schema.StringAttribute{Computed: true},
	}}
}
func (r *SecretPolicyResource) Configure(_ context.Context, q resource.ConfigureRequest, s *resource.ConfigureResponse) {
	if q.ProviderData == nil {
		return
	}
	var ok bool
	r.client, ok = q.ProviderData.(*client.Client)
	if !ok {
		s.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *client.Client.")
	}
}
func (r *SecretPolicyResource) ValidateConfig(ctx context.Context, q resource.ValidateConfigRequest, s *resource.ValidateConfigResponse) {
	var d secretPolicyModel
	s.Diagnostics.Append(q.Config.Get(ctx, &d)...)
	if s.Diagnostics.HasError() || d.Rules.IsUnknown() {
		return
	}
	var rules []secretPolicyRuleModel
	s.Diagnostics.Append(d.Rules.ElementsAs(ctx, &rules, false)...)
	names := map[string]bool{}
	for _, rule := range rules {
		if !rule.Name.IsUnknown() {
			name := rule.Name.ValueString()
			if names[name] {
				s.Diagnostics.AddError("Duplicate Secret Policy Rule", "Inline rule names must be unique.")
			}
			names[name] = true
		}
		if rule.Action.ValueString() == "ALLOW" && !rule.Selector.IsUnknown() {
			if rule.Selector.IsNull() || len(rule.Selector.Elements()) == 0 {
				s.Diagnostics.AddError("Unbounded ALLOW Rule", "An ALLOW rule requires an application-specific selector.")
			}
			for _, v := range rule.Selector.Elements() {
				if v.IsUnknown() {
					continue
				}
				if strings.TrimSpace(v.(types.String).ValueString()) == "" {
					s.Diagnostics.AddError("Empty ALLOW Selector", "An empty selector allows all clients.")
				}
			}
		}
	}
}
func policyMissing(err error) bool {
	var e *xcsherrors.XCSHError
	return errors.As(err, &e) && e.IsNotFound()
}
func (r *SecretPolicyResource) payload(ctx context.Context, d secretPolicyModel, previous *secretPolicyModel, observed *client.SecretPolicy) (*client.SecretPolicy, error) {
	p := &client.SecretPolicy{Metadata: client.SecretPolicyMetadata{Name: d.Name.ValueString(), Namespace: d.Namespace.ValueString(), Description: d.Description.ValueString()}, Spec: client.SecretPolicySpec{AllowF5XC: d.AllowF5XC.ValueBool(), DecryptCacheTimeout: d.Cache.ValueString(), RuleList: &client.SecretPolicyRuleList{Rules: []client.SecretPolicyRule{}}}}
	for _, field := range []struct {
		plan     types.Map
		old      types.Map
		observed map[string]string
		out      *map[string]string
	}{
		{d.Labels, types.MapNull(types.StringType), nil, &p.Metadata.Labels},
		{d.Annotations, types.MapNull(types.StringType), nil, &p.Metadata.Annotations},
	} {
		if previous != nil && observed != nil {
			if field.out == &p.Metadata.Labels {
				field.old = previous.Labels
				field.observed = observed.Metadata.Labels
			} else {
				field.old = previous.Annotations
				field.observed = observed.Metadata.Annotations
			}
		}
		m := map[string]string{}
		for k, v := range field.observed {
			m[k] = v
		}
		for k := range field.old.Elements() {
			delete(m, k)
		}
		if !field.plan.IsNull() {
			var owned map[string]string
			if diag := field.plan.ElementsAs(ctx, &owned, false); diag.HasError() {
				return nil, errors.New("invalid Secret Policy metadata")
			}
			for k, v := range owned {
				m[k] = v
			}
		}
		*field.out = m
	}
	var rules []secretPolicyRuleModel
	if diag := d.Rules.ElementsAs(ctx, &rules, false); diag.HasError() {
		return nil, errors.New("invalid Secret Policy rules")
	}
	for _, rule := range rules {
		v := client.SecretPolicyRule{Metadata: client.SecretPolicyMetadata{Name: rule.Name.ValueString(), Description: rule.Description.ValueString()}, Spec: client.SecretPolicyRuleSpec{Action: rule.Action.ValueString()}}
		if !rule.Selector.IsNull() {
			var expressions []string
			if diag := rule.Selector.ElementsAs(ctx, &expressions, false); diag.HasError() {
				return nil, errors.New("invalid Secret Policy selector")
			}
			v.Spec.ClientSelector = &client.SecretPolicySelector{Expressions: expressions}
		}
		p.Spec.RuleList.Rules = append(p.Spec.RuleList.Rules, v)
	}
	if observed != nil {
		p.ResourceVersion = observed.ResourceVersion
	}
	return p, nil
}
func policyMap(ctx context.Context, observed map[string]string, owned types.Map) (types.Map, error) {
	if owned.IsNull() {
		return owned, nil
	}
	selected := map[string]string{}
	for k := range owned.Elements() {
		if v, ok := observed[k]; ok {
			selected[k] = v
		}
	}
	if owned.IsUnknown() {
		selected = observed
	}
	value, diag := types.MapValueFrom(ctx, types.StringType, selected)
	if diag.HasError() {
		return types.Map{}, errors.New("invalid Secret Policy metadata readback")
	}
	return value, nil
}
func policyModel(ctx context.Context, d *secretPolicyModel, p *client.SecretPolicy) error {
	if !d.UID.IsNull() && !d.UID.IsUnknown() && d.UID.ValueString() != "" && d.UID.ValueString() != p.SystemMetadata.UID {
		return errors.New("secret policy UID changed; review ownership")
	}
	if d.ID.IsNull() || d.ID.IsUnknown() {
		if d.Labels.IsNull() && len(p.Metadata.Labels) > 0 {
			d.Labels = types.MapUnknown(types.StringType)
		}
		if d.Annotations.IsNull() && len(p.Metadata.Annotations) > 0 {
			d.Annotations = types.MapUnknown(types.StringType)
		}
	}
	d.ID = types.StringValue(d.Namespace.ValueString() + "/" + d.Name.ValueString())
	d.UID = types.StringValue(p.SystemMetadata.UID)
	d.MarkedForDelete = types.BoolValue(p.Spec.MarkedForDelete)
	d.DeletionTime = types.StringValue(p.Spec.DeletionTime)
	d.AllowF5XC = types.BoolValue(p.Spec.AllowF5XC)
	d.Cache = types.StringValue(p.Spec.DecryptCacheTimeout)
	d.Algorithm = types.StringValue(p.Spec.Algorithm)
	if !d.Description.IsNull() || p.Metadata.Description != "" {
		d.Description = types.StringValue(p.Metadata.Description)
	}
	var err error
	d.Labels, err = policyMap(ctx, p.Metadata.Labels, d.Labels)
	if err != nil {
		return err
	}
	d.Annotations, err = policyMap(ctx, p.Metadata.Annotations, d.Annotations)
	if err != nil {
		return err
	}
	if p.Spec.RuleList == nil {
		if p.Spec.MarkedForDelete {
			return nil
		}
		return errors.New("secret policy rules missing")
	}
	var old []secretPolicyRuleModel
	if !d.Rules.IsUnknown() && !d.Rules.IsNull() {
		if diag := d.Rules.ElementsAs(ctx, &old, false); diag.HasError() {
			return errors.New("invalid prior rules")
		}
	}
	rules := []secretPolicyRuleModel{}
	for i, rule := range p.Spec.RuleList.Rules {
		v := secretPolicyRuleModel{Name: types.StringValue(rule.Metadata.Name), Description: types.StringNull(), Action: types.StringValue(rule.Spec.Action), Selector: types.ListNull(types.StringType)}
		if rule.Metadata.Description != "" || (i < len(old) && !old[i].Description.IsNull()) {
			v.Description = types.StringValue(rule.Metadata.Description)
		}
		if rule.Spec.ClientSelector != nil && (len(rule.Spec.ClientSelector.Expressions) > 0 || (i < len(old) && !old[i].Selector.IsNull())) {
			selector, diagnostics := types.ListValueFrom(ctx, types.StringType, rule.Spec.ClientSelector.Expressions)
			v.Selector = selector
			if diagnostics.HasError() {
				return errors.New("invalid selector readback")
			}
		}
		rules = append(rules, v)
	}
	value, diag := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: secretPolicyRuleTypes}, rules)
	if diag.HasError() {
		return errors.New("invalid rules readback")
	}
	d.Rules = value
	return nil
}
func (r *SecretPolicyResource) Create(ctx context.Context, q resource.CreateRequest, s *resource.CreateResponse) {
	var d secretPolicyModel
	s.Diagnostics.Append(q.Plan.Get(ctx, &d)...)
	if s.Diagnostics.HasError() {
		return
	}
	_, err := r.client.GetSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString())
	if !policyMissing(err) {
		s.Diagnostics.AddError("Secret Policy Adoption Required", "Name exists or absence could not be verified. Review ownership and import namespace/name.")
		return
	}
	p, err := r.payload(ctx, d, nil, nil)
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Configuration", err.Error())
		return
	}
	err = r.client.CreateSecretPolicy(ctx, p)
	// Record identity before readback so a failed read cannot orphan a successful create.
	if err == nil {
		d.ID = types.StringValue(d.Namespace.ValueString() + "/" + d.Name.ValueString())
		d.UID = types.StringValue("")
		d.MarkedForDelete = types.BoolValue(false)
		d.DeletionTime = types.StringValue("")
		s.Diagnostics.Append(s.State.Set(ctx, &d)...)
	}
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Create Failed", "Create was sent once. Inspect the exact name and import an owned object before retrying.")
		return
	}
	observed, err := r.client.GetSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString())
	if err == nil {
		err = policyModel(ctx, &d, observed)
	}
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Readback Failed", "The issued identity is recorded. Readback failed; refresh it before continuing.")
		return
	}
	s.Diagnostics.Append(s.State.Set(ctx, &d)...)
}
func (r *SecretPolicyResource) Read(ctx context.Context, q resource.ReadRequest, s *resource.ReadResponse) {
	var d secretPolicyModel
	s.Diagnostics.Append(q.State.Get(ctx, &d)...)
	if s.Diagnostics.HasError() {
		return
	}
	p, err := r.client.GetSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString())
	if policyMissing(err) {
		s.State.RemoveResource(ctx)
		return
	}
	if err == nil {
		err = policyModel(ctx, &d, p)
	}
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Read Failed", err.Error())
		return
	}
	s.Diagnostics.Append(s.State.Set(ctx, &d)...)
}
func (r *SecretPolicyResource) Update(ctx context.Context, q resource.UpdateRequest, s *resource.UpdateResponse) {
	var d, previous secretPolicyModel
	s.Diagnostics.Append(q.Plan.Get(ctx, &d)...)
	s.Diagnostics.Append(q.State.Get(ctx, &previous)...)
	if s.Diagnostics.HasError() {
		return
	}
	observed, err := r.client.GetSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString())
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Preflight Failed", "Owned policy could not be read.")
		return
	}
	if observed.Spec.MarkedForDelete || (previous.UID.ValueString() != "" && previous.UID.ValueString() != observed.SystemMetadata.UID) {
		s.Diagnostics.AddError("Secret Policy Recovery Required", "Policy is marked for deletion or UID changed. Preserve ciphertext and state; review ownership.")
		return
	}
	p, err := r.payload(ctx, d, &previous, observed)
	if err == nil {
		err = r.client.ReplaceSecretPolicy(ctx, p)
	}
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Update Failed", "Replace was not replayed. Refresh the exact owned policy before retrying.")
		return
	}
	observed, err = r.client.GetSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString())
	if err == nil {
		err = policyModel(ctx, &d, observed)
	}
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Readback Failed", "Replace succeeded but readback failed. Refresh before continuing.")
		return
	}
	s.Diagnostics.Append(s.State.Set(ctx, &d)...)
}
func (r *SecretPolicyResource) Delete(ctx context.Context, q resource.DeleteRequest, s *resource.DeleteResponse) {
	var d secretPolicyModel
	s.Diagnostics.Append(q.State.Get(ctx, &d)...)
	if s.Diagnostics.HasError() {
		return
	}
	p, err := r.client.GetSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString())
	if policyMissing(err) {
		return
	}
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Delete Preflight Failed", "Owned policy could not be read.")
		return
	}
	if d.UID.ValueString() != "" && d.UID.ValueString() != p.SystemMetadata.UID {
		s.Diagnostics.AddError("Secret Policy Ownership Changed", "UID changed; deletion refused.")
		return
	}
	if !p.Spec.MarkedForDelete {
		if err = r.client.DeleteSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString()); err != nil {
			s.Diagnostics.AddError("Secret Policy Delete Failed", "DELETE was sent once. Refresh the exact policy before retrying.")
			return
		}
	}
	p, err = r.client.GetSecretPolicy(ctx, d.Namespace.ValueString(), d.Name.ValueString())
	if policyMissing(err) {
		return
	}
	if err != nil {
		s.Diagnostics.AddError("Secret Policy Deletion Unverified", "No absence proof; keep state and refresh.")
		return
	}
	_ = policyModel(ctx, &d, p)
	s.Diagnostics.Append(s.State.Set(ctx, &d)...)
	s.Diagnostics.AddError("Secret Policy Tombstone Pending", fmt.Sprintf("Policy still exists (marked_for_delete=%t, deletion_time=%s). State is retained. Retry destroy after expiry; marked policies never receive a second DELETE.", p.Spec.MarkedForDelete, p.Spec.DeletionTime))
}
func (r *SecretPolicyResource) ImportState(ctx context.Context, q resource.ImportStateRequest, s *resource.ImportStateResponse) {
	// Import ID format: namespace/name
	parts := strings.Split(q.ID, "/")
	if len(parts) != 2 || !secretPolicyIdentifier.MatchString(parts[0]) || !secretPolicyIdentifier.MatchString(parts[1]) {
		s.Diagnostics.AddError("Invalid Import ID", "Use namespace/name.")
		return
	}
	s.Diagnostics.Append(s.State.SetAttribute(ctx, path.Root("namespace"), parts[0])...)
	s.Diagnostics.Append(s.State.SetAttribute(ctx, path.Root("name"), parts[1])...)
}
