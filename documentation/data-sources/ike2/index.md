---
page_title: "xcsh_ike2"
subcategory: ""
description: "Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration."
xcsh_docs: {"aliases": ["ike2"], "body_bytes": 1269, "body_sha256": "sha256:0e29f611ee4fbb0b3721cb38a7bc6687caada1184722de03410966f7c9a915a7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:ike2:reference", "xcsh-docs:data-sources:ike2:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike2:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/ike2/index.md", "product": "distributed-cloud", "provider_name": "ike2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032", "registry_path": "docs/data-sources/ike2.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike2/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["ike2CreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike2 by name
data "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}

output "ike2_id" {
  value = data.xcsh_ike2.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike2/examples/)
