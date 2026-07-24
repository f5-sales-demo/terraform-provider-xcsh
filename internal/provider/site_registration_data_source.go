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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var (
	_ datasource.DataSource              = &SiteRegistrationDataSource{}
	_ datasource.DataSourceWithConfigure = &SiteRegistrationDataSource{}
)

// defaultRegistrationNamespace is the namespace registrations live in.
const defaultRegistrationNamespace = "system"

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
	Token        types.String `tfsdk:"token"`
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
  count     = data.xcsh_site_registration.ce.found ? 1 : 0
  name      = data.xcsh_site_registration.ce.name
  namespace = data.xcsh_site_registration.ce.namespace
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
			"token": schema.StringAttribute{
				MarkdownDescription: "Site token the CE registered with.",
				Computed:            true,
				Sensitive:           true,
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
// No candidate returns (nil, nil): a site whose CE has not registered yet is
// the normal case, not an error. More than one candidate returns an error
// naming the hostnames to choose between.
func selectRegistration(items []client.RegistrationListItem, siteName, hostname string) (*client.RegistrationListItem, error) {
	var candidates []client.RegistrationListItem
	for _, it := range items {
		if cn := it.GetSpec.Passport.ClusterName; cn != "" && cn != siteName {
			continue
		}
		if hostname != "" && it.GetSpec.Infra.Hostname != hostname {
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
		data.Token = types.StringNull()
		if data.Hostname.IsNull() {
			data.Hostname = types.StringNull()
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	uid := match.UID
	if uid == "" {
		uid = match.SystemMetadata.UID
	}

	data.ID = types.StringValue(match.Name)
	data.Found = types.BoolValue(true)
	data.Name = types.StringValue(match.Name)
	data.UID = types.StringValue(uid)
	data.State = types.StringValue(match.Object.Status.CurrentState)
	data.ClusterName = types.StringValue(match.GetSpec.Passport.ClusterName)
	data.ClusterSize = types.Int64Value(match.GetSpec.Passport.ClusterSize)
	data.ProviderType = types.StringValue(match.GetSpec.Infra.Provider)
	data.Token = types.StringValue(match.GetSpec.Token)
	if data.Hostname.IsNull() {
		data.Hostname = types.StringValue(match.GetSpec.Infra.Hostname)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
