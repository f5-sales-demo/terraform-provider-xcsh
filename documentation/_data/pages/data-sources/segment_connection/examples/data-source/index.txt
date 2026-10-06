---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_segment_connection."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1111, "body_sha256": "sha256:d0b3ea50ecb617165616b207021cef78658f63cc3ff2f6884445d3bb7ae093be", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:segment_connection:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d15a7947bb7cd84484433c713ef71c626c3d8bc137b6e6bda822fbc845d28228", "source_path": "examples/data-sources/xcsh_segment_connection/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:segment_connection:example:data-source", "parent_id": "xcsh-docs:data-sources:segment_connection:examples", "path": "documentation/data-sources/segment_connection/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "segment_connection", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0013320203022301-1021301010103003-1011302323121122-3120003012320313-0222211130122322-0322210203003023-0032230222331211-2213301133113003", "registry_path": "docs/guides/data-sources--segment_connection--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/segment_connection/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_segment_connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_segment_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/segment_connection/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment_connection/data-source.tf`; digest `sha256:d15a7947bb7cd84484433c713ef71c626c3d8bc137b6e6bda822fbc845d28228`.

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
