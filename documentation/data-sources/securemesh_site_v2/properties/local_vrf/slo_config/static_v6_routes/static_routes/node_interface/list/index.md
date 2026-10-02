---
page_title: "local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list"
subcategory: ""
description: "On a multinode site, this list holds the nodes and corresponding networking_interface."
xcsh_docs: {"aliases": ["local vrf slo config static v6 routes static routes node interface list"], "body_bytes": 3706, "body_sha256": "sha256:6e59a1822156a2601d7c43278bdf553fbb5fc4353726ca62bf54cf55ebc3ac5a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface", "path": "documentation/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "sections": [{"aliases": ["interface"], "anchor": "section", "description": "Interface reference on this node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["node"], "anchor": "schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--node", "description": "Node name on this site.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "node"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/)
- [local_vrf.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/)
- [local_vrf.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/)
- [local_vrf.slo_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/list/interface/): complete subsection reference.

<a id="schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--node"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/list/interface/)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
