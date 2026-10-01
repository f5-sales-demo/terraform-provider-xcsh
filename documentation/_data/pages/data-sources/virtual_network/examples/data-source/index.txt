---
page_title: "Data source"
subcategory: "Networking"
description: "Data source for xcsh_virtual_network."
xcsh_docs: {"aliases": [], "body_bytes": 1317, "body_sha256": "sha256:bb32c579fb8f40655fa7c04766e27660e467bcca645f42dff59848c5c35e5a79", "child_ids": [], "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:82f2d94c3d011a8d200fcd015e8a3406aea69e2a77a8f9d9eea9796b4b19e3bd", "source_path": "examples/data-sources/xcsh_virtual_network/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_network:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_network:examples", "path": "documentation/data-sources/virtual_network/examples/data-source/index.md", "provider_name": "virtual_network", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_virtual_network.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_network/data-source.tf`; digest `sha256:82f2d94c3d011a8d200fcd015e8a3406aea69e2a77a8f9d9eea9796b4b19e3bd`.

```terraform
# VirtualNetwork Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualNetwork by name
data "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}

output "virtual_network_id" {
  value = data.xcsh_virtual_network.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/examples/)
- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
