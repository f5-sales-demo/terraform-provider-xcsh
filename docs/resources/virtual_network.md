---
page_title: "xcsh_virtual_network"
subcategory: "Networking"
description: "xcsh_virtual_network for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 1476, "body_sha256": "sha256:95a3ed9a18fb43457a74e30a89a0d898a6a9460f60b89de018a95b6e163e7b95", "canonical_id": "xcsh-docs:resources:virtual_network:fundamentals", "child_ids": ["xcsh-docs:resources:virtual_network:reference", "xcsh-docs:resources:virtual_network:examples", "xcsh-docs:resources:virtual_network:import", "xcsh-docs:resources:virtual_network:timeouts"], "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:fundamentals", "parent_id": null, "path": "docs/resources/virtual_network.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_virtual_network for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_virtual_network

Breadcrumbs:

- xcsh_virtual_network

Manages virtual network in given namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `network_connector`.

- network_connector: Connect to external networks

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualNetwork Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--virtual_network--reference.md)
- [Examples](../guides/resources--virtual_network--examples.md)
- [Import](../guides/resources--virtual_network--import.md)
- [Timeouts](../guides/resources--virtual_network--timeouts.md)
