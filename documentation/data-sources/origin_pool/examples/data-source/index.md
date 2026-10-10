---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_origin_pool."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1041, "body_sha256": "sha256:4f1c50c9682754a4bbc7d4e6e2140ee5430a00acd57f980824109c3fd56897bc", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cae409ea2c1cde6f7ffac297039f556a77e9684a94017d350d26cb02cd85782a", "source_path": "examples/data-sources/xcsh_origin_pool/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:origin_pool:example:data-source", "parent_id": "xcsh-docs:data-sources:origin_pool:examples", "path": "documentation/data-sources/origin_pool/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3122032320030203-3303233222333212-2330222003120311-3332022103233112-3111101123220313-3313010302313031-1303233000223200-0202121233002013", "registry_path": "docs/guides/data-sources--origin_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_origin_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["origin_poolCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
