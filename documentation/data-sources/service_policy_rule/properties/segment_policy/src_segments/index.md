---
page_title: "segment_policy.src_segments"
subcategory: ""
description: "List of references to Segments."
xcsh_docs: {"aliases": ["segment policy src segments"], "body_bytes": 1539, "body_sha256": "sha256:9a91b793173440e58d9147b35abe1b5e8b28de1649c4a6167721652aba1810c7", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:src_segments:segments"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:src_segments", "parent_id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy", "path": "documentation/data-sources/service_policy_rule/properties/segment_policy/src_segments/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1220001122222130-0303323031120210-0203112221301021-0121332332221201-3321313332210022-1023022113320133-1123213021331003-3121302033103123", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_policy", "src_segments"], "schema_version": 1, "sections": [{"aliases": ["segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:segment_policy:src_segments:segments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["segment_policy", "src_segments", "segments"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/segment_policy/src_segments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy.src_segments

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/)
- segment_policy.src_segments

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

- [segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/src_segments/segments/): complete subsection reference.

## Next pages

- [segment_policy.src_segments.segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/src_segments/segments/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/segment_policy/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
