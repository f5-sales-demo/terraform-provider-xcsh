---
page_title: "rule_list.rules.spec.segment_policy"
subcategory: "Security"
description: "rule_list.rules.spec.segment_policy for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3140, "body_sha256": "sha256:ae0a6641869bfeb5588fd16d9bc0e32686699a6784421902ca3f6fc0e9bdf430", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_any", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:intra_segment", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:src_any", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments"], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--segment_policy.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "segment_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.segment_policy for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.segment_policy

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- rule_list.rules.spec.segment_policy

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

- [dst_any](resources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_any.md): complete subsection reference.

- [dst_segments](resources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_segments.md): complete subsection reference.

- [intra_segment](resources--service_policy--properties--rule_list--rules--spec--segment_policy--intra_segment.md): complete subsection reference.

- [src_any](resources--service_policy--properties--rule_list--rules--spec--segment_policy--src_any.md): complete subsection reference.

- [src_segments](resources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.segment_policy.dst_any](resources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_any.md)
- [rule_list.rules.spec.segment_policy.dst_segments](resources--service_policy--properties--rule_list--rules--spec--segment_policy--dst_segments.md)
- [rule_list.rules.spec.segment_policy.intra_segment](resources--service_policy--properties--rule_list--rules--spec--segment_policy--intra_segment.md)
- [rule_list.rules.spec.segment_policy.src_any](resources--service_policy--properties--rule_list--rules--spec--segment_policy--src_any.md)
- [rule_list.rules.spec.segment_policy.src_segments](resources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [xcsh_service_policy](../resources/service_policy.md)
