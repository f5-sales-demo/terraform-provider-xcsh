---
page_title: "api_rate_limit.custom_ip_allowed_list"
subcategory: "Load Balancing"
description: "api_rate_limit.custom_ip_allowed_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1258, "body_sha256": "sha256:e3b5caec08bba7bd1c1590b2a3f69be390a7a5d44fb060b616242ceabf9db843", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_rate_limit--custom_ip_allowed_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "custom_ip_allowed_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.custom_ip_allowed_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.custom_ip_allowed_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- api_rate_limit.custom_ip_allowed_list

<a id="section"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

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

- [rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--properties--api_rate_limit--custom_ip_allowed_list--rate_limiter_allowed_prefixes.md): complete subsection reference.

## Next pages

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--properties--api_rate_limit--custom_ip_allowed_list--rate_limiter_allowed_prefixes.md)
- [api_rate_limit](data-sources--http_loadbalancer--properties--api_rate_limit.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
