// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNetworkAllowlistCIDRConversionAndOrdering(t *testing.T) {
	got, err := networkAllowlistCIDRs([]string{"203.0.113.9", "192.0.2.0/24", "203.0.113.9", "198.51.100.7/32"})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"192.0.2.0/24", "198.51.100.7/32", "203.0.113.9/32"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CIDRs = %v, want %v", got, want)
	}
}

func TestNetworkAllowlistPinnedManifest(t *testing.T) {
	version, err := os.ReadFile("../../tools/spec-version.txt")
	if err != nil {
		t.Fatal(err)
	}
	wantReleaseTag := strings.TrimSpace(string(version))
	if bundledNetworkAllowlist.APIReleaseTag != wantReleaseTag {
		t.Fatalf("release tag = %q, want %q from tools/spec-version.txt", bundledNetworkAllowlist.APIReleaseTag, wantReleaseTag)
	}
	if bundledNetworkAllowlist.SourceSHA256 != "0bb6fd6bd561aef1ea9fbc772119d4cfea8c26cbe82d7ea6e76f9aeb03366555" {
		t.Fatalf("source digest = %q", bundledNetworkAllowlist.SourceSHA256)
	}
	if got := bundledNetworkAllowlist.Services["global_log_receiver"].SourceEntries[2]; got != "3.64.239.6" {
		t.Fatalf("mixed source entry = %q", got)
	}
	if got := bundledNetworkAllowlist.CustomerEdge.SecureMeshV2.Domains[0]; got != ".volterra.io" {
		t.Fatalf("leading-dot domain = %q", got)
	}
}

func TestNetworkAllowlistRegionalEdgesCredentialFreePlan(t *testing.T) {
	clearCredentialEnvironment(t)
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testProviderFactories(),
		Steps: []resource.TestStep{{
			Config: "provider \"xcsh\" {}\n" +
				"data \"xcsh_network_regional_edges\" \"selected\" {\n" +
				"  regions = [\"europe\", \"americas\"]\n" +
				"}\n",
			PlanOnly: true,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("data.xcsh_network_regional_edges.selected", "regions.#", "2"),
				resource.TestCheckResourceAttrSet("data.xcsh_network_regional_edges.selected", "source_sha256"),
				resource.TestCheckResourceAttrSet("data.xcsh_network_regional_edges.selected", "cidr_blocks_by_region.americas.#"),
			),
		}},
	})
}

func TestNetworkAllowlistRejectsUnknownRegion(t *testing.T) {
	clearCredentialEnvironment(t)
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testProviderFactories(),
		Steps: []resource.TestStep{{
			Config: "provider \"xcsh\" {}\n" +
				"data \"xcsh_network_data_intelligence\" \"selected\" {\n" +
				"  regions = [\"moon\"]\n" +
				"}\n",
			PlanOnly:    true,
			ExpectError: regexp.MustCompile("Unknown Network Allowlist Region"),
		}},
	})
}

func TestAPIBackedDataSourceRequiresCredentialsBeforeRequest(t *testing.T) {
	clearCredentialEnvironment(t)
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testProviderFactories(),
		Steps: []resource.TestStep{{
			Config: "provider \"xcsh\" { api_url = \"http://127.0.0.1:1\" }\n" +
				"data \"xcsh_namespace\" \"probe\" {\n" +
				"  namespace = \"system\"\n" +
				"  name = \"probe\"\n" +
				"}\n",
			PlanOnly:    true,
			ExpectError: regexp.MustCompile("F5XC API credentials are required"),
		}},
	})
}

func testProviderFactories() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"xcsh": providerserver.NewProtocol6WithError(New("test")()),
	}
}

func clearCredentialEnvironment(t *testing.T) {
	t.Helper()
	for _, name := range []string{"XCSH_API_TOKEN", "XCSH_P12_FILE", "XCSH_P12_PASSWORD", "XCSH_CERT", "XCSH_KEY", "XCSH_CACERT"} {
		t.Setenv(name, "")
	}
}

func TestStableNetworkAllowlistID(t *testing.T) {
	all := stableNetworkAllowlistID("regional_edges", []string{"americas", "europe"})
	reordered := stableNetworkAllowlistID("regional_edges", []string{"europe", "americas"})
	if all != reordered {
		t.Fatalf("ID depends on selection order: %q != %q", all, reordered)
	}
	if all == stableNetworkAllowlistID("regional_edges", []string{"americas"}) {
		t.Fatal("ID does not distinguish region selection")
	}
}

func TestNetworkAllowlistRegionsDefaultToAllPublished(t *testing.T) {
	available := map[string]networkAllowlistServiceGroup{
		"us": {},
		"eu": {},
	}
	got, err := selectedNetworkAllowlistRegions(available, types.SetNull(types.StringType))
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"eu", "us"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("default regions = %v, want %v", got, want)
	}
}

func TestNetworkAllowlistConstructorsExposeExpectedNames(t *testing.T) {
	p := &XCSHProvider{}
	got := dataSourceNames(t, p.DataSources(context.Background()))
	for _, want := range []string{
		"xcsh_network_regional_edges",
		"xcsh_network_cdn",
		"xcsh_network_secondary_dns_zone_transfer",
		"xcsh_network_global_log_receiver",
		"xcsh_network_dnslb_health_checks",
		"xcsh_network_global_controller_sso_egress",
		"xcsh_network_bot_defense",
		"xcsh_network_data_intelligence",
		"xcsh_network_customer_edge_defaults",
		"xcsh_network_customer_edge_egress",
	} {
		found := false
		for _, name := range got {
			if name == want {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("missing data source %s in %v", want, got)
		}
	}
}
