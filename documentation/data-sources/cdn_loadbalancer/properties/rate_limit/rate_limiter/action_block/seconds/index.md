---
page_title: "rate_limit.rate_limiter.action_block.seconds"
subcategory: "Load Balancing"
description: "Input Duration Seconds."
xcsh_docs: {"aliases": ["rate limit rate limiter action block seconds"], "body_bytes": 2057, "body_sha256": "sha256:bc2d9e66fa3cb8db2c9db1b7a37c31469f2cdbab828cdb7c63e2205bba55f8ac", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "path": "documentation/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/seconds/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3231033012222010-2120232321333112-0312012110302313-2013223121230200-3132113103010121-3220000003010221-1221021310321010-2213212111200310", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block", "seconds"], "schema_version": 1, "sections": [{"aliases": ["rate limit rate limiter action block seconds duration"], "anchor": "schema-rate_limit--rate_limiter--action_block--seconds--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "seconds", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/seconds/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Input Duration Seconds.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block.seconds

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/)
- [rate_limit.rate_limiter.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/)
- rate_limit.rate_limiter.action_block.seconds

<a id="section"></a>

Type: `"single"`. Computed.

Seconds. Input Duration Seconds.

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

<a id="schema-rate_limit--rate_limiter--action_block--seconds--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```
