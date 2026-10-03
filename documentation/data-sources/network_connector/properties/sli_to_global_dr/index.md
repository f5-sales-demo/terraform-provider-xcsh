---
page_title: "sli_to_global_dr"
subcategory: "Networking"
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["sli to global dr"], "body_bytes": 1961, "body_sha256": "sha256:fd03a6039cc1d21b84b43f8626de55bec40cb0c0d69a677fd68ebf51ac7134eb", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "documentation/data-sources/network_connector/properties/sli_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3202003000033330-0132321012323121-2201002113322223-1301301113020232-1303210303223002-1110011130333032-0202031330113122-0100012123123232", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sli_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["sli to global dr global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:network_connector:properties:sli_to_global_dr:global_vn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sli_to_global_dr", "global_vn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_connectorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sli_to_global_dr

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- sli_to_global_dr

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: sli\_to\_global\_dr, sli\_to\_slo\_snat, slo\_to\_global\_dr\] Global network reference for
direct connection.

Upstream description:

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

OneOf alternatives in this subsection:

- [sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_global_dr/#section)
- [sli_to_slo_snat](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_slo_snat/#section)
- [slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/slo_to_global_dr/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/sli_to_global_dr/global_vn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
