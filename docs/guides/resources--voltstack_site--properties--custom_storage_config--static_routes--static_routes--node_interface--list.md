---
page_title: "custom_storage_config.static_routes.static_routes.node_interface.list"
subcategory: ""
description: "custom_storage_config.static_routes.static_routes.node_interface.list for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3096, "body_sha256": "sha256:79afa67cd9bda7c16316f0ee18bdaae214aca8e091ad49c194f523f1a46b609e", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface:list", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface:list:interface"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes:static_routes:node_interface", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--static_routes--static_routes--node_interface--list.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "static_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/static_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.static_routes.static_routes.node_interface.list for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.static_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.static_routes](resources--voltstack_site--properties--custom_storage_config--static_routes.md)
- [custom_storage_config.static_routes.static_routes](resources--voltstack_site--properties--custom_storage_config--static_routes--static_routes.md)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--properties--custom_storage_config--static_routes--static_routes--node_interface.md)
- custom_storage_config.static_routes.static_routes.node_interface.list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [interface](resources--voltstack_site--properties--custom_storage_config--static_routes--static_routes--node_interface--list--interface.md): complete subsection reference.

<a id="schema-custom_storage_config--static_routes--static_routes--node_interface--list--node"></a>

### node property

Type: `"string"`. Optional.

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

- [custom_storage_config.static_routes.static_routes.node_interface.list.interface](resources--voltstack_site--properties--custom_storage_config--static_routes--static_routes--node_interface--list--interface.md)
- [custom_storage_config.static_routes.static_routes.node_interface](resources--voltstack_site--properties--custom_storage_config--static_routes--static_routes--node_interface.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
