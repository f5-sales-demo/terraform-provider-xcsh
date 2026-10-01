---
page_title: "local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list"
subcategory: ""
description: "local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 3169, "body_sha256": "sha256:630a75362cc49779a0cb3fb6d7383b91f36675d9679c7d01dac49389aec0d8fb", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface:list", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface:list:interface"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:local_vrf:sli_config:static_v6_routes:static_routes:node_interface", "path": "docs/guides/data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--node_interface--list.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_vrf", "sli_config", "static_v6_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/local_vrf/sli_config/static_v6_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [local_vrf](data-sources--securemesh_site_v2--properties--local_vrf.md)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--properties--local_vrf--sli_config.md)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes.md)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes.md)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--node_interface.md)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

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

- [interface](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md): complete subsection reference.

<a id="schema-local_vrf--sli_config--static_v6_routes--static_routes--node_interface--list--node"></a>

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

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--node_interface--list--interface.md)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--properties--local_vrf--sli_config--static_v6_routes--static_routes--node_interface.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
