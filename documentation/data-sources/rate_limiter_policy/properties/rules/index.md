---
page_title: "rules"
subcategory: "Security"
description: "rules for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2268, "body_sha256": "sha256:b17d54b9d22fb701a569acf243e39d1d8e44d9190e46cd4856fa360ba954fd38", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:metadata", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec"], "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:reference", "path": "documentation/data-sources/rate_limiter_policy/properties/rules/index.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- rules

<a id="section"></a>

Type: `"list"`. Computed.

List of RateLimiterRules that are evaluated sequentially till a matching rule is identified.
Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/metadata/): complete subsection reference.

- [spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/): complete subsection reference.

## Next pages

- [rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/metadata/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
