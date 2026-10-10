---
page_title: "xcsh_network_firewall"
subcategory: "Security"
description: "Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["network firewall"], "body_bytes": 1715, "body_sha256": "sha256:05333954478631395a1e72713c2d24bf7056d03b3adb54b4092baaff37f6957e", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_firewall:reference", "xcsh-docs:resources:network_firewall:examples", "xcsh-docs:resources:network_firewall:import", "xcsh-docs:resources:network_firewall:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/network_firewall/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333", "registry_path": "docs/resources/network_firewall.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["network_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_firewall

Breadcrumbs:

- xcsh_network_firewall

Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users
in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkFirewall Resource Example
# Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkFirewall configuration
resource "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/lifecycle/timeouts/)
