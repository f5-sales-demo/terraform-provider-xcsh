---
page_title: "rules.spec.segment_policy"
subcategory: "Security"
description: "rules.spec.segment_policy for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2920, "body_sha256": "sha256:72a2f93382312a3011d2ae45b71533c7d373fcb3de8c4f36c50ae521ce297711", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_any", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:intra_segment", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_segments"], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--segment_policy.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "segment_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.segment_policy for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
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

- [dst_any](resources--rate_limiter_policy--properties--rules--spec--segment_policy--dst_any.md): complete subsection reference.

- [dst_segments](resources--rate_limiter_policy--properties--rules--spec--segment_policy--dst_segments.md): complete subsection reference.

- [intra_segment](resources--rate_limiter_policy--properties--rules--spec--segment_policy--intra_segment.md): complete subsection reference.

- [src_any](resources--rate_limiter_policy--properties--rules--spec--segment_policy--src_any.md): complete subsection reference.

- [src_segments](resources--rate_limiter_policy--properties--rules--spec--segment_policy--src_segments.md): complete subsection reference.

## Next pages

- [rules.spec.segment_policy.dst_any](resources--rate_limiter_policy--properties--rules--spec--segment_policy--dst_any.md)
- [rules.spec.segment_policy.dst_segments](resources--rate_limiter_policy--properties--rules--spec--segment_policy--dst_segments.md)
- [rules.spec.segment_policy.intra_segment](resources--rate_limiter_policy--properties--rules--spec--segment_policy--intra_segment.md)
- [rules.spec.segment_policy.src_any](resources--rate_limiter_policy--properties--rules--spec--segment_policy--src_any.md)
- [rules.spec.segment_policy.src_segments](resources--rate_limiter_policy--properties--rules--spec--segment_policy--src_segments.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
