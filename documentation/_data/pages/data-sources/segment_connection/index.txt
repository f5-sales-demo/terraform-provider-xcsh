---
page_title: "xcsh_segment_connection"
subcategory: ""
description: "Reads Segment Connection information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["segment connection"], "body_bytes": 1383, "body_sha256": "sha256:383d8ceabf66b8c4e15320673b813bb383b4aebcca316350a5281abb99a2ccf8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:segment_connection:reference", "xcsh-docs:data-sources:segment_connection:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment_connection:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/segment_connection/index.md", "product": "distributed-cloud", "provider_name": "segment_connection", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0030011210101331-2323210032230002-1133223312021002-2000100321331332-0221323131333131-3323103302112300-2012332113200223-2222011321202323", "registry_path": "docs/data-sources/segment_connection.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Reads Segment Connection information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": [], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_segment_connection

Breadcrumbs:

- xcsh_segment_connection

Reads Segment Connection information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/examples/)
