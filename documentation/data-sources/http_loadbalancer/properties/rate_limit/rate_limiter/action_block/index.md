---
page_title: "rate_limit.rate_limiter.action_block"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter.action_block for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2568, "body_sha256": "sha256:7306c55ca89805d66ee1e2658a220f982dcf0cbbfaa196de687ad1e4956c527a", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter", "path": "documentation/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter.action_block for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/)
- rate_limit.rate_limiter.action_block

<a id="section"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/): complete subsection reference.

- [seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/seconds/): complete subsection reference.

## Next pages

- [rate_limit.rate_limiter.action_block.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/hours/)
- [rate_limit.rate_limiter.action_block.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/)
- [rate_limit.rate_limiter.action_block.seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/seconds/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
