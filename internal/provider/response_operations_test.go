// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

type responseOperationRequest struct {
	Method string
	Path   string
	Query  url.Values
	Body   map[string]interface{}
}

func responseOperationRaw(t *testing.T, model interface{}, schemaType attr.Type) tftypes.Value {
	t.Helper()
	ctx := context.Background()
	var value types.Object
	if diagnostics := tfsdk.ValueFrom(ctx, model, schemaType, &value); diagnostics.HasError() {
		t.Fatalf("encode operation config: %v", diagnostics)
	}
	raw, err := value.ToTerraformValue(ctx)
	if err != nil {
		t.Fatalf("convert operation config: %v", err)
	}
	return raw
}

func TestGeneratedResponseOperationsRouteAndDecode(t *testing.T) {
	ctx := context.Background()
	requests := make(chan responseOperationRequest, 8)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		var body map[string]interface{}
		if request.Body != nil {
			_ = json.NewDecoder(request.Body).Decode(&body)
		}
		requests <- responseOperationRequest{Method: request.Method, Path: request.URL.Path, Query: request.URL.Query(), Body: body}
		w.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/register/namespaces/system/get-image-download-url":
			_, _ = w.Write([]byte(`{"image_download_url":"https://download.example/image","image_md5_download_url":"https://download.example/md5"}`))
		case "/api/register/namespaces/tenant-a/registrations":
			_, _ = w.Write([]byte(`{"items":[{"name":"r-1","get_spec":{"infra":{"instance_id":"instance-1"}}}],"errors":[]}`))
		case "/api/register/namespaces/system/registrations_by_site/site-a",
			"/api/register/namespaces/system/listregistrationsbystate":
			_, _ = w.Write([]byte(`{"items":[],"errors":[]}`))
		case "/api/register/namespaces/system/get-cloud-init-config":
			_, _ = w.Write([]byte(`{"cloud_init_config":"sensitive-cloud-init"}`))
		default:
			w.WriteHeader(http.StatusAccepted)
			_, _ = w.Write([]byte(`{}`))
		}
	}))
	defer server.Close()
	apiClient := client.NewClient(server.URL, "test-token", client.WithMaxRetries(0))

	image := &SiteImageDataSource{client: apiClient}
	imageSchema := &datasource.SchemaResponse{}
	image.Schema(ctx, datasource.SchemaRequest{}, imageSchema)
	for _, name := range []string{"image_download_url", "image_md5_download_url"} {
		if !imageSchema.Schema.Attributes[name].IsSensitive() {
			t.Fatalf("site image attribute %s is not sensitive", name)
		}
	}
	imageConfig := SiteImageDataSourceModel{
		ProviderRef: types.StringValue("KVM"), ImageDownloadURL: types.StringNull(), ImageMD5DownloadURL: types.StringNull(),
	}
	imageResp := datasource.ReadResponse{State: tfsdk.State{Schema: imageSchema.Schema}}
	image.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: imageSchema.Schema, Raw: responseOperationRaw(t, imageConfig, imageSchema.Schema.Type())}}, &imageResp)
	if imageResp.Diagnostics.HasError() {
		t.Fatalf("site image read: %v", imageResp.Diagnostics)
	}
	var imageState SiteImageDataSourceModel
	imageResp.Diagnostics.Append(imageResp.State.Get(ctx, &imageState)...)
	if imageState.ImageDownloadURL.ValueString() != "https://download.example/image" || imageState.ImageMD5DownloadURL.ValueString() != "https://download.example/md5" {
		t.Fatalf("site image response was not decoded: %+v", imageState)
	}
	assertResponseOperationRequest(t, <-requests, http.MethodPost, "/api/register/namespaces/system/get-image-download-url", nil, map[string]interface{}{"provider": "KVM"})

	registrations := &SiteRegistrationsDataSource{client: apiClient}
	registrationsSchema := &datasource.SchemaResponse{}
	registrations.Schema(ctx, datasource.SchemaRequest{}, registrationsSchema)
	registrationsConfig := SiteRegistrationsDataSourceModel{
		Namespace: types.StringValue("tenant-a"), LabelFilter: types.StringValue("env=demo"),
		ReportFields:       types.ListValueMust(types.StringType, []attr.Value{types.StringValue("name")}),
		ReportStatusFields: types.ListValueMust(types.StringType, []attr.Value{types.StringValue("state")}),
		Errors:             types.ListNull(types.ObjectType{AttrTypes: SiteRegistrationsErrorsModelAttrTypes}),
		Items:              types.ListNull(types.ObjectType{AttrTypes: SiteRegistrationsItemsModelAttrTypes}),
	}
	registrationsResp := datasource.ReadResponse{State: tfsdk.State{Schema: registrationsSchema.Schema}}
	registrations.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: registrationsSchema.Schema, Raw: responseOperationRaw(t, registrationsConfig, registrationsSchema.Schema.Type())}}, &registrationsResp)
	if registrationsResp.Diagnostics.HasError() {
		t.Fatalf("registrations read: %v", registrationsResp.Diagnostics)
	}
	var registrationsState SiteRegistrationsDataSourceModel
	registrationsResp.Diagnostics.Append(registrationsResp.State.Get(ctx, &registrationsState)...)
	var registrationItems []SiteRegistrationsItemsModel
	registrationsResp.Diagnostics.Append(registrationsState.Items.ElementsAs(ctx, &registrationItems, false)...)
	if registrationsResp.Diagnostics.HasError() || len(registrationItems) != 1 || registrationItems[0].GetSpec == nil || registrationItems[0].GetSpec.Infra == nil || registrationItems[0].GetSpec.Infra.InstanceID.ValueString() != "instance-1" {
		t.Fatalf("typed registration response was not decoded: items=%+v diagnostics=%v", registrationItems, registrationsResp.Diagnostics)
	}
	assertResponseOperationRequest(t, <-requests, http.MethodGet, "/api/register/namespaces/tenant-a/registrations", url.Values{"label_filter": {"env=demo"}, "report_fields": {"name"}, "report_status_fields": {"state"}}, nil)

	bySite := &SiteRegistrationsBySiteDataSource{client: apiClient}
	bySiteSchema := &datasource.SchemaResponse{}
	bySite.Schema(ctx, datasource.SchemaRequest{}, bySiteSchema)
	bySiteConfig := SiteRegistrationsBySiteDataSourceModel{
		SiteName: types.StringValue("site-a"), Namespace: types.StringNull(),
		Errors: types.ListNull(types.ObjectType{AttrTypes: SiteRegistrationsBySiteErrorsModelAttrTypes}),
		Items:  types.ListNull(types.ObjectType{AttrTypes: SiteRegistrationsBySiteItemsModelAttrTypes}),
	}
	bySiteResp := datasource.ReadResponse{State: tfsdk.State{Schema: bySiteSchema.Schema}}
	bySite.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: bySiteSchema.Schema, Raw: responseOperationRaw(t, bySiteConfig, bySiteSchema.Schema.Type())}}, &bySiteResp)
	if bySiteResp.Diagnostics.HasError() {
		t.Fatalf("registrations-by-site read: %v", bySiteResp.Diagnostics)
	}
	assertResponseOperationRequest(t, <-requests, http.MethodGet, "/api/register/namespaces/system/registrations_by_site/site-a", nil, nil)

	byState := &SiteRegistrationsByStateDataSource{client: apiClient}
	byStateSchema := &datasource.SchemaResponse{}
	byState.Schema(ctx, datasource.SchemaRequest{}, byStateSchema)
	byStateConfig := SiteRegistrationsByStateDataSourceModel{
		State: types.StringValue("ONLINE"), Namespace: types.StringNull(),
		Errors: types.ListNull(types.ObjectType{AttrTypes: SiteRegistrationsByStateErrorsModelAttrTypes}),
		Items:  types.ListNull(types.ObjectType{AttrTypes: SiteRegistrationsByStateItemsModelAttrTypes}),
	}
	byStateResp := datasource.ReadResponse{State: tfsdk.State{Schema: byStateSchema.Schema}}
	byState.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: byStateSchema.Schema, Raw: responseOperationRaw(t, byStateConfig, byStateSchema.Schema.Type())}}, &byStateResp)
	if byStateResp.Diagnostics.HasError() {
		t.Fatalf("registrations-by-state read: %v", byStateResp.Diagnostics)
	}
	assertResponseOperationRequest(t, <-requests, http.MethodPost, "/api/register/namespaces/system/listregistrationsbystate", nil, map[string]interface{}{"namespace": "system", "state": "ONLINE"})

	issuance := &SiteCloudInitResource{client: apiClient}
	issuanceSchema := &resource.SchemaResponse{}
	issuance.Schema(ctx, resource.SchemaRequest{}, issuanceSchema)
	if !issuanceSchema.Schema.Attributes["cloud_init_config"].IsSensitive() {
		t.Fatal("cloud_init_config is not sensitive")
	}
	if _, supportsImport := interface{}(issuance).(resource.ResourceWithImportState); supportsImport {
		t.Fatal("create-once issuance must not support import")
	}
	issuancePlan := SiteCloudInitResourceModel{
		Provider: types.StringValue("KVM"), SiteName: types.StringValue("site-a"), EnableManagementNetwork: types.BoolNull(),
		CloudInitConfig: types.StringNull(), ID: types.StringNull(),
	}
	issuanceResp := resource.CreateResponse{State: tfsdk.State{Schema: issuanceSchema.Schema}}
	issuance.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan{Schema: issuanceSchema.Schema, Raw: responseOperationRaw(t, issuancePlan, issuanceSchema.Schema.Type())}}, &issuanceResp)
	if issuanceResp.Diagnostics.HasError() {
		t.Fatalf("cloud-init issuance: %v", issuanceResp.Diagnostics)
	}
	var issuanceState SiteCloudInitResourceModel
	issuanceResp.Diagnostics.Append(issuanceResp.State.Get(ctx, &issuanceState)...)
	if issuanceState.CloudInitConfig.ValueString() != "sensitive-cloud-init" || issuanceState.ID.IsNull() || issuanceState.ID.ValueString() == "" {
		t.Fatalf("issuance response/state not retained: %+v", issuanceState)
	}
	assertResponseOperationRequest(t, <-requests, http.MethodGet, "/api/register/namespaces/system/get-cloud-init-config", url.Values{"provider": {"KVM"}, "site_name": {"site-a"}}, nil)
	issuanceReadResp := resource.ReadResponse{State: tfsdk.State{Schema: issuanceSchema.Schema}}
	issuance.Read(ctx, resource.ReadRequest{State: issuanceResp.State}, &issuanceReadResp)
	if issuanceReadResp.Diagnostics.HasError() {
		t.Fatalf("cloud-init retained read: %v", issuanceReadResp.Diagnostics)
	}
	var retainedIssuance SiteCloudInitResourceModel
	issuanceReadResp.Diagnostics.Append(issuanceReadResp.State.Get(ctx, &retainedIssuance)...)
	if retainedIssuance.CloudInitConfig.ValueString() != "sensitive-cloud-init" || retainedIssuance.ID.ValueString() != issuanceState.ID.ValueString() {
		t.Fatalf("issuance read did not retain create-once state: %+v", retainedIssuance)
	}
	issuanceDeleteResp := &resource.DeleteResponse{}
	issuance.Delete(ctx, resource.DeleteRequest{State: issuanceResp.State}, issuanceDeleteResp)
	if issuanceDeleteResp.Diagnostics.HasError() || len(requests) != 0 {
		t.Fatalf("issuance delete contacted API or returned diagnostics: requests=%d diagnostics=%v", len(requests), issuanceDeleteResp.Diagnostics)
	}

	invokeUpgradeAction(t, ctx, &SiteUpgradeSwAction{client: apiClient}, SiteUpgradeSwActionModel{Name: types.StringValue("site-a"), Namespace: types.StringValue("system"), Version: types.StringValue("9.0.0"), Force: types.BoolNull()})
	assertResponseOperationRequest(t, <-requests, http.MethodPost, "/api/config/namespaces/system/sites/site-a/upgrade_sw", nil, map[string]interface{}{"name": "site-a", "namespace": "system", "version": "9.0.0", "force": false})
	invokeUpgradeAction(t, ctx, &SiteUpgradeOSAction{client: apiClient}, SiteUpgradeOSActionModel{Name: types.StringValue("site-a"), Namespace: types.StringValue("system"), Version: types.StringValue("10.0.0"), Force: types.BoolValue(true)})
	assertResponseOperationRequest(t, <-requests, http.MethodPost, "/api/config/namespaces/system/sites/site-a/upgrade_os", nil, map[string]interface{}{"name": "site-a", "namespace": "system", "version": "10.0.0", "force": true})
	if len(requests) != 0 {
		t.Fatalf("actions performed unexpected polling requests: %+v", <-requests)
	}
}

type invokableAction interface {
	Schema(context.Context, action.SchemaRequest, *action.SchemaResponse)
	Invoke(context.Context, action.InvokeRequest, *action.InvokeResponse)
}

func invokeUpgradeAction(t *testing.T, ctx context.Context, implementation invokableAction, model interface{}) {
	t.Helper()
	schemaResponse := &action.SchemaResponse{}
	implementation.Schema(ctx, action.SchemaRequest{}, schemaResponse)
	response := &action.InvokeResponse{}
	implementation.Invoke(ctx, action.InvokeRequest{Config: tfsdk.Config{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type())}}, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("invoke upgrade action: %v", response.Diagnostics)
	}
}

func assertResponseOperationRequest(t *testing.T, got responseOperationRequest, method, path string, query url.Values, body map[string]interface{}) {
	t.Helper()
	if got.Method != method || got.Path != path {
		t.Fatalf("request = %s %s, want %s %s", got.Method, got.Path, method, path)
	}
	if query != nil && got.Query.Encode() != query.Encode() {
		t.Fatalf("request query = %q, want %q", got.Query.Encode(), query.Encode())
	}
	if body != nil {
		encodedGot, _ := json.Marshal(got.Body)
		encodedWant, _ := json.Marshal(body)
		if string(encodedGot) != string(encodedWant) {
			t.Fatalf("request body = %s, want %s", encodedGot, encodedWant)
		}
	}
}
