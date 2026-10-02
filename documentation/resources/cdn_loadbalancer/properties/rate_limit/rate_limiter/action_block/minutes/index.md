---
page_title: "rate_limit.rate_limiter.action_block.minutes"
subcategory: "Load Balancing"
description: "Input Duration Minutes."
xcsh_docs: {"aliases": ["rate limit rate limiter action block minutes"], "body_bytes": 2698, "body_sha256": "sha256:00f0f86060bb91921bd2c0ef7c037e209ec9cb0a524277a9ef9dc7c2c7faf0e0", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block", "path": "documentation/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2232102003211200-3310003313003312-0012200212113100-0002323313331101-1020102203000121-1302032211020003-3000302231323112-1301000013201333", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block", "minutes"], "schema_version": 1, "sections": [{"aliases": ["duration"], "anchor": "schema-rate_limit--rate_limiter--action_block--minutes--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "minutes", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Input Duration Minutes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block.minutes

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/)
- [rate_limit.rate_limiter.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/)
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

- [rate_limit.rate_limiter.action_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/rate_limit/rate_limiter/action_block/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
