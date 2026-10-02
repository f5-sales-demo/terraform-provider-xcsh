---
page_title: "rule_list.rules.spec.segment_policy.src_segments"
subcategory: "Security"
description: "List of references to Segments."
xcsh_docs: {"aliases": ["rule list rules spec segment policy src segments"], "body_bytes": 2080, "body_sha256": "sha256:0b96c42805c9beb8ae18c7eab9646f96b3e98922a2a1bd5f0fcac6cd384c5342", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments:segments"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy", "path": "documentation/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233", "registry_path": "docs/guides/data-sources--service_policy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "segment_policy", "src_segments"], "schema_version": 1, "sections": [{"aliases": ["segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:segment_policy:src_segments:segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "segment_policy", "src_segments", "segments"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.segment_policy.src_segments

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- [rule_list.rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/)
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

- [segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/segments/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.segment_policy.src_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/src_segments/segments/)
- [rule_list.rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/segment_policy/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
