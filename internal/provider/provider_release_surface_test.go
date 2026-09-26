package provider

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/releasesurface"
)

func TestProviderReleaseSurfaceIsExact(t *testing.T) {
	p := &XCSHProvider{}
	surface, err := releasesurface.Load("../../provider-release-surface.json")
	if err != nil {
		t.Fatalf("load release surface: %v", err)
	}
	resources := resourceNames(t, p.Resources(context.Background()))
	dataSources := dataSourceNames(t, p.DataSources(context.Background()))
	actions := actionNames(t, p.Actions(context.Background()))
	assertSurfaceNames(t, resources, prefixed(surface.Resources))
	assertSurfaceNames(t, dataSources, prefixed(surface.DataSources))
	assertSurfaceNames(t, actions, prefixed(surface.Actions))

	var functions []func() function.Function
	if withFunctions, ok := any(p).(frameworkprovider.ProviderWithFunctions); ok {
		functions = withFunctions.Functions(context.Background())
	}
	if len(functions) != len(surface.Functions) {
		t.Fatalf("registered functions = %d, release surface requires %d", len(functions), len(surface.Functions))
	}
}

func prefixed(names []string) []string {
	result := make([]string, len(names))
	for index, name := range names {
		result[index] = "xcsh_" + name
	}
	return result
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
