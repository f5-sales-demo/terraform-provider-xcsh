---
page_title: "rate_limit.custom_ip_allowed_list"
subcategory: "Load Balancing"
description: "rate_limit.custom_ip_allowed_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1577, "body_sha256": "sha256:cc67565a410032af897c919521930c0e37d29852d54e059f3fb6a56cb33eea3a", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit", "path": "documentation/data-sources/http_loadbalancer/properties/rate_limit/custom_ip_allowed_list/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["rate_limit", "custom_ip_allowed_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/rate_limit/custom_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.custom_ip_allowed_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.custom_ip_allowed_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/)
- rate_limit.custom_ip_allowed_list

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

- [rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/): complete subsection reference.

## Next pages

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
