---
page_title: "segment_vrf.segment_config.static_routes.static_routes.node_interface.list"
subcategory: ""
description: "On a multinode site, this list holds the nodes and corresponding networking_interface."
xcsh_docs: {"aliases": ["segment vrf segment config static routes static routes node interface list"], "body_bytes": 3764, "body_sha256": "sha256:1d1579340b9ab18bdb23cbaff5f882acdb7ec172643958c0d65ad92bef30cc8f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface:list:interface"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface", "path": "documentation/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2113202010113230-3032001331112030-3120300302011222-3020201303330131-0212311302213032-0113001101100232-0322210002022111-3322002001103121", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "sections": [{"aliases": ["segment vrf segment config static routes static routes node interface list interface"], "anchor": "section", "description": "Interface reference on this node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface:list:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["segment_vrf", "segment_config", "static_routes", "static_routes", "node_interface", "list", "interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["segment vrf segment config static routes static routes node interface list node"], "anchor": "schema-segment_vrf--segment_config--static_routes--static_routes--node_interface--list--node", "description": "Node name on this site.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["segment_vrf", "segment_config", "static_routes", "static_routes", "node_interface", "list", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config.static_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [segment_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/)
- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- [segment_vrf.segment_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/)
- [segment_vrf.segment_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/)
- segment_vrf.segment_config.static_routes.static_routes.node_interface.list

<a id="section"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

## Direct properties

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/list/interface/): complete subsection reference.

<a id="schema-segment_vrf--segment_config--static_routes--static_routes--node_interface--list--node"></a>

### node property

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [segment_vrf.segment_config.static_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/list/interface/)
- [segment_vrf.segment_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/node_interface/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
