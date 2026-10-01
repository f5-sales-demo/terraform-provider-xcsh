---
page_title: "caching_policy"
subcategory: "Load Balancing"
description: "caching_policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1666, "body_sha256": "sha256:75245faa1154a4c6e4336c05be49d7a335b456be02553041c05b8a7f61452038", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:default_cache_action"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--caching_policy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["caching_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/caching_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "caching_policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- caching_policy

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Upstream description:

Caching Policies for the CDN.

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

OneOf alternatives in this subsection:

- [caching_policy](data-sources--http_loadbalancer--properties--caching_policy.md#section)
- [disable_caching](data-sources--http_loadbalancer--properties--disable_caching.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [custom_cache_rule](data-sources--http_loadbalancer--properties--caching_policy--custom_cache_rule.md): complete subsection reference.

- [default_cache_action](data-sources--http_loadbalancer--properties--caching_policy--default_cache_action.md): complete subsection reference.

## Next pages

- [caching_policy.custom_cache_rule](data-sources--http_loadbalancer--properties--caching_policy--custom_cache_rule.md)
- [caching_policy.default_cache_action](data-sources--http_loadbalancer--properties--caching_policy--default_cache_action.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
