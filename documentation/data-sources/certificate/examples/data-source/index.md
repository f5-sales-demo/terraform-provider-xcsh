---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_certificate."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1043, "body_sha256": "sha256:654e62a42680bb52304f1352ead1149fe29a103f179b7242c1617970a246a392", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e", "source_path": "examples/data-sources/xcsh_certificate/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certificate:example:data-source", "parent_id": "xcsh-docs:data-sources:certificate:examples", "path": "documentation/data-sources/certificate/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0230322312131310-1100223021111133-2013003200312120-3001030011211123-2210201231002213-3101333023233123-2120323312113233-3103212213321302", "registry_path": "docs/guides/data-sources--certificate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["certificateCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certificate/data-source.tf`; digest `sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e`.

```terraform
# Certificate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Certificate by name
data "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"
}

output "certificate_id" {
  value = data.xcsh_certificate.example.id
}
```
