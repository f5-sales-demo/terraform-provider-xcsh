---
page_title: "xcsh_network_firewall"
subcategory: "Security"
description: "xcsh_network_firewall for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1513, "body_sha256": "sha256:5891cb8424d51185e20352c22eedbac623426871fe7b04195246216b7d2bbf31", "canonical_id": "xcsh-docs:resources:network_firewall:fundamentals", "child_ids": ["xcsh-docs:resources:network_firewall:reference", "xcsh-docs:resources:network_firewall:examples", "xcsh-docs:resources:network_firewall:import", "xcsh-docs:resources:network_firewall:timeouts"], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:fundamentals", "parent_id": null, "path": "docs/resources/network_firewall.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_firewall for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--network_firewall--reference.md)
- [Examples](../guides/resources--network_firewall--examples.md)
- [Import](../guides/resources--network_firewall--import.md)
- [Timeouts](../guides/resources--network_firewall--timeouts.md)
