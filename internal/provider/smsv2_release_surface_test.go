package provider

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestSMSv2ReleaseSurfaceIsExact(t *testing.T) {
	p := &XCSHProvider{}
	resources := resourceNames(t, p.Resources(context.Background()))
	dataSources := dataSourceNames(t, p.DataSources(context.Background()))
	actions := actionNames(t, p.Actions(context.Background()))
	assertSurfaceNames(t, resources, []string{"xcsh_bgp", "xcsh_dns_zone", "xcsh_external_connector", "xcsh_http_loadbalancer", "xcsh_origin_pool", "xcsh_registration_approval", "xcsh_securemesh_site_v2", "xcsh_token", "xcsh_virtual_site"})
	assertSurfaceNames(t, dataSources, []string{"xcsh_dns_zone", "xcsh_namespace", "xcsh_site_bgp_status", "xcsh_site_cloud_init", "xcsh_site_image", "xcsh_site_registration", "xcsh_site_registrations_by_site", "xcsh_site_upgrade_status", "xcsh_smsv2_aws_runtime", "xcsh_smsv2_contract", "xcsh_smsv2_kvm_runtime"})
	assertSurfaceNames(t, actions, []string{"xcsh_site_upgrade_os", "xcsh_site_upgrade_sw"})
}

func resourceNames(t *testing.T, constructors []func() resource.Resource) []string {
	t.Helper()
	names := make([]string, 0, len(constructors))
	for _, constructor := range constructors {
		var response resource.MetadataResponse
		constructor().Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "xcsh"}, &response)
		names = append(names, response.TypeName)
	}
	return names
}
func dataSourceNames(t *testing.T, constructors []func() datasource.DataSource) []string {
	t.Helper()
	names := make([]string, 0, len(constructors))
	for _, constructor := range constructors {
		var response datasource.MetadataResponse
		constructor().Metadata(context.Background(), datasource.MetadataRequest{ProviderTypeName: "xcsh"}, &response)
		names = append(names, response.TypeName)
	}
	return names
}
func actionNames(t *testing.T, constructors []func() action.Action) []string {
	t.Helper()
	names := make([]string, 0, len(constructors))
	for _, constructor := range constructors {
		var response action.MetadataResponse
		constructor().Metadata(context.Background(), action.MetadataRequest{ProviderTypeName: "xcsh"}, &response)
		names = append(names, response.TypeName)
	}
	return names
}
func assertSurfaceNames(t *testing.T, got, want []string) {
	t.Helper()
	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("surface = %v, want %v", got, want)
	}
}
