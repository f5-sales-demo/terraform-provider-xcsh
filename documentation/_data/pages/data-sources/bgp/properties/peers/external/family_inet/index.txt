---
page_title: "peers.external.family_inet"
subcategory: ""
description: "Parameters for inet family."
xcsh_docs: {"aliases": ["peers external family inet"], "body_bytes": 1890, "body_sha256": "sha256:2623c00225a129b3565093662eef3cac9f68ecac18800926874782c8071a62c4", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:disable_spec", "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external", "path": "documentation/data-sources/bgp/properties/peers/external/family_inet/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "family_inet"], "schema_version": 1, "sections": [{"aliases": ["peers external family inet disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "family_inet", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers external family inet enable"], "anchor": "section", "description": "IPv4 Unicast.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["peers", "external", "family_inet", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/family_inet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters for inet family.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgpCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- peers.external.family_inet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/): complete subsection reference.

## Next pages

- [peers.external.family_inet.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/disable_spec/)
- [peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
