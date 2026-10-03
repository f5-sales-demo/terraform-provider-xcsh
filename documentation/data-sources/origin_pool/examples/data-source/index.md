---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_origin_pool."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1266, "body_sha256": "sha256:dca569fa24034fc7454f92591ea3a03cc6447b8ae1f3b8f0944e6418061ad3c0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cae409ea2c1cde6f7ffac297039f556a77e9684a94017d350d26cb02cd85782a", "source_path": "examples/data-sources/xcsh_origin_pool/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:origin_pool:example:data-source", "parent_id": "xcsh-docs:data-sources:origin_pool:examples", "path": "documentation/data-sources/origin_pool/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3122032320030203-3303233222333212-2330222003120311-3332022103233112-3111101123220313-3313010302313031-1303233000223200-0202121233002013", "registry_path": "docs/guides/data-sources--origin_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_origin_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_origin_pool/data-source.tf`; digest `sha256:cae409ea2c1cde6f7ffac297039f556a77e9684a94017d350d26cb02cd85782a`.

```terraform
# OriginPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing OriginPool by name
data "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}

output "origin_pool_id" {
  value = data.xcsh_origin_pool.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/examples/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
