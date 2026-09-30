---
page_title: "allow_list.prefix_list"
subcategory: "Security"
description: "allow_list.prefix_list for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2030, "body_sha256": "sha256:8090cc731f2849e48fa2a7d17b742b18ff7f835be515c9e37c7ac68f8b6b852f", "child_ids": [], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:allow_list:prefix_list", "parent_id": "xcsh-docs:data-sources:service_policy:properties:allow_list", "path": "documentation/data-sources/service_policy/properties/allow_list/prefix_list/index.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["allow_list", "prefix_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/allow_list/prefix_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_list.prefix_list for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allow_list.prefix_list

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/allow_list/)
- allow_list.prefix_list

<a id="section"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="schema-allow_list--prefix_list--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [allow_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/allow_list/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
