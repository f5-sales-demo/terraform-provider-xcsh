---
page_title: "advanced_options.default_circuit_breaker"
subcategory: "Load Balancing"
description: "advanced_options.default_circuit_breaker for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1072, "body_sha256": "sha256:9b2ae15f37a0d216f0ac42f73aa36071ef01441e6ec02ff876a6e4a977e508d5", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:default_circuit_breaker", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options:default_circuit_breaker", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "path": "docs/guides/data-sources--origin_pool--properties--advanced_options--default_circuit_breaker.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "default_circuit_breaker"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/default_circuit_breaker/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.default_circuit_breaker for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.default_circuit_breaker

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- advanced_options.default_circuit_breaker

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default circuit breaker. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_options](data-sources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
