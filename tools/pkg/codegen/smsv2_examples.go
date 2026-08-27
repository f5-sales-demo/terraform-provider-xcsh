// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package codegen

import (
	"fmt"
	"os"
	"path/filepath"
)

// SecuremeshSiteV2ProviderChoices is the complete provider-choice group from
// the released enriched SMSv2 contract.
var SecuremeshSiteV2ProviderChoices = []string{
	"aws",
	"azure",
	"baremetal",
	"equinix",
	"gcp",
	"kvm",
	"nutanix",
	"oci",
	"openshift_virtualization",
	"openstack",
	"vmware",
}

const smsv2ExamplePreamble = `terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

`

func writeSMSv2Example(root, variant, body string) error {
	directory := filepath.Join(root, "examples", "resources", "xcsh_securemesh_site_v2", "variants", variant)
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(directory, "resource.tf"), []byte(smsv2ExamplePreamble+body), 0o644)
}

// WriteSecuremeshSiteV2Examples writes one structurally complete example for
// every current provider choice plus the named Segment VRF contract.
func WriteSecuremeshSiteV2Examples(root string) error {
	for _, providerChoice := range SecuremeshSiteV2ProviderChoices {
		body := fmt.Sprintf(`resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2-%s"
  namespace = "system"

  %s {
    not_managed {}
  }

  disable_ha {}
  no_network_policy {}
  no_forward_proxy {}
  logs_streaming_disabled {}
  block_all_services {}
}
`, providerChoice, providerChoice)
		if err := writeSMSv2Example(root, providerChoice, body); err != nil {
			return err
		}
	}

	return writeSMSv2Example(root, "segment-vrf", `resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2-segment-vrf"
  namespace = "system"

  baremetal {
    not_managed {}
  }

  disable_ha {}
  no_network_policy {}
  no_forward_proxy {}
  logs_streaming_disabled {}
  block_all_services {}

  segment_vrf {
    segment_network {
      name      = "example-segment"
      namespace = "system"
    }
  }
}
`)
}
