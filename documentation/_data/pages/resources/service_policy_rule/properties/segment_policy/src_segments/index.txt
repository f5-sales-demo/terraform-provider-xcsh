---
page_title: "segment_policy.src_segments"
subcategory: ""
description: "List of references to Segments."
xcsh_docs: {"aliases": ["segment policy src segments"], "body_bytes": 1380, "body_sha256": "sha256:fd4cb35ad90a6a0a8a3d8cd1b58d046fc02a19fc00260484f63e83dd4e8b26c9", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy", "path": "documentation/resources/service_policy_rule/properties/segment_policy/src_segments/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0101332320113122-2200203113113133-1130203002313322-0202221012003110-2321113003133011-3130122110120132-1322213110010211-0010122011230310", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_policy.src_segments:RequiredObjectAttributes:segments", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_policy", "src_segments"], "schema_version": 1, "sections": [{"aliases": ["segment policy src segments segments"], "anchor": "section", "description": "Select list of segments.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-segment_policy--src_segments--segments--name", "enforcement": "provider-schema", "group": "segment_policy.src_segments.segments:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:segment_policy:src_segments:segments", "type": "requires"}], "schema_path": ["segment_policy", "src_segments", "segments"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/segment_policy/src_segments/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of references to Segments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_policy.src_segments

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/)
- segment_policy.src_segments

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for src segments.

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
src_segments {
  # Configure direct properties listed below.
}
```

## Direct properties

- [segments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/segment_policy/src_segments/segments/): complete subsection reference.
