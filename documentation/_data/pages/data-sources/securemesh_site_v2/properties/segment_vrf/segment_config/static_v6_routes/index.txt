---
page_title: "segment_vrf.segment_config.static_v6_routes"
subcategory: ""
description: "List of IPv6 static routes."
xcsh_docs: {"aliases": ["segment vrf segment config static v6 routes"], "body_bytes": 1807, "body_sha256": "sha256:c97ca2c6ae1bdbfe53ac0fdcedcebba63ed717a6bc2491c45e2809aeaf993dc7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config", "path": "documentation/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_v6_routes"], "schema_version": 1, "sections": [{"aliases": ["segment vrf segment config static v6 routes static routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["segment_vrf", "segment_config", "static_v6_routes", "static_routes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "List of IPv6 static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config.static_v6_routes

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [segment_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/)
- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- segment_vrf.segment_config.static_v6_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/): complete subsection reference.

## Next pages

- [segment_vrf.segment_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/static_routes/)
- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
