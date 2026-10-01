---
page_title: "caching_policy.custom_cache_rule"
subcategory: "Load Balancing"
description: "caching_policy.custom_cache_rule for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1217, "body_sha256": "sha256:833ba3231f39698bed81927cf2d5a876f6a3b4ed0945b1e2183af8092e72a4ed", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy", "path": "docs/guides/data-sources--http_loadbalancer--properties--caching_policy--custom_cache_rule.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["caching_policy", "custom_cache_rule"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/caching_policy/custom_cache_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "caching_policy.custom_cache_rule for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy.custom_cache_rule

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [caching_policy](data-sources--http_loadbalancer--properties--caching_policy.md)
- caching_policy.custom_cache_rule

<a id="section"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

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

- [cdn_cache_rules](data-sources--http_loadbalancer--properties--caching_policy--custom_cache_rule--cdn_cache_rules.md): complete subsection reference.

## Next pages

- [caching_policy.custom_cache_rule.cdn_cache_rules](data-sources--http_loadbalancer--properties--caching_policy--custom_cache_rule--cdn_cache_rules.md)
- [caching_policy](data-sources--http_loadbalancer--properties--caching_policy.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
