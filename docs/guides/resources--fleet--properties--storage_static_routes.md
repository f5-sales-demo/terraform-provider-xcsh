---
page_title: "storage_static_routes"
subcategory: ""
description: "storage_static_routes for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1239, "body_sha256": "sha256:77ffbbe9fc481a8bf0ff9cde94589e1651079130ec9b0f68014e3074f940ae45", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_static_routes", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--storage_static_routes.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- storage_static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage static routes.

Upstream description:

List of storage static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_routes")}
```

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
storage_static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_routes](resources--fleet--properties--storage_static_routes--storage_routes.md): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes](resources--fleet--properties--storage_static_routes--storage_routes.md)
- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
