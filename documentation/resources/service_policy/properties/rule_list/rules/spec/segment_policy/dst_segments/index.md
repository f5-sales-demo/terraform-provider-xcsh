---
page_title: "rule_list.rules.spec.segment_policy.dst_segments"
subcategory: "Security"
description: "List of references to Segments."
xcsh_docs: {"aliases": ["rule list rules spec segment policy dst segments"], "body_bytes": 1848, "body_sha256": "sha256:12736e8b390dc7f6aa5809571ee35871b552dfb458bd6dbde62c04d2caaba18f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments:segments"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2023102300100100-0010302223303330-1300202232022033-0103023130302032-1023012301210211-3320130333031013-2033320130101133-1131333303332132", "registry_path": "docs/guides/resources--service_policy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy.dst_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments:segments", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "segment_policy", "dst_segments"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec segment policy dst segments segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments:segments", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-rule_list--rules--spec--segment_policy--dst_segments--segments--name", "enforcement": "provider-schema", "group": "rule_list.rules.spec.segment_policy.dst_segments.segments:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:segment_policy:dst_segments:segments", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "segment_policy", "dst_segments", "segments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["service_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.segment_policy.dst_segments

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [rule_list.rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/)
- rule_list.rules.spec.segment_policy.dst_segments

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dst segments.

Additional upstream details:

List of references to Segments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("segments")}
```

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

Terraform syntax:

```terraform
dst_segments {
  # Configure direct properties listed below.
}
```

## Direct properties

- [segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/segment_policy/dst_segments/segments/): complete subsection reference.
