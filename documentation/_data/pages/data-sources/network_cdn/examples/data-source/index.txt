---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_cdn."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1262, "body_sha256": "sha256:c329dfe067afedf15acea8cd1ebaa5e263648d3b0093a97ce28f95f8d3348c80", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_cdn:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e239b2d01c82b76466c9b21dfd9551a3649d04c98c8f874243856b6d5396e241", "source_path": "examples/data-sources/xcsh_network_cdn/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_cdn:example:data-source", "parent_id": "xcsh-docs:data-sources:network_cdn:examples", "path": "documentation/data-sources/network_cdn/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_cdn", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0001220000332333-3022231313010002-2000312113102232-3222102302012321-2330313023300330-2130333320210031-2010202211031130-2200210112230020", "registry_path": "docs/guides/data-sources--network_cdn--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_cdn/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_network_cdn.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/examples/)
- [xcsh_network_cdn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_cdn/)
