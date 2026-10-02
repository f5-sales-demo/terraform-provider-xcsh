---
page_title: "xcsh_segment"
subcategory: ""
description: "Manages a Segment resource in F5 Distributed Cloud for segment. configuration."
xcsh_docs: {"aliases": ["segment"], "body_bytes": 1263, "body_sha256": "sha256:865a47c6f33557cace56b61dc7c3cc0636dd26fa1a0561bda31435edfa0d71b3", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:segment:reference", "xcsh-docs:data-sources:segment:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:segment:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/segment/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1322202313211133-3232200032023212-3030131113213211-1232130001132103-2232301233230330-2022002332021123-3031321121233232-3101311003121100", "registry_path": "docs/data-sources/segment.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Segment resource in F5 Distributed Cloud for segment. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["segmentCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_segment

Breadcrumbs:

- xcsh_segment

Manages a Segment resource in F5 Distributed Cloud for segment. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Segment Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Segment by name
data "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.xcsh_segment.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/examples/)
