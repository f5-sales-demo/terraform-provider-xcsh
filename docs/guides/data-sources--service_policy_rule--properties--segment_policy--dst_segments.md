---
page_title: "segment_policy.dst_segments"
subcategory: ""
description: "segment_policy.dst_segments for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1085, "body_sha256": "sha256:afc37f704558751319ca5deb3d3d27dd5a46681a3ebc528a3ec26f4f3cd7c44d", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:dst_segments", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:dst_segments:segments"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:dst_segments", "parent_id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy", "path": "docs/guides/data-sources--service_policy_rule--properties--segment_policy--dst_segments.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["segment_policy", "dst_segments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/segment_policy/dst_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "segment_policy.dst_segments for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# segment_policy.dst_segments

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- [segment_policy](data-sources--service_policy_rule--properties--segment_policy.md)
- segment_policy.dst_segments

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

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

- [segments](data-sources--service_policy_rule--properties--segment_policy--dst_segments--segments.md): complete subsection reference.

## Next pages

- [segment_policy.dst_segments.segments](data-sources--service_policy_rule--properties--segment_policy--dst_segments--segments.md)
- [segment_policy](data-sources--service_policy_rule--properties--segment_policy.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
