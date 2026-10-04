---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_attack_signatures."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1275, "body_sha256": "sha256:52f7adf3d3da57efff90a419b6db0207080a199b2a590f60709d1b15d6831533", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5857e356bd3dea3a26e0c6c17fd28eb9757681e0e4bfa5915066871a8e38e467", "source_path": "examples/data-sources/xcsh_waf_attack_signatures/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_attack_signatures:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_attack_signatures:examples", "path": "documentation/data-sources/waf_attack_signatures/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2311322111003122-1330012212233023-2223001232321231-3021213221010010-3032301301220122-0110233000232021-1322330311323001-3111302312322213", "registry_path": "docs/guides/data-sources--waf_attack_signatures--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_waf_attack_signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_attack_signatures/data-source.tf`; digest `sha256:5857e356bd3dea3a26e0c6c17fd28eb9757681e0e4bfa5915066871a8e38e467`.

```terraform
# WAFAttackSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_attack_signatures" "example" {
}

output "waf_attack_signatures_result" {
  value = data.xcsh_waf_attack_signatures.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/examples/)
- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
