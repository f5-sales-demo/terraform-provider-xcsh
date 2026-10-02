---
page_title: "xcsh_network_interface"
subcategory: ""
description: "Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device. it is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["network interface"], "body_bytes": 1750, "body_sha256": "sha256:c0a6b3571e9a1c2552d1428c587f09d00d48d8413fc4334cfbfc6f211d7760b7", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:reference", "xcsh-docs:resources:network_interface:examples", "xcsh-docs:resources:network_interface:import", "xcsh-docs:resources:network_interface:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/network_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310", "registry_path": "docs/resources/network_interface.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device. it is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/lifecycle/timeouts/)
