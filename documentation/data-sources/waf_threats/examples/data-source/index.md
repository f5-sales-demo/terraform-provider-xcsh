---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_threats."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 941, "body_sha256": "sha256:97f4706125e66821adeddc614699db985b2ada783cafe8dbe5f44d017663afe3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9ac0cdd709cbc9a9ca710dea9bf40dd5b343227de03ab5599937418e7cc2c055", "source_path": "examples/data-sources/xcsh_waf_threats/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_threats:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_threats:examples", "path": "documentation/data-sources/waf_threats/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_threats", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2202312212030212-0201011323231132-1103233031011022-3010302202012333-1211332010302213-0131321311030321-3000213121010220-0210013231310221", "registry_path": "docs/guides/data-sources--waf_threats--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_waf_threats.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_waf_threats](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_threats/data-source.tf`; digest `sha256:9ac0cdd709cbc9a9ca710dea9bf40dd5b343227de03ab5599937418e7cc2c055`.

```terraform
# WAFThreats DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threats" "example" {
}

output "waf_threats_result" {
  value = data.xcsh_waf_threats.example
}
```
