---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_cdn."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1037, "body_sha256": "sha256:7199ad5d6c4531b86713670ef2102b3342ad92ee1295dfd9275660d0ae2a49dd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_cdn:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e239b2d01c82b76466c9b21dfd9551a3649d04c98c8f874243856b6d5396e241", "source_path": "examples/data-sources/xcsh_network_cdn/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_cdn:example:data-source", "parent_id": "xcsh-docs:data-sources:network_cdn:examples", "path": "documentation/data-sources/network_cdn/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_cdn", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0001220000332333-3022231313010002-2000312113102232-3222102302012321-2330313023300330-2130333320210031-2010202211031130-2200210112230020", "registry_path": "docs/guides/data-sources--network_cdn--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_cdn/examples/data-source/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data source for xcsh_network_cdn.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_cdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_cdn/data-source.tf`; digest `sha256:e239b2d01c82b76466c9b21dfd9551a3649d04c98c8f874243856b6d5396e241`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_cdn" "origin_ingress" {}

output "cdn_https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_cdn.origin_ingress.cidr_blocks
  }
}
```
