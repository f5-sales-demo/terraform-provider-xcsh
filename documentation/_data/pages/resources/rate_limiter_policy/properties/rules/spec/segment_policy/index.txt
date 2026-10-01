---
page_title: "rules.spec.segment_policy"
subcategory: "Security"
description: "rules.spec.segment_policy for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3706, "body_sha256": "sha256:62ae02a3d86040c0bf391bf4d00af07df1f0f1cfa3e713f4e0616b19ba01bf3f", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments"], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "documentation/resources/rate_limiter_policy/properties/rules/spec/segment_policy/index.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["rules", "spec", "segment_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.segment_policy for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/)
- rules.spec.segment_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure source and destination segment for policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dst_any",
    "dst_segments"),
  validators.ConflictingObjectAttributes("dst_any",
    "intra_segment"),
  validators.ConflictingObjectAttributes("dst_segments",
    "intra_segment"),
  validators.ConflictingObjectAttributes("src_any",
    "src_segments")}
```

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

Terraform syntax:

```terraform
segment_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_any/): complete subsection reference.

- [dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/): complete subsection reference.

- [intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/): complete subsection reference.

- [src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_any/): complete subsection reference.

- [src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/): complete subsection reference.

## Next pages

- [rules.spec.segment_policy.dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_any/)
- [rules.spec.segment_policy.dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/)
- [rules.spec.segment_policy.intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/intra_segment/)
- [rules.spec.segment_policy.src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_any/)
- [rules.spec.segment_policy.src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_segments/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/rules/spec/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
