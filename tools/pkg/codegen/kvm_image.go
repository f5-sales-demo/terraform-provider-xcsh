package codegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"text/template"
)

// Exact supported semantics bind this implementation to the source-owned join.
// A similar endpoint, object UID, or incomplete declaration must fail generation.
const kvmImageSemantics = `{
 "availability":"evidence_backed","enforcement":"required","strategy":"site_uid_os_image","namespace":"system",
 "terraform_data_source":"site_image","replaces_operation":"ves.io.schema.registration.CustomAPI.GetImageDownloadUrl",
 "configuration_get":{"method":"GET","path":"/api/config/namespaces/system/securemesh_site_v2s/{site_name}","operation_id":"ves.io.schema.views.securemesh_site_v2.API.Get"},
 "configuration_list":{"method":"GET","path":"/api/config/namespaces/system/securemesh_site_v2s","operation_id":"ves.io.schema.views.securemesh_site_v2.API.List"},
 "site_list":{"method":"GET","path":"/api/config/namespaces/system/sites","operation_id":"ves.io.schema.site.API.List"},
 "identity":{"configuration_name":"items[].name","configuration_uid":"items[].uid","site_uid":"items[].uid","owner_kind":"items[].owner_view.kind","required_owner_kind":"securemesh_site_v2","owner_uid":"items[].owner_view.uid","request_uid":"site_uid","configuration_matches":1,"owner_matches":1},
 "query":{"method":"POST","path":"/api/maurice/software_os_version","operation_id":"ves.io.schema.virtual_appliance.SoftwareVersionOsImageCustomApi.GetImage","request_schema":"virtual_applianceGetImageRequest","response_schema":"virtual_applianceGetImageResponse","request_field":"uids","request_cardinality":1,"side_effects":"none"},
 "response":{"mapping":"images","mapping_key":"site_uid","download_url":"download_image_link","image_name":"copy_image_name","md5":"image_md5_sum","error":"error_description"},
 "validation":{"caller_supplied_uid":false,"static_fallback":false,"require_current_owner_mapping":true,"required_platform_field":"spec.kvm","reject_nonempty_error":true,"require_complete_image_fields":true,"download_scheme":"https","download_hosts":["downloads.volterra.io"],"md5_pattern":"^[0-9a-fA-F]{32}$","verify_artifact_checksum":true,"boot_acceptance_required":true}
}`

func validateKVMImageResolution(image map[string]any) error {
	var expected map[string]any
	if err := json.Unmarshal([]byte(kvmImageSemantics), &expected); err != nil {
		return err
	}
	if len(image) != len(expected)+1 {
		return fmt.Errorf("KVM image resolution contract is incomplete")
	}
	for key, want := range expected {
		if !reflect.DeepEqual(image[key], want) {
			return fmt.Errorf("KVM image resolution %s semantics are unsupported", key)
		}
	}
	provenance, ok := image["provenance"].(map[string]any)
	if !ok || len(provenance) != 5 || provenance["source_issue"] != "f5-sales-demo/api-specs-enriched#1805" || provenance["receipt_path"] != "config/evidence/kvm-site-image-resolution-20260919.json" {
		return fmt.Errorf("KVM image resolution provenance is incomplete")
	}
	for key, length := range map[string]int{"source_commit": 40, "upstream_spec_sha256": 64, "receipt_sha256": 64} {
		value, ok := provenance[key].(string)
		if !ok || !regexp.MustCompile(fmt.Sprintf("^[a-f0-9]{%d}$", length)).MatchString(value) {
			return fmt.Errorf("KVM image resolution %s is malformed", key)
		}
	}
	return nil
}

// SMSv2ImageSuppressedOperations selects the API surfaces that the
// owner-verified resolver consumes internally. Neither the retired endpoint nor
// the raw Site-UID query may become a caller-controlled Terraform data source.
func SMSv2ImageSuppressedOperations(contractJSON []byte) (map[string]struct{}, error) {
	var contract smsv2ReleaseContract
	if err := json.Unmarshal(contractJSON, &contract); err != nil {
		return nil, err
	}
	image := contract.Providers.KVM.ImageResolution
	if err := validateKVMImageResolution(image); err != nil {
		return nil, err
	}
	query := image["query"].(map[string]any)
	return map[string]struct{}{
		image["replaces_operation"].(string): {},
		query["operation_id"].(string):       {},
	}, nil
}

func generateKVMImageDataSource(image map[string]any, outputDir string) error {
	if err := validateKVMImageResolution(image); err != nil {
		return err
	}
	path := filepath.Join(outputDir, "site_image_data_source.go")
	if err := ensureResponseOperationTarget(path); err != nil {
		return err
	}
	tmpl, err := template.New("KVM image").Parse(kvmImageDataSourceTemplate)
	if err != nil {
		return err
	}
	var source bytes.Buffer
	if err := tmpl.Execute(&source, image); err != nil {
		return err
	}
	formatted, err := format.Source(source.Bytes())
	if err != nil {
		return fmt.Errorf("format KVM image data source: %w", err)
	}
	return os.WriteFile(path, formatted, 0o644)
}

const kvmImageDataSourceTemplate = `// Code generated from the immutable SMSv2 KVM image contract. DO NOT EDIT.
// Source: F5 XC enriched API response-operation contract
package provider

import (
 "context"
 "github.com/hashicorp/terraform-plugin-framework/datasource"
 "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
 "github.com/hashicorp/terraform-plugin-framework/types"
 "github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var _ datasource.DataSource = &SiteImageDataSource{}
var _ datasource.DataSourceWithConfigure = &SiteImageDataSource{}
func NewSiteImageDataSource() datasource.DataSource { return &SiteImageDataSource{} }
type SiteImageDataSource struct { client *client.Client }
type SiteImageDataSourceModel struct {
 SiteName types.String ` + "`tfsdk:\"site_name\"`" + `
 ImageDownloadURL types.String ` + "`tfsdk:\"image_download_url\"`" + `
 ImageName types.String ` + "`tfsdk:\"image_name\"`" + `
 ImageMD5Sum types.String ` + "`tfsdk:\"image_md5_sum\"`" + `
}
func (d *SiteImageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) { resp.TypeName = req.ProviderTypeName + "_site_image" }
func (d *SiteImageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
 resp.Schema = schema.Schema{MarkdownDescription: "Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution does not imply successful boot.", Attributes: map[string]schema.Attribute{
  "site_name": schema.StringAttribute{Required: true, MarkdownDescription: "Existing KVM SMSv2 configuration name in system. Ownership is revalidated on each read."},
  "image_download_url": schema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "Validated HTTPS image URL. Protect Terraform state."},
  "image_name": schema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "Image name returned by the Site-UID query. May contain a download URL; protect Terraform state."},
  "image_md5_sum": schema.StringAttribute{Computed: true, MarkdownDescription: "Expected artifact MD5, which the consumer must verify before boot."},
 }}
}
func (d *SiteImageDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
 if req.ProviderData == nil { return }
 configured, ok := req.ProviderData.(*client.Client)
 if !ok { resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected *client.Client"); return }
 d.client = configured
}
func (d *SiteImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
 var data SiteImageDataSourceModel
 resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
 if resp.Diagnostics.HasError() { return }
 if data.SiteName.IsNull() || data.SiteName.IsUnknown() || d.client == nil {
  resp.Diagnostics.AddError("KVM Image Resolution", "A known site name and configured client are required."); return
 }
 image, err := d.client.ResolveKVMImage(ctx, data.SiteName.ValueString(), client.KVMImageResolverContract{
  ConfigurationList: {{printf "%q" .configuration_list.path}},
  ConfigurationGet: {{printf "%q" .configuration_get.path}},
  SiteList: {{printf "%q" .site_list.path}},
  Query: {{printf "%q" .query.path}},
 })
 if err != nil { resp.Diagnostics.AddError("KVM Image Resolution", err.Error()); return }
 data.ImageDownloadURL = types.StringValue(image.DownloadURL)
 data.ImageName = types.StringValue(image.ImageName)
 data.ImageMD5Sum = types.StringValue(image.MD5)
 resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
`
