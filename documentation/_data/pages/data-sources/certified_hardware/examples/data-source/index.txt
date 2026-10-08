---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_certified_hardware."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1111, "body_sha256": "sha256:aea982dc628b558965886f22913d69f17a069227c31d84d1ab07f0c340909916", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba", "source_path": "examples/data-sources/xcsh_certified_hardware/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certified_hardware:example:data-source", "parent_id": "xcsh-docs:data-sources:certified_hardware:examples", "path": "documentation/data-sources/certified_hardware/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1001223122213112-2321321330020223-3001031021223130-0312000001213120-2202031212330020-1033003021022202-3021321100302202-2313303211212213", "registry_path": "docs/guides/data-sources--certified_hardware--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_certified_hardware.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certified_hardware/data-source.tf`; digest `sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba`.

```terraform
# CertifiedHardware Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertifiedHardware by name
data "xcsh_certified_hardware" "example" {
  name      = "example-certified-hardware"
  namespace = "staging"
}

output "certified_hardware_id" {
  value = data.xcsh_certified_hardware.example.id
}
```
