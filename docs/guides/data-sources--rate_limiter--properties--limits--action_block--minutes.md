---
page_title: "limits.action_block.minutes"
subcategory: "Security"
description: "limits.action_block.minutes for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1754, "body_sha256": "sha256:af5ac372925495f80a46e8b92ac2847137518ee13a3fc9efbb936d16edb85c3c", "canonical_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:minutes", "child_ids": [], "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:minutes", "parent_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block", "path": "docs/guides/data-sources--rate_limiter--properties--limits--action_block--minutes.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["limits", "action_block", "minutes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/properties/limits/action_block/minutes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "limits.action_block.minutes for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# limits.action_block.minutes

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md)
- [Property reference](data-sources--rate_limiter--reference.md)
- [limits](data-sources--rate_limiter--properties--limits.md)
- [limits.action_block](data-sources--rate_limiter--properties--limits--action_block.md)
- limits.action_block.minutes

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-limits--action_block--minutes--duration"></a>

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

- [limits.action_block](data-sources--rate_limiter--properties--limits--action_block.md)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md)
