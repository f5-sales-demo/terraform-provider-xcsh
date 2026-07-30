// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package openapi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/naming"
)

// decoratedSiteResources are the resources whose API objects F5 XC decorates with
// hardware/OS discovery labels, by Terraform resource name.
//
// Measured in the f5-sales-demo tenant on 2026-07-30 by listing each kind and
// counting objects carrying `domain`, `host-os-version` or an `hw-` key:
// aws_vpc_site 4/4, azure_vnet_site 2/2, securemesh_site_v2 9/13. fleet 0/19 and
// virtual_site 0/6 carried none, and are deliberately excluded. The kinds with no
// live objects to measure share the node-backed site shape and the same generator.
var decoratedSiteResources = []string{
	"aws_tgw_site",
	"aws_vpc_site",
	"azure_vnet_site",
	"gcp_vpc_site",
	"securemesh_site",
	"securemesh_site_v2",
	"site",
	"voltstack_site",
}

// The data file is keyed by TitleCase, and TitleCase is not guessable — the
// generator renders `azure_vnet_site` as `AzureVNETSite`, not `AzureVnetSite`. A key
// that does not match simply never fires, silently leaving that resource on the
// broken behaviour. That happened once already during #1391: four of the eight keys
// were wrong and only `securemesh_site_v2` was actually filtered. Derive the key the
// same way the generator does instead of trusting the spelling in the file.
func TestDiscoveredSiteLabels_KeysMatchGeneratedTitleCase(t *testing.T) {
	for _, resource := range decoratedSiteResources {
		titleCase := naming.ToResourceTypeName(resource)
		if !LoadDiscoveredSiteLabels(titleCase) {
			t.Errorf("resource %q (TitleCase %q) is not opted in to discovery-label filtering; "+
				"tools/discovered-site-labels.json is missing that key or spells it differently",
				resource, titleCase)
		}
	}
}

// The converse: no key in the file may be dead. A stale key is a key that stopped
// matching, which fails open.
func TestDiscoveredSiteLabels_NoKeyIsUnaccountedFor(t *testing.T) {
	expected := map[string]bool{}
	for _, resource := range decoratedSiteResources {
		expected[naming.ToResourceTypeName(resource)] = true
	}

	data, err := os.ReadFile(filepath.Join("..", "..", "discovered-site-labels.json"))
	if err != nil {
		t.Fatalf("reading tools/discovered-site-labels.json: %v", err)
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("parsing tools/discovered-site-labels.json: %v", err)
	}
	for key := range raw {
		if key == "_comment" {
			continue
		}
		if !expected[key] {
			t.Errorf("tools/discovered-site-labels.json has key %q, which no resource in "+
				"decoratedSiteResources produces — it is dead and filters nothing", key)
		}
	}
}

// A resource F5 XC does not decorate must keep prefix-only filtering, so a user can
// still own a label called `domain` on it.
func TestDiscoveredSiteLabels_UndecoratedResourcesAreNotOptedIn(t *testing.T) {
	for _, resource := range []string{"fleet", "virtual_site", "site_mesh_group", "http_loadbalancer"} {
		titleCase := naming.ToResourceTypeName(resource)
		if LoadDiscoveredSiteLabels(titleCase) {
			t.Errorf("resource %q (TitleCase %q) must not filter the discovery labels: "+
				"nothing writes them there, and the names are ordinary enough for a user to own",
				resource, titleCase)
		}
	}
}
