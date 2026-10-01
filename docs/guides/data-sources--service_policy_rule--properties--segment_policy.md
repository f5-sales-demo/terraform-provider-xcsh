---
page_title: "segment_policy"
subcategory: ""
description: "segment_policy for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2092, "body_sha256": "sha256:4badd85c1feffea48053d40e7344dcc460e0c20f351b2b39a703e88a112931c5", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:dst_any", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:dst_segments", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:intra_segment", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:src_any", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:src_segments"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "docs/guides/data-sources--service_policy_rule--properties--segment_policy.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_policy for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- segment_policy

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

- [dst_any](data-sources--service_policy_rule--properties--segment_policy--dst_any.md): complete subsection reference.

- [dst_segments](data-sources--service_policy_rule--properties--segment_policy--dst_segments.md): complete subsection reference.

- [intra_segment](data-sources--service_policy_rule--properties--segment_policy--intra_segment.md): complete subsection reference.

- [src_any](data-sources--service_policy_rule--properties--segment_policy--src_any.md): complete subsection reference.

- [src_segments](data-sources--service_policy_rule--properties--segment_policy--src_segments.md): complete subsection reference.

## Next pages

- [segment_policy.dst_any](data-sources--service_policy_rule--properties--segment_policy--dst_any.md)
- [segment_policy.dst_segments](data-sources--service_policy_rule--properties--segment_policy--dst_segments.md)
- [segment_policy.intra_segment](data-sources--service_policy_rule--properties--segment_policy--intra_segment.md)
- [segment_policy.src_any](data-sources--service_policy_rule--properties--segment_policy--src_any.md)
- [segment_policy.src_segments](data-sources--service_policy_rule--properties--segment_policy--src_segments.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
