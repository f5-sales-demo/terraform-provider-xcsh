---
page_title: "custom_network_config.slo_config.static_routes.static_routes.node_interface.list"
subcategory: ""
description: "custom_network_config.slo_config.static_routes.static_routes.node_interface.list for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 3384, "body_sha256": "sha256:5fd16c13e6a1042652b2f6a24c67bc44b1b8218a891051f8d8e67183a6ca3732", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list:interface"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface:list", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_routes", "static_routes", "node_interface", "list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.slo_config.static_routes.static_routes.node_interface.list for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_routes.static_routes.node_interface.list

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [custom_network_config.slo_config](resources--securemesh_site--properties--custom_network_config--slo_config.md)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes.md)
- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes.md)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface.md)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

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

- [interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md): complete subsection reference.

<a id="schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--node"></a>

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

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface.md)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--properties--custom_network_config--slo_config--static_routes--static_routes--node_interface.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
