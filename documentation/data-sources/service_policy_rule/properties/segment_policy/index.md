---
page_title: "segment_policy"
subcategory: ""
description: "segment_policy for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2800, "body_sha256": "sha256:82ca5702a382d759037f955dfce4ee43b27ed6154e9c68c30c6feda0c100163f", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:dst_any", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:dst_segments", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:intra_segment", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:src_any", "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:src_segments"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/segment_policy/index.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["segment_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/segment_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_policy for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
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

- [dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/dst_any/): complete subsection reference.

- [dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/dst_segments/): complete subsection reference.

- [intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/intra_segment/): complete subsection reference.

- [src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/src_any/): complete subsection reference.

- [src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/src_segments/): complete subsection reference.

## Next pages

- [segment_policy.dst_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/dst_any/)
- [segment_policy.dst_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/dst_segments/)
- [segment_policy.intra_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/intra_segment/)
- [segment_policy.src_any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/src_any/)
- [segment_policy.src_segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/src_segments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
