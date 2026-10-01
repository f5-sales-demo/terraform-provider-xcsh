---
page_title: "segment_policy"
subcategory: ""
description: "segment_policy for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2568, "body_sha256": "sha256:9856ea06b4a821ee3c46dd8c98244eae0d641e4858613271fb67bd8efe4db280", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_any", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:dst_segments", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:intra_segment", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_any", "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "docs/guides/resources--service_policy_rule--properties--segment_policy.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_policy for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- segment_policy

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

- [dst_any](resources--service_policy_rule--properties--segment_policy--dst_any.md): complete subsection reference.

- [dst_segments](resources--service_policy_rule--properties--segment_policy--dst_segments.md): complete subsection reference.

- [intra_segment](resources--service_policy_rule--properties--segment_policy--intra_segment.md): complete subsection reference.

- [src_any](resources--service_policy_rule--properties--segment_policy--src_any.md): complete subsection reference.

- [src_segments](resources--service_policy_rule--properties--segment_policy--src_segments.md): complete subsection reference.

## Next pages

- [segment_policy.dst_any](resources--service_policy_rule--properties--segment_policy--dst_any.md)
- [segment_policy.dst_segments](resources--service_policy_rule--properties--segment_policy--dst_segments.md)
- [segment_policy.intra_segment](resources--service_policy_rule--properties--segment_policy--intra_segment.md)
- [segment_policy.src_any](resources--service_policy_rule--properties--segment_policy--src_any.md)
- [segment_policy.src_segments](resources--service_policy_rule--properties--segment_policy--src_segments.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
