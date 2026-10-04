---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_zone_cryptokeys."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1253, "body_sha256": "sha256:8061bfd903f1a6e7b459770893d156f65908539cb9cf7276726f3dc584362f3c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b7a0c08b1e74769c7e06fd735c8df6c0ed430cc8d5516a712859e61adfe28bc8", "source_path": "examples/data-sources/xcsh_dns_zone_cryptokeys/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:examples", "path": "documentation/data-sources/dns_zone_cryptokeys/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0221323200003000-0201121202321133-1030213300123123-1121031321031132-0022112333000122-2121303111222113-3101312100231313-1123200101000122", "registry_path": "docs/guides/data-sources--dns_zone_cryptokeys--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_dns_zone_cryptokeys.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_zone_cryptokeys/data-source.tf`; digest `sha256:b7a0c08b1e74769c7e06fd735c8df6c0ed430cc8d5516a712859e61adfe28bc8`.

```terraform
# DNSZoneCryptokeys DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_dns_zone_cryptokeys" "example" {
}

output "dns_zone_cryptokeys_result" {
  value = data.xcsh_dns_zone_cryptokeys.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/examples/)
- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
