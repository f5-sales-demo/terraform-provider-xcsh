---
page_title: "ipv4_prefixes"
subcategory: ""
description: "ipv4_prefixes for xcsh_ip_prefix_set."
xcsh_docs: {"aliases": [], "body_bytes": 2504, "body_sha256": "sha256:fb32e9cb14b754e20b22310e824e785a1aa18d43a3dda5667c607e2354da2bdd", "canonical_id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ip_prefix_set:properties:ipv4_prefixes", "parent_id": "xcsh-docs:data-sources:ip_prefix_set:reference", "path": "docs/guides/data-sources--ip_prefix_set--properties--ipv4_prefixes.md", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipv4_prefixes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/properties/ipv4_prefixes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipv4_prefixes for xcsh_ip_prefix_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipv4_prefixes

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md)
- [Property reference](data-sources--ip_prefix_set--reference.md)
- ipv4_prefixes

<a id="section"></a>

Type: `"list"`. Computed.

IPv4 Prefixes. List of IPv4 prefixes with description.

Upstream description:

List of IPv4 prefixes with description.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-ipv4_prefixes--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Human-readable description text

<a id="schema-ipv4_prefixes--ipv4_prefix"></a>

### ipv4_prefix property

Type: `"string"`. Computed.

IPv4 Prefix. IP address configuration

Upstream description:

IP address configuration

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

## Next pages

- [Property reference](data-sources--ip_prefix_set--reference.md)
- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md)
