---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 1381, "body_sha256": "sha256:e973e820022231c4a0d5d3f837f8fceda80403d55260916c262a9401cf9b2ce7", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b98a932b2c1a0b5e2b9c608b25b6b3f74aea2df606bdee9675da50e166ce11d", "source_path": "examples/data-sources/xcsh_bigip_virtual_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bigip_virtual_server:example:data-source", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:examples", "path": "documentation/data-sources/bigip_virtual_server/examples/data-source/index.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bigip_virtual_server/data-source.tf`; digest `sha256:4b98a932b2c1a0b5e2b9c608b25b6b3f74aea2df606bdee9675da50e166ce11d`.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/examples/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
