package releasesurface

import (
	"path/filepath"
	"slices"
	"testing"
)

func TestLoadProviderReleaseSurfaceIsExact(t *testing.T) {
	surface, err := Load(filepath.Join("..", "..", "..", "provider-release-surface.json"))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(surface.Resources), 126; got != want {
		t.Fatalf("resources = %d, want %d", got, want)
	}
	if got, want := len(surface.DataSources), 193; got != want {
		t.Fatalf("data sources = %d, want %d", got, want)
	}
	if got, want := len(surface.Actions), 11; got != want {
		t.Fatalf("actions = %d, want %d", got, want)
	}
	if got, want := len(surface.EphemeralResources), 2; got != want {
		t.Fatalf("ephemeral resources = %d, want %d", got, want)
	}
	for _, name := range []string{"namespace", "protected_domain", "public_ip_binding", "smsv2_kvm_runtime_interface"} {
		if !slices.Contains(surface.Resources, name) {
			t.Errorf("resources does not contain %q", name)
		}
	}
	for _, name := range []string{"addon_service_activation_status", "network_regional_edges", "smsv2_contract"} {
		if !slices.Contains(surface.DataSources, name) {
			t.Errorf("data sources does not contain %q", name)
		}
	}
	for _, name := range []string{"aws_vpc_site", "azure_vnet_site", "gcp_vpc_site", "securemesh_site", "cloud_connect"} {
		if slices.Contains(surface.Resources, name) || slices.Contains(surface.DataSources, name) {
			t.Errorf("retired type %q remains in the release surface", name)
		}
	}
	if len(surface.Functions) != 0 {
		t.Fatalf("functions = %v, want none", surface.Functions)
	}
}
