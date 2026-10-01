---
page_title: "storage_static_routes.storage_routes.subnets"
subcategory: ""
description: "storage_static_routes.storage_routes.subnets for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2365, "body_sha256": "sha256:3816e521f2110ab64e63106feb69f962064108218e98555fbc426f2cbc377ee6", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv4", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv6"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "path": "docs/guides/resources--fleet--properties--storage_static_routes--storage_routes--subnets.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.subnets for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.subnets

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_static_routes](resources--fleet--properties--storage_static_routes.md)
- [storage_static_routes.storage_routes](resources--fleet--properties--storage_static_routes--storage_routes.md)
- storage_static_routes.storage_routes.subnets

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](resources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv4.md): complete subsection reference.

- [ipv6](resources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv6.md): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.subnets.ipv4](resources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv4.md)
- [storage_static_routes.storage_routes.subnets.ipv6](resources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv6.md)
- [storage_static_routes.storage_routes](resources--fleet--properties--storage_static_routes--storage_routes.md)
- [xcsh_fleet](../resources/fleet.md)
