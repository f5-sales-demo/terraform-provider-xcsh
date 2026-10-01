---
page_title: "storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4"
subcategory: ""
description: "storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2526, "body_sha256": "sha256:15f6903040833769f83a41f83f26545b26fd261eb89bc57f72fb2081cc0613df", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address:ipv4", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:nexthop:nexthop_address", "path": "docs/guides/resources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address--ipv4.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "nexthop", "nexthop_address", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/nexthop/nexthop_address/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4 for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [storage_static_routes](resources--fleet--properties--storage_static_routes.md)
- [storage_static_routes.storage_routes](resources--fleet--properties--storage_static_routes--storage_routes.md)
- [storage_static_routes.storage_routes.nexthop](resources--fleet--properties--storage_static_routes--storage_routes--nexthop.md)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address.md)
- storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_static_routes--storage_routes--nexthop--nexthop_address--ipv4--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [storage_static_routes.storage_routes.nexthop.nexthop_address](resources--fleet--properties--storage_static_routes--storage_routes--nexthop--nexthop_address.md)
- [xcsh_fleet](../resources/fleet.md)
