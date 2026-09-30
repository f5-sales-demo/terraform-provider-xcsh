---
page_title: "rate_limit.rate_limiter.action_block.hours"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter.action_block.hours for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2005, "body_sha256": "sha256:f5e1a488331ef1be7425e3dd798a84596afc3272cefd8f2217432b5768e8386d", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block--hours.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block", "hours"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/hours/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter.action_block.hours for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rate_limit.rate_limiter.action_block.hours

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [rate_limit](data-sources--cdn_loadbalancer--properties--rate_limit.md)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter.md)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block.md)
- rate_limit.rate_limiter.action_block.hours

<a id="section"></a>

Type: `"single"`. Computed.

Hours. Input Duration Hours.

Upstream description:

Input Duration Hours.

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

<a id="schema-rate_limit--rate_limiter--action_block--hours--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

## Next pages

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--properties--rate_limit--rate_limiter--action_block.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
