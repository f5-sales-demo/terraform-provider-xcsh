---
page_title: "storage_static_routes.storage_routes.subnets.ipv4"
subcategory: ""
description: "storage_static_routes.storage_routes.subnets.ipv4 for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 3035, "body_sha256": "sha256:0fd23d6fd15254d7c0cdf7bc3d6b52205d1b9095d269e1d466e5fef88958d435", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv4", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_static_routes:storage_routes:subnets", "path": "documentation/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/ipv4/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "subnets", "ipv4"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/ipv4/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_static_routes.storage_routes.subnets.ipv4 for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# storage_static_routes.storage_routes.subnets.ipv4

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/)
- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/)
- [storage_static_routes.storage_routes.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/)
- storage_static_routes.storage_routes.subnets.ipv4

<a id="section"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

## Direct properties

<a id="schema-storage_static_routes--storage_routes--subnets--ipv4--plen"></a>

### plen property

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="schema-storage_static_routes--storage_routes--subnets--ipv4--prefix"></a>

### prefix property

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

- [storage_static_routes.storage_routes.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_static_routes/storage_routes/subnets/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
