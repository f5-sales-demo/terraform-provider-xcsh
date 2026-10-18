// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/netip"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/networkallowlist"
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
	// This is compiled-artifact provenance. A pending spec pin changes before
	// canonical regeneration, so it cannot describe this binary's embedded data.
	var manifest struct {
		APIReleaseTag string `json:"api_release_tag"`
		SourceSHA256  string `json:"source_sha256"`
	}
	if err := json.Unmarshal([]byte(bundledNetworkAllowlistJSON), &manifest); err != nil {
		t.Fatal(err)
	}
	if !regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+$`).MatchString(manifest.APIReleaseTag) {
		t.Fatalf("invalid embedded release tag %q", manifest.APIReleaseTag)
	}
	if bundledNetworkAllowlist.APIReleaseTag != manifest.APIReleaseTag {
		t.Fatal("embedded provenance was not preserved")
	}
	if err := reviewedNetworkAllowlist([]byte(bundledNetworkAllowlistJSON)); err != nil {
		t.Fatal(err)
	}
	if !regexp.MustCompile(`^[a-f0-9]{64}$`).MatchString(manifest.SourceSHA256) || bundledNetworkAllowlist.SourceSHA256 != manifest.SourceSHA256 {
		t.Fatalf("source digest = %q", bundledNetworkAllowlist.SourceSHA256)
	}
	// The exact address identity is bound by the reviewed full artifact digest
	// above. Verify the representative mixed entry remains a bare IPv4 address
	// without copying its operational value into a separate test fixture.
	logReceiver, err := netip.ParseAddr(bundledNetworkAllowlist.Services["global_log_receiver"].SourceEntries[2])
	if err != nil || !logReceiver.Is4() {
		t.Fatal("mixed log receiver source entry must be a bare IPv4 address")
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

// Review content independently of its release label. OpenAPI extraction binds the
// label to the measured release receipt; new releases with identical reviewed
// content do not require a version-specific exception.
func reviewedNetworkAllowlist(raw []byte) error {
	var artifact networkallowlist.Artifact
	if !json.Valid(raw) {
		return fmt.Errorf("invalid network allowlist JSON")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&artifact); err != nil {
		return err
	}
	if !regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+$`).MatchString(artifact.APIReleaseTag) {
		return fmt.Errorf("unresolved network allowlist release")
	}
	artifact.APIReleaseTag = ""
	content, err := json.Marshal(artifact)
	if err != nil {
		return err
	}
	// Reviewed against immutable v12.1.1 OpenAPI SHA256
	// d4334dbf10c5a2b7d367de8d76356886a646ba017edf53e8d211c23365fc9f60.
	if fmt.Sprintf("%x", sha256.Sum256(content)) != "a3c293cc9e4b7e859bcba83fb8e672c25ca76ad7d72d49edd8475398a8367685" {
		return fmt.Errorf("unreviewed network allowlist content")
	}
	return nil
}

func TestNetworkAllowlistRejectsUnreviewedProvenance(t *testing.T) {
	var original map[string]any
	if err := json.Unmarshal([]byte(bundledNetworkAllowlistJSON), &original); err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"api_release_tag", "source_sha256", "source_url", "generated_at"} {
		t.Run(field, func(t *testing.T) {
			changed := make(map[string]any, len(original))
			for key, value := range original {
				changed[key] = value
			}
			changed[field] = "unreviewed"
			raw, err := json.Marshal(changed)
			if err != nil {
				t.Fatal(err)
			}
			if reviewedNetworkAllowlist(raw) == nil {
				t.Fatal("accepted changed artifact")
			}
		})
	}
}

func TestNetworkAllowlistReviewedContentAcceptsNewReleaseLabel(t *testing.T) {
	raw := strings.Replace(bundledNetworkAllowlistJSON, bundledNetworkAllowlist.APIReleaseTag, "v999.0.0", 1)
	if err := reviewedNetworkAllowlist([]byte(raw)); err != nil {
		t.Fatalf("unchanged reviewed content required a version exception: %v", err)
	}
}
