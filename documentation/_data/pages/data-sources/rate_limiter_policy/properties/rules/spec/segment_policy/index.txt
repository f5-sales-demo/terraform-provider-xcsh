---
page_title: "rules.spec.segment_policy"
subcategory: "Security"
description: "Configure source and destination segment for policy."
xcsh_docs: {"aliases": ["rules spec segment policy"], "body_bytes": 3236, "body_sha256": "sha256:d13ae8777dae309e464532f3df5754294acb2ccd0eb8c1d63ed1f5270b1a9d95", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "path": "documentation/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132", "registry_path": "docs/guides/data-sources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec", "segment_policy"], "schema_version": 1, "sections": [{"aliases": ["dst any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "dst_any"], "syntax": "attribute", "type": "object"}, {"aliases": ["dst segments"], "anchor": "section", "description": "List of references to Segments.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "dst_segments"], "syntax": "attribute", "type": "object"}, {"aliases": ["intra segment"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "intra_segment"], "syntax": "attribute", "type": "object"}, {"aliases": ["src any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "src_any"], "syntax": "attribute", "type": "object"}, {"aliases": ["src segments"], "anchor": "section", "description": "List of references to Segments.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "src_segments"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure source and destination segment for policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- rules.spec.segment_policy

<a id="section"></a>

Type: `"single"`. Computed.

Configure source and destination segment for policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

## Direct properties

- [dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_any/): complete subsection reference.

- [dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/): complete subsection reference.

- [intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/): complete subsection reference.

- [src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/src_any/): complete subsection reference.

- [src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/): complete subsection reference.

## Next pages

- [rules.spec.segment_policy.dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_any/)
- [rules.spec.segment_policy.dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/)
- [rules.spec.segment_policy.intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/)
- [rules.spec.segment_policy.src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/src_any/)
- [rules.spec.segment_policy.src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
