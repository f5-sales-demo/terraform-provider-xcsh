---
page_title: "site_acl.fast_acl_rules"
subcategory: ""
description: "site_acl.fast_acl_rules for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2528, "body_sha256": "sha256:371b27ff18516dbff666834d3d9a503839c019b67257bd621b21c1dbdd77e73e", "canonical_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:metadata", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:port", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:prefix"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl", "path": "docs/guides/data-sources--fast_acl--properties--site_acl--fast_acl_rules.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_acl", "fast_acl_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/site_acl/fast_acl_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl.fast_acl_rules for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md)
- [Property reference](data-sources--fast_acl--reference.md)
- [site_acl](data-sources--fast_acl--properties--site_acl.md)
- site_acl.fast_acl_rules

<a id="section"></a>

Type: `"list"`. Computed.

Rules. Fast ACL rules to match.

Upstream description:

Fast ACL rules to match.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action.md): complete subsection reference.

- [ip_prefix_set](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set.md): complete subsection reference.

- [metadata](data-sources--fast_acl--properties--site_acl--fast_acl_rules--metadata.md): complete subsection reference.

- [port](data-sources--fast_acl--properties--site_acl--fast_acl_rules--port.md): complete subsection reference.

- [prefix](data-sources--fast_acl--properties--site_acl--fast_acl_rules--prefix.md): complete subsection reference.

## Next pages

- [site_acl.fast_acl_rules.action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action.md)
- [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--properties--site_acl--fast_acl_rules--ip_prefix_set.md)
- [site_acl.fast_acl_rules.metadata](data-sources--fast_acl--properties--site_acl--fast_acl_rules--metadata.md)
- [site_acl.fast_acl_rules.port](data-sources--fast_acl--properties--site_acl--fast_acl_rules--port.md)
- [site_acl.fast_acl_rules.prefix](data-sources--fast_acl--properties--site_acl--fast_acl_rules--prefix.md)
- [site_acl](data-sources--fast_acl--properties--site_acl.md)
- [xcsh_fast_acl](../data-sources/fast_acl.md)
