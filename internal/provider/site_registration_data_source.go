// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Site Registration Data Source for F5 XC
// Resolves the runtime registration of a site's Customer Edge node. A
// registration is named "r-<uuid>", never after the site, so this lookup is
// the only way to obtain the name that xcsh_registration_approval needs.

package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var (
	_ datasource.DataSource              = &SiteRegistrationDataSource{}
	_ datasource.DataSourceWithConfigure = &SiteRegistrationDataSource{}
)

// defaultRegistrationNamespace is the namespace registrations live in.
const defaultRegistrationNamespace = "system"

// terminalRegistrationStates are the registrationObjectState values a
// registration never leaves. F5 XC keeps such registrations indefinitely, and
// registrations_by_site keeps returning them, so rebuilding a CE under the same
// site name yields both the dead registration and the live one — with the same
// cluster_name and the same hostname. They are the registrations that can never
// be approved, so they are not candidates.
//
// The remaining registrationObjectState values — NOTSET, NEW, APPROVED,
// ADMITTED, PENDING, ONLINE, UPGRADING, MAINTENANCE — are live or still moving
// and stay candidates.
var terminalRegistrationStates = map[string]struct{}{
	"RETIRED":         {},
	"FAILED":          {},
	"FAILED_INACTIVE": {},
	"DONE":            {},
}

// isTerminalRegistrationState reports whether state is one a registration never
// leaves. An empty state is NOT terminal: a missing field must never turn a
// real match into a miss.
func isTerminalRegistrationState(state string) bool {
	_, ok := terminalRegistrationStates[state]
	return ok
}

func NewSiteRegistrationDataSource() datasource.DataSource {
	return &SiteRegistrationDataSource{}
}

type SiteRegistrationDataSource struct {
	client *client.Client
}

type SiteRegistrationDataSourceModel struct {
	ID           types.String `tfsdk:"id"`
	SiteName     types.String `tfsdk:"site_name"`
	Namespace    types.String `tfsdk:"namespace"`
	Hostname     types.String `tfsdk:"hostname"`
	Found        types.Bool   `tfsdk:"found"`
	Name         types.String `tfsdk:"name"`
	UID          types.String `tfsdk:"uid"`
	State        types.String `tfsdk:"state"`
	ClusterName  types.String `tfsdk:"cluster_name"`
	ClusterSize  types.Int64  `tfsdk:"cluster_size"`
	ProviderType types.String `tfsdk:"provider_type"`
	InstanceID   types.String `tfsdk:"instance_id"`
}

func (d *SiteRegistrationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_registration"
}

func (d *SiteRegistrationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Resolves the runtime registration of a site's Customer Edge (CE) node in F5 Distributed Cloud.

A registration is named ` + "`r-<uuid>`" + `, **not** after the site it belongs to, so it cannot be
read by site name. This data source lists the registrations belonging to a site and returns the
one that matches, giving you the name that ` + "`xcsh_registration_approval`" + ` requires.

The registration only exists once the CE has booted and registered with its token. Until then
this data source reports ` + "`found = false`" + ` **without raising an error**, so an approval can
safely be gated on it:

` + "```terraform" + `
data "xcsh_site_registration" "ce" {
  site_name = xcsh_securemesh_site_v2.ce.name
}

resource "xcsh_registration_approval" "ce" {
  count        = data.xcsh_site_registration.ce.found ? 1 : 0
  name         = data.xcsh_site_registration.ce.name
  namespace    = data.xcsh_site_registration.ce.namespace
  cluster_size = 1
}
` + "```" + `

**Possible ` + "`state`" + ` values:** ` + "`NOTSET`, `NEW`, `APPROVED`, `ADMITTED`, `RETIRED`, `FAILED`, `DONE`, `PENDING`, `ONLINE`, `UPGRADING`, `MAINTENANCE`, `FAILED_INACTIVE`" + `.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifier of this lookup: the registration name when one is found, otherwise null.",
				Computed:            true,
			},
			"site_name": schema.StringAttribute{
				MarkdownDescription: "Name of the F5 XC site whose CE registration should be resolved. Matched against each registration's `get_spec.passport.cluster_name`.",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Namespace holding the registrations. Defaults to `system`, where site registrations live.",
				Optional:            true,
				Computed:            true,
				// Reject an explicit empty string: the lookup would fall back to
				// `system` but report "" back, and a consumer feeding that into
				// xcsh_registration_approval.namespace would send an empty
				// namespace. Omit the argument to get the default instead.
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Node hostname used to pick one registration when a multi-node site has several. Optional for a single-node site; when omitted, the resolved node's hostname is returned here. Hostnames are only unique within a site.",
				Optional:            true,
				Computed:            true,
			},
			"found": schema.BoolAttribute{
				MarkdownDescription: "Whether a registration was resolved. `false` (with no error) while the CE has not registered yet — gate an approval's `count` on this.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Registration name (`r-<uuid>`) to pass to `xcsh_registration_approval`. Null when `found` is `false`.",
				Computed:            true,
			},
			"uid": schema.StringAttribute{
				MarkdownDescription: "Unique identifier of the registration (the `<uuid>` part of the name).",
				Computed:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "Current registration state, e.g. `PENDING` (awaiting approval) or `ONLINE` (node admitted and healthy).",
				Computed:            true,
			},
			"cluster_name": schema.StringAttribute{
				MarkdownDescription: "Cluster name the CE registered with, as reported in its passport. Equals `site_name` for a correctly configured site.",
				Computed:            true,
			},
			"cluster_size": schema.Int64Attribute{
				MarkdownDescription: "Number of nodes the CE reported for its cluster (1 for a single-node site, 3 for a three-node site).",
				Computed:            true,
			},
			"provider_type": schema.StringAttribute{
				MarkdownDescription: "Infrastructure provider the CE reported, e.g. `AZURE`, `AWS`, `GCP`, `VMWARE`.",
				Computed:            true,
			},
			"instance_id": schema.StringAttribute{
				MarkdownDescription: "Infrastructure instance identifier reported by the CE registration (`get_spec.infra.instance_id`). This distinguishes rebuilt nodes that reuse the same site and hostname.",
				Computed:            true,
			},
		},
	}
}

func (d *SiteRegistrationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// selectRegistration picks the one registration of siteName from the items the
// registrations_by_site endpoint returned.
//
// The endpoint is already site-scoped, so an item is a candidate unless it
// explicitly reports a different cluster_name (an item that reports none is
// kept rather than dropped, which would be a false negative). When hostname is
// set it further narrows the candidates — hostnames are only unique within a
// site, so this is a within-site discriminator.
//
// Registrations in a terminal state are then dropped. Rebuilding a CE under the
// same site name leaves the old registration behind, and it reports the same
// cluster_name and the same hostname as the fresh one, so neither filter above
// can separate them: without this the rebuild would resolve to the ambiguity
// error below and fail the consumer's whole plan.
//
// No candidate returns (nil, nil): a site whose CE has not registered yet — or
// one whose only registrations are dead — is the normal case, not an error.
// More than one live candidate returns an error naming the hostnames to choose
// between.
func selectRegistration(items []client.RegistrationListItem, siteName, hostname string) (*client.RegistrationListItem, error) {
	var candidates []client.RegistrationListItem
	for _, it := range items {
		if cn := it.GetSpec.Passport.ClusterName; cn != "" && cn != siteName {
			continue
		}
		if hostname != "" && it.GetSpec.Infra.Hostname != hostname {
			continue
		}
		if isTerminalRegistrationState(it.Object.Status.CurrentState) {
			continue
		}
		candidates = append(candidates, it)
	}

	switch len(candidates) {
	case 0:
		return nil, nil
	case 1:
		return &candidates[0], nil
	}

	hostnames := make([]string, 0, len(candidates))
	for _, c := range candidates {
		h := c.GetSpec.Infra.Hostname
		if h == "" {
			h = fmt.Sprintf("(no hostname, registration %s)", c.Name)
		}
		hostnames = append(hostnames, h)
	}
	sort.Strings(hostnames)

	return nil, fmt.Errorf(
		"site %q has %d matching registrations (node hostnames: %s); set the hostname argument to select one",
		siteName, len(candidates), strings.Join(hostnames, ", "),
	)
}

func (d *SiteRegistrationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SiteRegistrationDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	namespace := data.Namespace.ValueString()
	if data.Namespace.IsNull() || namespace == "" {
		namespace = defaultRegistrationNamespace
	}
	// Echo the configured value back verbatim so the read stays consistent
	// with the configuration.
	if data.Namespace.IsNull() {
		data.Namespace = types.StringValue(namespace)
	}

	siteName := data.SiteName.ValueString()
	configuredHostname := data.Hostname.ValueString()
	if data.Hostname.IsNull() {
		configuredHostname = ""
	}

	list, err := d.client.ListRegistrationsBySite(ctx, namespace, siteName)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to list registrations for site %q in namespace %q: %s", siteName, namespace, err),
		)
		return
	}

	if len(list.Errors) > 0 {
		msgs := make([]string, 0, len(list.Errors))
		for _, e := range list.Errors {
			msgs = append(msgs, strings.TrimSpace(e.Code+" "+e.Message))
		}
		resp.Diagnostics.AddError(
			"F5 XC API Error",
			fmt.Sprintf("Listing registrations for site %q in namespace %q returned errors: %s", siteName, namespace, strings.Join(msgs, "; ")),
		)
		return
	}

	match, err := selectRegistration(list.Items, siteName, configuredHostname)
	if err != nil {
		resp.Diagnostics.AddError("Ambiguous Site Registration", err.Error())
		return
	}

	if match == nil {
		// The CE has not registered yet (or the site does not exist). This is
		// an expected steady state before the node boots, so it is reported
		// without an error: consumers gate an approval's count on found.
		data.ID = types.StringNull()
		data.Found = types.BoolValue(false)
		data.Name = types.StringNull()
		data.UID = types.StringNull()
		data.State = types.StringNull()
		data.ClusterName = types.StringNull()
		data.ClusterSize = types.Int64Null()
		data.ProviderType = types.StringNull()
		data.InstanceID = types.StringNull()
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	applyRegistrationMatch(&data, match)
	if data.Hostname.IsNull() {
		data.Hostname = stringOrNull(match.GetSpec.Infra.Hostname)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func applyRegistrationMatch(data *SiteRegistrationDataSourceModel, match *client.RegistrationListItem) {
	uid := match.UID
	if uid == "" {
		uid = match.SystemMetadata.UID
	}
	data.ID = types.StringValue(match.Name)
	data.Found = types.BoolValue(true)
	data.Name = types.StringValue(match.Name)
	data.UID = stringOrNull(uid)
	// Fields the API may omit are reported null rather than "" / 0, so a consumer
	// can tell "not reported" from a real empty value (and to stay consistent with
	// the not-found branch above).
	data.State = stringOrNull(match.Object.Status.CurrentState)
	data.ClusterName = stringOrNull(match.GetSpec.Passport.ClusterName)
	data.ClusterSize = int64OrNull(match.GetSpec.Passport.ClusterSize)
	data.ProviderType = stringOrNull(match.GetSpec.Infra.Provider)
	data.InstanceID = stringOrNull(match.GetSpec.Infra.InstanceID)
}

// stringOrNull maps an omitted (empty) API string to a null attribute value so
// callers can distinguish "not reported" from a genuine empty string.
func stringOrNull(s string) types.String {
	if s == "" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

// int64OrNull maps an omitted (zero) API integer to a null attribute value.
// cluster_size is never legitimately 0 for a registered node.
func int64OrNull(i int64) types.Int64 {
	if i == 0 {
		return types.Int64Null()
	}
	return types.Int64Value(i)
}
