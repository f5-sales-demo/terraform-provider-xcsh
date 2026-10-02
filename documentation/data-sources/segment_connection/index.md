---
page_title: "xcsh_segment_connection"
subcategory: ""
description: "Manages a Segment Connection resource in F5 Distributed Cloud for segment connector specification. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["segment connection"], "body_bytes": 1444, "body_sha256": "sha256:26b32898058e301a16ee541585b4ec5abc7543df189690506f3405c683c4c4f3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:segment_connection:reference", "xcsh-docs:data-sources:segment_connection:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/segment_connection/index.md", "product": "distributed-cloud", "provider_name": "segment_connection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0030011210101331-2323210032230002-1133223312021002-2000100321331332-0221323131333131-3323103302112300-2012332113200223-2222011321202323", "registry_path": "docs/data-sources/segment_connection.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Segment Connection resource in F5 Distributed Cloud for segment connector specification. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
