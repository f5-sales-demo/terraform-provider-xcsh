---
page_title: "rate_limit.rate_limiter.action_block.minutes"
subcategory: "Load Balancing"
description: "rate_limit.rate_limiter.action_block.minutes for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2353, "body_sha256": "sha256:8206ec04212db86577e26346231e56d5ffa09fd8d50a0d542afca30e3f54d968", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "path": "docs/guides/resources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block--minutes.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block", "minutes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.rate_limiter.action_block.minutes for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block.minutes

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [rate_limit](resources--http_loadbalancer--properties--rate_limit.md)
- [rate_limit.rate_limiter](resources--http_loadbalancer--properties--rate_limit--rate_limiter.md)
- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block.md)
- rate_limit.rate_limiter.action_block.minutes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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
minutes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rate_limit--rate_limiter--action_block--minutes--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
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
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

## Next pages

- [rate_limit.rate_limiter.action_block](resources--http_loadbalancer--properties--rate_limit--rate_limiter--action_block.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
