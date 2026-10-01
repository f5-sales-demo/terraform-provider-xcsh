---
page_title: "Resource"
subcategory: "Networking"
description: "Resource for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 1070, "body_sha256": "sha256:e17570e5dbecde647faa7913e8dc48b500d3aa0e09a45139a6cebe9a41ff611a", "canonical_id": "xcsh-docs:resources:virtual_network:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473", "source_path": "examples/resources/xcsh_virtual_network/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:virtual_network:example:resource", "parent_id": "xcsh-docs:resources:virtual_network:examples", "path": "docs/guides/resources--virtual_network--example--resource.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md)
- [Examples](resources--virtual_network--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_network/resource.tf`; digest `sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473`.

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

## Next pages

- [Examples](resources--virtual_network--examples.md)
- [xcsh_virtual_network](../resources/virtual_network.md)
