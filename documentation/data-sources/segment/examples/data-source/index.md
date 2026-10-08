---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_segment."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1002, "body_sha256": "sha256:b036588badddedaa58ade90481eba5ef6e4b9b2c37a5c502973ea139f6dca208", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2e9bfc7c69e50200df48c3eca01ce203db4e0516ed11b6bf987fdf485d490ac4", "source_path": "examples/data-sources/xcsh_segment/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:segment:example:data-source", "parent_id": "xcsh-docs:data-sources:segment:examples", "path": "documentation/data-sources/segment/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2232010232021203-2221031113210023-0131212110111310-0003221321320031-2023113000120320-0311333313333233-1012011003000102-3021220000112222", "registry_path": "docs/guides/data-sources--segment--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_segment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["segmentCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment/data-source.tf`; digest `sha256:2e9bfc7c69e50200df48c3eca01ce203db4e0516ed11b6bf987fdf485d490ac4`.

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
