// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Site Image Data Source for F5 XC.
// Resolves the current F5-provided Customer Edge image for a platform. The
// public image-download action is separate from Secure Mesh Site v2 objects.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var (
	_ datasource.DataSource              = &SiteImageDataSource{}
	_ datasource.DataSourceWithConfigure = &SiteImageDataSource{}
)

func NewSiteImageDataSource() datasource.DataSource {
	return &SiteImageDataSource{}
}

type SiteImageDataSource struct {
	client *client.Client
}

type SiteImageDataSourceModel struct {
	ID                  types.String `tfsdk:"id"`
	Platform            types.String `tfsdk:"platform"`
	ImageDownloadURL    types.String `tfsdk:"image_download_url"`
	ImageMD5DownloadURL types.String `tfsdk:"image_md5_download_url"`
}

func (d *SiteImageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_image"
}

func (d *SiteImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Returns the current F5 Distributed Cloud Customer Edge image download URLs for one platform.

This data source calls the F5 XC public image-download action. It is deliberately separate from
` + "`xcsh_securemesh_site_v2`" + ` because an image is a platform artifact, not a site configuration field.
The returned URLs can be short-lived or authorize artifact access, so Terraform treats both as sensitive.

` + "```terraform" + `
data "xcsh_site_image" "kvm" {
  platform = "Kvm"
}
` + "```" + ``,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Identifier of this lookup. Equal to `platform`.",
				Computed:            true,
			},
			"platform": schema.StringAttribute{
				MarkdownDescription: "Customer Edge platform requested from the F5 XC image service, for example `Kvm`.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"image_download_url": schema.StringAttribute{
				MarkdownDescription: "URL for the platform image. Sensitive because the URL can authorize access to the artifact.",
				Computed:            true,
				Sensitive:           true,
			},
			"image_md5_download_url": schema.StringAttribute{
				MarkdownDescription: "URL for the image MD5 checksum. Sensitive because the URL can authorize access to the artifact.",
				Computed:            true,
				Sensitive:           true,
			},
		},
	}
}

func (d *SiteImageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SiteImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SiteImageDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	platform := data.Platform.ValueString()
	image, err := d.client.GetSiteImage(ctx, platform)
	if err != nil {
		resp.Diagnostics.AddError(
			"F5 XC Site Image Lookup Failed",
			fmt.Sprintf("Unable to retrieve the %q Customer Edge image URLs: %s", platform, err),
		)
		return
	}
	if image.ImageDownloadURL == "" || image.ImageMD5DownloadURL == "" {
		resp.Diagnostics.AddError(
			"F5 XC Site Image Lookup Returned an Incomplete Response",
			fmt.Sprintf("The %q image lookup did not return both image_download_url and image_md5_download_url.", platform),
		)
		return
	}

	data.ID = types.StringValue(platform)
	data.ImageDownloadURL = types.StringValue(image.ImageDownloadURL)
	data.ImageMD5DownloadURL = types.StringValue(image.ImageMD5DownloadURL)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
