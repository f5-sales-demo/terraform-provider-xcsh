---
page_title: "segment_vrf.segment_config.static_v6_routes"
subcategory: ""
description: "List of IPv6 static routes."
xcsh_docs: {"aliases": ["segment vrf segment config static v6 routes"], "body_bytes": 1300, "body_sha256": "sha256:ccb64ea67a8f646297c61a00f21eeacc84e4b551a71a945c36cac94bf2cb97a1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config", "path": "documentation/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1303120210133032-0020301133000323-3311321220022302-3113222021233210-2213313002331302-1100333233332222-0121333220032113-2133332330221312", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_v6_routes"], "schema_version": 1, "sections": [{"aliases": ["segment vrf segment config static v6 routes static routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_v6_routes:static_routes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["segment_vrf", "segment_config", "static_v6_routes", "static_routes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_v6_routes/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "List of IPv6 static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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

Additional upstream details:

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
