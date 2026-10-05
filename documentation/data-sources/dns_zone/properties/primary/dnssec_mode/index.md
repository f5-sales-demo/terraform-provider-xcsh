---
page_title: "primary.dnssec_mode"
subcategory: "DNS"
description: "DNSSEC Mode."
xcsh_docs: {"aliases": ["primary dnssec mode"], "body_bytes": 1670, "body_sha256": "sha256:271cf5fc3143ec4ffbc4421c4e4bdb66c8077f4c9ce52055a76eeaa9f7d76d75", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode:disable_spec", "xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode:enable"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode", "parent_id": "xcsh-docs:data-sources:dns_zone:properties:primary", "path": "documentation/data-sources/dns_zone/properties/primary/dnssec_mode/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102", "registry_path": "docs/guides/data-sources--dns_zone--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "dnssec_mode"], "schema_version": 1, "sections": [{"aliases": ["primary dnssec mode disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "dnssec_mode", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["primary dnssec mode enable"], "anchor": "section", "description": "DNSSEC enable.", "document_id": "xcsh-docs:data-sources:dns_zone:properties:primary:dnssec_mode:enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "dnssec_mode", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/primary/dnssec_mode/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "DNSSEC Mode.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.dnssec_mode

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- primary.dnssec_mode

<a id="section"></a>

Type: `"single"`. Computed.

DNSSEC Mode.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/enable/): complete subsection reference.

## Next pages

- [primary.dnssec_mode.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/disable_spec/)
- [primary.dnssec_mode.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/enable/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
