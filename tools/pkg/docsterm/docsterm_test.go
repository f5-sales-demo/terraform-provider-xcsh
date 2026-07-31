// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package docsterm

import (
	"strings"
	"testing"
)

// TestFixUpstreamTerminology_PreservesEncodingBase64 is the regression test for
// the S2b fix: the terminology normaliser must NOT lowercase "Base64" inside the
// real API enum token "EncodingBase64" (which would corrupt it into
// "Encodingbase64"). Go's regexp has no lookbehind, so a naive Base64->base64
// rewrite cannot exclude the enum token; the rewrite was therefore removed.
func TestFixUpstreamTerminology_PreservesEncodingBase64(t *testing.T) {
	input := "Possible values are `EncodingNone`, `EncodingBase64`"
	got := FixUpstreamTerminology(input)

	if strings.Contains(got, "Encodingbase64") {
		t.Errorf("FixUpstreamTerminology corrupted enum token: got %q, must not contain %q", got, "Encodingbase64")
	}
	if !strings.Contains(got, "EncodingBase64") {
		t.Errorf("FixUpstreamTerminology dropped enum token: got %q, must contain %q", got, "EncodingBase64")
	}
}

// TestFixUpstreamTerminology_NormalisesInternet is a characterization test proving
// the function still performs a real terminology transform (lowercasing prose
// "Internet" -> "internet"), i.e. the extraction did not gut the logic.
func TestFixUpstreamTerminology_NormalisesInternet(t *testing.T) {
	input := "Traffic reaches the Internet directly."
	want := "Traffic reaches the internet directly."
	got := FixUpstreamTerminology(input)
	if got != want {
		t.Errorf("FixUpstreamTerminology(%q) = %q, want %q", input, got, want)
	}
}

// Terminology must never rewrite an identifier. Generated provider documentation is
// mostly identifiers, and the `javascript` -> `JavaScript` entry rewrote the real
// attribute `javascript_location` into `JavaScript_location` — HCL that Terraform
// rejects — in five committed pages (#1414). URLs were already protected here; code
// spans were not, which is where attribute names live.
func TestFixUpstreamTerminologyLeavesCodeSpansAlone(t *testing.T) {
	cases := []struct {
		name string
		in   string
	}{
		{
			"attribute name in a link label",
			"&#x2022; [`javascript_location`](#location-a9a4d9) - Optional String",
		},
		{"bare code span", "Set `javascript_mode` to control caching."},
		{"docker-prefixed attribute", "The `docker_registry` block configures it."},
		{"azure-prefixed attribute", "Use `azure_vnet_site` for orchestrated sites."},
		{"ubuntu value", "Set `os_flavor` to `ubuntu` for the default image."},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FixUpstreamTerminology(tc.in); got != tc.in {
				t.Errorf("code span was rewritten\n in: %s\nout: %s", tc.in, got)
			}
		})
	}
}

// Prose outside code spans must still be corrected — the point is to distinguish the
// two, not to disable terminology.
func TestFixUpstreamTerminologyStillCorrectsProse(t *testing.T) {
	got := FixUpstreamTerminology("Bot Defense injects javascript into the page.")
	want := "Bot Defense injects JavaScript into the page."
	if got != want {
		t.Errorf("prose not corrected\ngot:  %s\nwant: %s", got, want)
	}
}

// Both at once: the identifier is preserved and the prose around it is corrected.
func TestFixUpstreamTerminologyMixedLine(t *testing.T) {
	in := "Web Client javascript Mode. Set `javascript_mode` to change it."
	want := "Web Client JavaScript Mode. Set `javascript_mode` to change it."
	if got := FixUpstreamTerminology(in); got != want {
		t.Errorf("mixed line\ngot:  %s\nwant: %s", got, want)
	}
}

// Fenced blocks are code for the same reason inline spans are. Generated resource
// pages embed the example HCL verbatim, and `azure` -> `Azure` rewrote the resource
// NAME inside them: examples/resources/xcsh_azure_vnet_site/resource.tf ships
// `example-azure-vnet-site` while the fenced copy in docs/resources/azure_vnet_site.md
// read `example-Azure-vnet-site`. A reader copying the documented example got
// different HCL from the one in the repository.
func TestFixUpstreamTerminologyLeavesFencedBlocksAlone(t *testing.T) {
	in := "Prose about javascript here.\n\n" +
		"```terraform\n" +
		"resource \"xcsh_azure_vnet_site\" \"example\" {\n" +
		"  name      = \"example-azure-vnet-site\"\n" +
		"  image     = \"ubuntu\"\n" +
		"  registry  = \"docker.io\"\n" +
		"}\n" +
		"```\n\n" +
		"More javascript prose.\n"

	got := FixUpstreamTerminology(in)

	for _, forbidden := range []string{
		"example-Azure-vnet-site",
		"\"Ubuntu\"",
		"Docker.io",
		"xcsh_Azure_vnet_site",
	} {
		if strings.Contains(got, forbidden) {
			t.Errorf("fenced block was rewritten: found %q in\n%s", forbidden, got)
		}
	}

	// Prose on both sides of the fence is still corrected.
	if strings.Count(got, "JavaScript") != 2 {
		t.Errorf("prose around the fence should still be corrected, got:\n%s", got)
	}
}

// `[Enum: ...]` lists the exact literals the provider accepts. They sit outside
// backticks, so terminology rewrote them: the generated azure_vnet_site page
// advertised `[Enum: Azure-byol-multi-nic-voltmesh]` while the resource's own
// stringvalidator.OneOf accepts only "azure-byol-multi-nic-voltmesh". A reader
// copying the documented value gets a Terraform validation error.
func TestFixUpstreamTerminologyLeavesEnumTokensAlone(t *testing.T) {
	in := "[Enum: azure-byol-multi-nic-voltmesh] Name for Azure certified hardware."
	want := "[Enum: azure-byol-multi-nic-voltmesh] Name for Azure certified hardware."
	if got := FixUpstreamTerminology(in); got != want {
		t.Errorf("enum token rewritten\ngot:  %s\nwant: %s", got, want)
	}

	multi := "[Enum: ubuntu|docker|azure] pick one"
	if got := FixUpstreamTerminology(multi); got != multi {
		t.Errorf("multi-value enum rewritten\ngot:  %s\nwant: %s", got, multi)
	}
}
