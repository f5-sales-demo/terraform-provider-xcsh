---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 1076, "body_sha256": "sha256:ccedcc00cba0fc15cf1c9e6988be049f5d8248714f13935e31ad6e3bf5cf2fd0", "canonical_id": "xcsh-docs:data-sources:bigip_virtual_server:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b98a932b2c1a0b5e2b9c608b25b6b3f74aea2df606bdee9675da50e166ce11d", "source_path": "examples/data-sources/xcsh_bigip_virtual_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bigip_virtual_server:example:data-source", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:examples", "path": "docs/guides/data-sources--bigip_virtual_server--example--data-source.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
- [Examples](data-sources--bigip_virtual_server--examples.md)
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

- [Examples](data-sources--bigip_virtual_server--examples.md)
- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
