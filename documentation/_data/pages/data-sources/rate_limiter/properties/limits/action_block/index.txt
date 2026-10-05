---
page_title: "limits.action_block"
subcategory: "Security"
description: "Action where a user is blocked from making further requests after exceeding rate limit threshold."
xcsh_docs: {"aliases": ["limits action block"], "body_bytes": 2122, "body_sha256": "sha256:f713a96078b0ad90ff0d7ce74ecc83c0af33911fc6008e9c7069d7f3c8417a59", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:hours", "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:minutes", "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:seconds"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block", "parent_id": "xcsh-docs:data-sources:rate_limiter:properties:limits", "path": "documentation/data-sources/rate_limiter/properties/limits/action_block/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3221130221020103-0312010202000323-2021222110132133-0013331101311010-2012013120020130-0323321210003020-2032322033000210-3211121201221012", "registry_path": "docs/guides/data-sources--rate_limiter--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["limits", "action_block"], "schema_version": 1, "sections": [{"aliases": ["limits action block hours"], "anchor": "section", "description": "Input Duration Hours.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:hours", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "hours"], "syntax": "attribute", "type": "object"}, {"aliases": ["limits action block minutes"], "anchor": "section", "description": "Input Duration Minutes.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:minutes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "minutes"], "syntax": "attribute", "type": "object"}, {"aliases": ["limits action block seconds"], "anchor": "section", "description": "Input Duration Seconds.", "document_id": "xcsh-docs:data-sources:rate_limiter:properties:limits:action_block:seconds", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["limits", "action_block", "seconds"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter/properties/limits/action_block/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.action_block

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/)
- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/)
- limits.action_block

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

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/minutes/): complete subsection reference.

- [seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/seconds/): complete subsection reference.

## Next pages

- [limits.action_block.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/hours/)
- [limits.action_block.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/minutes/)
- [limits.action_block.seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/action_block/seconds/)
- [limits](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/properties/limits/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter/)
