---
page_title: "custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list"
subcategory: ""
description: "custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3333, "body_sha256": "sha256:a800a94fa9fe72d1187b9d6c2d3849dd5f4227f51a7f5d047066e4a2137b352f", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list:interface"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes:node_interface", "path": "docs/guides/data-sources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.slo_config](data-sources--voltstack_site--properties--custom_network_config--slo_config.md)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes.md)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes.md)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface.md)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

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

- [interface](data-sources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md): complete subsection reference.

<a id="schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--node"></a>

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

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface.md)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--properties--custom_network_config--slo_config--static_v6_routes--static_routes--node_interface.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
