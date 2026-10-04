---
page_title: "xcsh_cloud_connect"
subcategory: ""
description: "Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud provider networks."
xcsh_docs: {"aliases": ["cloud connect"], "body_bytes": 1371, "body_sha256": "sha256:26779ab49e4d898849ec37f558572583b90df98f6f587442012656e953a0ea21", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_connect:reference", "xcsh-docs:data-sources:cloud_connect:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/cloud_connect/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000", "registry_path": "docs/data-sources/cloud_connect.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud provider networks.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cloud_connect

Breadcrumbs:

- xcsh_cloud_connect

Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud
provider networks.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/examples/)
