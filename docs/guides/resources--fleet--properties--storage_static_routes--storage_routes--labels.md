---
page_title: "storage_static_routes.storage_routes.labels"
subcategory: ""
description: "storage_static_routes.storage_routes.labels for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1128, "body_sha256": "sha256:da322affe203f425886279876b9c29c887d33f81aa38345b7be9e18f77be1259", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:labels", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:labels", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "path": "docs/guides/resources--fleet--properties--storage_static_routes--storage_routes--labels.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.labels for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.labels

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_static_routes](resources--fleet--properties--storage_static_routes.md)
- [storage_static_routes.storage_routes](resources--fleet--properties--storage_static_routes--storage_routes.md)
- storage_static_routes.storage_routes.labels

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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

Terraform syntax:

```terraform
labels {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [storage_static_routes.storage_routes](resources--fleet--properties--storage_static_routes--storage_routes.md)
- [xcsh_fleet](../resources/fleet.md)
