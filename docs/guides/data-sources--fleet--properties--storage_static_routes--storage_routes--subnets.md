---
page_title: "storage_static_routes.storage_routes.subnets"
subcategory: ""
description: "storage_static_routes.storage_routes.subnets for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2020, "body_sha256": "sha256:ce81fbc715863710e337fd8a87bac7467d55c4e5f2284d042c139c200545549d", "canonical_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets", "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv4", "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv6"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes", "path": "docs/guides/data-sources--fleet--properties--storage_static_routes--storage_routes--subnets.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.subnets for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_static_routes.storage_routes.subnets

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md)
- [Property reference](data-sources--fleet--reference.md)
- [storage_static_routes](data-sources--fleet--properties--storage_static_routes.md)
- [storage_static_routes.storage_routes](data-sources--fleet--properties--storage_static_routes--storage_routes.md)
- storage_static_routes.storage_routes.subnets

<a id="section"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

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

## Direct properties

- [ipv4](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv4.md): complete subsection reference.

- [ipv6](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv6.md): complete subsection reference.

## Next pages

- [storage_static_routes.storage_routes.subnets.ipv4](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv4.md)
- [storage_static_routes.storage_routes.subnets.ipv6](data-sources--fleet--properties--storage_static_routes--storage_routes--subnets--ipv6.md)
- [storage_static_routes.storage_routes](data-sources--fleet--properties--storage_static_routes--storage_routes.md)
- [xcsh_fleet](../data-sources/fleet.md)
