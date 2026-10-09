---
page_title: "rules.spec.segment_policy.dst_segments"
subcategory: "Security"
description: "List of references to Segments."
xcsh_docs: {"aliases": ["rules spec segment policy dst segments"], "body_bytes": 1395, "body_sha256": "sha256:5abd1db557e1b116cd6e1a2151e9e54e67ef8db7163029ffdcb1a15571550952", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments:segments"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy", "path": "documentation/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3001212202030123-2011032303010332-3011112032220233-3222313111213313-0223122011203130-1311231101202311-2200012000010031-1313320010021221", "registry_path": "docs/guides/data-sources--rate_limiter_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "spec", "segment_policy", "dst_segments"], "schema_version": 1, "sections": [{"aliases": ["rules spec segment policy dst segments segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy:dst_segments:segments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "spec", "segment_policy", "dst_segments", "segments"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec.segment_policy.dst_segments

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- [rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/)
- [rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/)
- rules.spec.segment_policy.dst_segments

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Additional upstream details:

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

- [segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/dst_segments/segments/): complete subsection reference.
