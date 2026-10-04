---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_connect."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1292, "body_sha256": "sha256:b962755457280eb6d1870a7eb0f92969e770c906d91e61f99b5a2795dca7e914", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:eed109fcd2f0afc5835242d94d0d30137f2b13d899b31f9e138f00b05427dcd6", "source_path": "examples/data-sources/xcsh_cloud_connect/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_connect:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_connect:examples", "path": "documentation/data-sources/cloud_connect/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3022330130322333-0112021300131120-0032310030220112-3100103123333021-1133133312302123-3211333121033220-3000230211200231-1021111103113332", "registry_path": "docs/guides/data-sources--cloud_connect--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_cloud_connect.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_connect/data-source.tf`; digest `sha256:eed109fcd2f0afc5835242d94d0d30137f2b13d899b31f9e138f00b05427dcd6`.

```terraform
# CloudConnect Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudConnect by name
data "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}

output "cloud_connect_id" {
  value = data.xcsh_cloud_connect.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/examples/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
