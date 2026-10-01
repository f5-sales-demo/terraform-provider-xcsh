---
page_title: "routes.simple_route.advanced_options.common_hash_policy"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.common_hash_policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1313, "body_sha256": "sha256:9478062339c40c66f952cd6539dc8adcd7acc1cd8537382d5df28eb232e040a7", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_hash_policy", "child_ids": [], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_hash_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--common_hash_policy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "common_hash_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/common_hash_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.common_hash_policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.common_hash_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](data-sources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- routes.simple_route.advanced_options.common_hash_policy

<a id="section"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
