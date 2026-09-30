---
page_title: "xcsh_network_interface"
subcategory: ""
description: "xcsh_network_interface for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1462, "body_sha256": "sha256:cb9a1ac5568b3bf126499208cd0d55c1b48b4168588dda6228caae0cab0dcb78", "canonical_id": "xcsh-docs:resources:network_interface:fundamentals", "child_ids": ["xcsh-docs:resources:network_interface:reference", "xcsh-docs:resources:network_interface:examples", "xcsh-docs:resources:network_interface:import", "xcsh-docs:resources:network_interface:timeouts"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:fundamentals", "parent_id": null, "path": "docs/resources/network_interface.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_interface for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_network_interface

Breadcrumbs:

- xcsh_network_interface

Manages a Network Interface resource in F5 Distributed Cloud for network interface represents
configuration of a network device. it is created by users in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkInterface Resource Example
# Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkInterface configuration
resource "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--network_interface--reference.md)
- [Examples](../guides/resources--network_interface--examples.md)
- [Import](../guides/resources--network_interface--import.md)
- [Timeouts](../guides/resources--network_interface--timeouts.md)
