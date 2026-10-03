---
page_title: "custom_network_config.global_network_list.global_network_connections.slo_to_global_dr"
subcategory: ""
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["custom network config global network list global network connections slo to global dr"], "body_bytes": 2284, "body_sha256": "sha256:d0496ac0e2630031e3f71b892b86bf47367b60e357258b8e67849a19f26059a9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections", "path": "documentation/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1322332210320133-2130213000130130-0012122220211010-0002022231001220-1012321010111212-0331322023332033-1231030132311312-0003332320200321", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "slo_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["custom network config global network list global network connections slo to global dr global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr:global_vn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "global_network_list", "global_network_connections", "slo_to_global_dr", "global_vn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list.global_network_connections.slo_to_global_dr

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/)
- [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/)
- custom_network_config.global_network_list.global_network_connections.slo_to_global_dr

<a id="section"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

## Direct properties

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/)
- [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_network_config/global_network_list/global_network_connections/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
