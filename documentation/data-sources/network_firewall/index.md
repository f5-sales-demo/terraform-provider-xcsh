---
page_title: "xcsh_network_firewall"
subcategory: "Security"
description: "Reads a Network Firewall created in the system namespace."
xcsh_docs: {"aliases": ["network firewall"], "body_bytes": 1377, "body_sha256": "sha256:a0d2e4659b87717497407727a56a7ce6308bfc28263f4c76175ede3a94033b45", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_firewall:reference", "xcsh-docs:data-sources:network_firewall:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_firewall:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_firewall/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112", "registry_path": "docs/data-sources/network_firewall.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads a Network Firewall created in the system namespace.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_firewall

Breadcrumbs:

- xcsh_network_firewall

Reads a Network Firewall created in the system namespace.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkFirewall by name
data "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.xcsh_network_firewall.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/examples/)
