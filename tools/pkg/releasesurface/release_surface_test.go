package releasesurface

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadSMSv2ReleaseSurfaceIsExact(t *testing.T) {
	surface, err := Load(filepath.Join("..", "..", "..", "smsv2-release-surface.json"))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := surface.Resources, []string{"securemesh_site_v2", "smsv2_kvm_runtime_interface", "token", "registration_approval", "external_connector", "bgp", "virtual_site", "origin_pool", "http_loadbalancer", "dns_zone"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("resources = %v, want %v", got, want)
	}
	if got, want := surface.DataSources, []string{"network_regional_edges", "network_cdn", "network_secondary_dns_zone_transfer", "network_global_log_receiver", "network_dnslb_health_checks", "network_global_controller_sso_egress", "network_bot_defense", "network_data_intelligence", "network_customer_edge_defaults", "network_customer_edge_egress", "smsv2_contract", "smsv2_aws_runtime", "smsv2_kvm_runtime", "site_bgp_status", "site_upgrade_status", "site_image", "site_cloud_init", "site_registration", "site_registrations_by_site", "namespace", "dns_zone"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("data sources = %v, want %v", got, want)
	}
	if got, want := surface.Actions, []string{"site_upgrade_sw", "site_upgrade_os"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("actions = %v, want %v", got, want)
	}
	if len(surface.Functions) != 0 {
		t.Fatalf("functions = %v, want none", surface.Functions)
	}
}
