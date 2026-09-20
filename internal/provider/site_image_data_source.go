// Code generated from the immutable SMSv2 KVM image contract. DO NOT EDIT.
// Source: F5 XC enriched API response-operation contract
package provider

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &SiteImageDataSource{}
var _ datasource.DataSourceWithConfigure = &SiteImageDataSource{}

func NewSiteImageDataSource() datasource.DataSource { return &SiteImageDataSource{} }

type SiteImageDataSource struct{ client *client.Client }
type SiteImageDataSourceModel struct {
	SiteName         types.String `tfsdk:"site_name"`
	ImageDownloadURL types.String `tfsdk:"image_download_url"`
	ImageName        types.String `tfsdk:"image_name"`
	ImageMD5Sum      types.String `tfsdk:"image_md5_sum"`
}

func (d *SiteImageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_image"
}
func (d *SiteImageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{MarkdownDescription: "Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution does not imply successful boot.", Attributes: map[string]schema.Attribute{
		"site_name":          schema.StringAttribute{Required: true, MarkdownDescription: "Existing KVM SMSv2 configuration name in system. Ownership is revalidated on each read."},
		"image_download_url": schema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "Validated HTTPS image URL. Protect Terraform state."},
		"image_name":         schema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "Image name returned by the Site-UID query. May contain a download URL; protect Terraform state."},
		"image_md5_sum":      schema.StringAttribute{Computed: true, MarkdownDescription: "Expected artifact MD5, which the consumer must verify before boot."},
	}}
}
func (d *SiteImageDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	configured, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected *client.Client")
		return
	}
	d.client = configured
}
func (d *SiteImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SiteImageDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.SiteName.IsNull() || data.SiteName.IsUnknown() || d.client == nil {
		resp.Diagnostics.AddError("KVM Image Resolution", "A known site name and configured client are required.")
		return
	}
	image, err := d.client.ResolveKVMImage(ctx, data.SiteName.ValueString(), client.KVMImageResolverContract{
		ConfigurationList: "/api/config/namespaces/system/securemesh_site_v2s",
		ConfigurationGet:  "/api/config/namespaces/system/securemesh_site_v2s/{site_name}",
		SiteList:          "/api/config/namespaces/system/sites",
		Query:             "/api/maurice/software_os_version",
	})
	if err != nil {
		resp.Diagnostics.AddError("KVM Image Resolution", err.Error())
		return
	}
	data.ImageDownloadURL = types.StringValue(image.DownloadURL)
	data.ImageName = types.StringValue(image.ImageName)
	data.ImageMD5Sum = types.StringValue(image.MD5)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
