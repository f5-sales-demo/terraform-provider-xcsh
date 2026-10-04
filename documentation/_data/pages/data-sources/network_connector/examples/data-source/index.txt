---
page_title: "Data source"
subcategory: "Networking"
description: "Data source for xcsh_network_connector."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1344, "body_sha256": "sha256:6363b19c9e772638f0e51d7f0342b089ebe04843ae3e2eeecf8c41ed4f6dbb36", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:842679078281a092ba5ea7fed221d6b2c452c242a5d7cc3a8d19f42931b95be3", "source_path": "examples/data-sources/xcsh_network_connector/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_connector:example:data-source", "parent_id": "xcsh-docs:data-sources:network_connector:examples", "path": "documentation/data-sources/network_connector/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0023000011301211-0311132300203303-1010013212013123-0311323010132310-3232010101200001-0331121220320220-2333103333313230-0012312220023213", "registry_path": "docs/guides/data-sources--network_connector--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_network_connector.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["network_connectorCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_connector/data-source.tf`; digest `sha256:842679078281a092ba5ea7fed221d6b2c452c242a5d7cc3a8d19f42931b95be3`.

```terraform
# NetworkConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkConnector by name
data "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}

output "network_connector_id" {
  value = data.xcsh_network_connector.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/examples/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
