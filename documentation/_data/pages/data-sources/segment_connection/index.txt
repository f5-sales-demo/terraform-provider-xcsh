---
page_title: "xcsh_segment_connection"
subcategory: ""
description: "xcsh_segment_connection for xcsh_segment_connection."
xcsh_docs: {"aliases": [], "body_bytes": 1444, "body_sha256": "sha256:26b32898058e301a16ee541585b4ec5abc7543df189690506f3405c683c4c4f3", "child_ids": ["xcsh-docs:data-sources:segment_connection:reference", "xcsh-docs:data-sources:segment_connection:examples"], "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:fundamentals", "parent_id": null, "path": "documentation/data-sources/segment_connection/index.md", "provider_name": "segment_connection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_segment_connection for xcsh_segment_connection.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_segment_connection

Breadcrumbs:

- xcsh_segment_connection

Manages a Segment Connection resource in F5 Distributed Cloud for segment connector specification.
configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SegmentConnection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SegmentConnection by name
data "xcsh_segment_connection" "example" {
  name      = "example-segment-connection"
  namespace = "staging"
}

output "segment_connection_id" {
  value = data.xcsh_segment_connection.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/examples/)
