---
page_title: "rule_list.rules.spec.segment_policy.src_segments"
subcategory: "Security"
description: "rule_list.rules.spec.segment_policy.src_segments for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1587, "body_sha256": "sha256:b89666410663124cb3483bb197a05beb71314b531a9b2b12ff74aa11f9f49cac", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments:segments"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy", "path": "docs/guides/data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "segment_policy", "src_segments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.segment_policy.src_segments for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.segment_policy.src_segments

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- [rule_list](data-sources--service_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy.md)
- rule_list.rules.spec.segment_policy.src_segments

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for src segments.

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

- [segments](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments--segments.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.segment_policy.src_segments.segments](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy--src_segments--segments.md)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--properties--rule_list--rules--spec--segment_policy.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
