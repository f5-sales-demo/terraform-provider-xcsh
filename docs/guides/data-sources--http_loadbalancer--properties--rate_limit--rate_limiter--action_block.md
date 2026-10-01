---
page_title: "rate_limit.rate_limiter.action_block"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter.action_block for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1974, "body_sha256": "sha256:1dc0132a588120ce420629bbba628ff57f9ee06e3295eb1dcc807643f5fb01b7", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter", "path": "docs/guides/data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter.action_block for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [rate_limit](data-sources--http_loadbalancer--properties--rate_limit.md)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter.md)
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

- [hours](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block--hours.md): complete subsection reference.

- [minutes](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block--minutes.md): complete subsection reference.

- [seconds](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block--seconds.md): complete subsection reference.

## Next pages

- [rate_limit.rate_limiter.action_block.hours](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block--hours.md)
- [rate_limit.rate_limiter.action_block.minutes](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block--minutes.md)
- [rate_limit.rate_limiter.action_block.seconds](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block--seconds.md)
- [rate_limit.rate_limiter](data-sources--http_loadbalancer--properties--rate_limit--rate_limiter.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
