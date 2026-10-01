---
page_title: "cache_rules.cache_bypass"
subcategory: ""
description: "cache_rules.cache_bypass for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1008, "body_sha256": "sha256:6b71467a21acc2011ce38efc67a2a334007d9db8505874c6372d32f31f21fdf1", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:cache_bypass", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:cache_bypass", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "path": "docs/guides/resources--cdn_cache_rule--properties--cache_rules--cache_bypass.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "cache_bypass"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/cache_bypass/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.cache_bypass for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.cache_bypass

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- [Property reference](resources--cdn_cache_rule--reference.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- cache_rules.cache_bypass

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for cache bypass.

Upstream description:

This can be used for messages where no values are needed.

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
cache_bypass = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
