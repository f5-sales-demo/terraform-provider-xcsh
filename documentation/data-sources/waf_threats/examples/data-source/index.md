---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_threats."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1166, "body_sha256": "sha256:d841754b8c1b9cade9c340024c138ed584dd3e9a018ec82a09bc8245d3fca439", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threats:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9ac0cdd709cbc9a9ca710dea9bf40dd5b343227de03ab5599937418e7cc2c055", "source_path": "examples/data-sources/xcsh_waf_threats/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_threats:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_threats:examples", "path": "documentation/data-sources/waf_threats/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_threats", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2202312212030212-0201011323231132-1103233031011022-3010302202012333-1211332010302213-0131321311030321-3000213121010220-0210013231310221", "registry_path": "docs/guides/data-sources--waf_threats--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threats/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_waf_threats.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/examples/)
- [xcsh_waf_threats](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threats/)
