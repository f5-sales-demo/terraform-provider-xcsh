---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1344, "body_sha256": "sha256:60eb54c0e5d45149105b1fe09fb8e7a8d065ca29e2ee1fa55d49aaf03349ff82", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2b23ed57ede50484e6ffd1c4ce6fdf9ba70d1bbaddf2b7e8ceb68afb488d14e2", "source_path": "examples/data-sources/xcsh_network_interface/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_interface:example:data-source", "parent_id": "xcsh-docs:data-sources:network_interface:examples", "path": "documentation/data-sources/network_interface/examples/data-source/index.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_interface/data-source.tf`; digest `sha256:2b23ed57ede50484e6ffd1c4ce6fdf9ba70d1bbaddf2b7e8ceb68afb488d14e2`.

```terraform
# NetworkInterface Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkInterface by name
data "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}

output "network_interface_id" {
  value = data.xcsh_network_interface.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/examples/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
