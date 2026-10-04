---
page_title: "primary.dnssec_mode.enable"
subcategory: "DNS"
description: "DNSSEC enable."
xcsh_docs: {"aliases": ["primary dnssec mode enable"], "body_bytes": 1299, "body_sha256": "sha256:46a26e7c1bac2c5217eba1e8da5bdd078af3024fc011c646195a4e82489de05e", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "path": "documentation/resources/dns_zone/properties/primary/dnssec_mode/enable/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1310313012332122-1030002133303121-1003102013010323-3332112323301223-3110312030122203-0312021133201032-1100323202001132-2131130301303301", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "dnssec_mode", "enable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/dnssec_mode/enable/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "DNSSEC enable.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.dnssec_mode.enable

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.dnssec_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/)
- primary.dnssec_mode.enable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable. DNSSEC enable.

Upstream description:

DNSSEC enable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [primary.dnssec_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
