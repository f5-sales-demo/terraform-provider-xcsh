---
page_title: "xcsh_network_firewall"
subcategory: "Security"
description: "Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["network firewall"], "body_bytes": 1702, "body_sha256": "sha256:14db7303b5baaa9d22bcfd8dda2c6fc3873609afaa88f47c23b80118e569a3b8", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_firewall:reference", "xcsh-docs:resources:network_firewall:examples", "xcsh-docs:resources:network_firewall:import", "xcsh-docs:resources:network_firewall:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/network_firewall/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0002103131303200-3202220110210100-2001233131212311-3322200221303002-0011132331222322-2101102312302011-0113033313013211-2001132200001333", "registry_path": "docs/resources/network_firewall.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/lifecycle/timeouts/)
