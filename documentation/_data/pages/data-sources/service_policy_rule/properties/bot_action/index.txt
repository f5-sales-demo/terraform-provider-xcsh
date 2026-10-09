---
page_title: "bot_action"
subcategory: ""
description: "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing."
xcsh_docs: {"aliases": ["bot action"], "body_bytes": 1171, "body_sha256": "sha256:e687cd734a6a799eaf2c677a34fe57aa4757340070002755935a5a41fbc61872", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:bot_action:bot_skip_processing", "xcsh-docs:data-sources:service_policy_rule:properties:bot_action:none"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:bot_action", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/bot_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_action"], "schema_version": 1, "sections": [{"aliases": ["bot action bot skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:bot_action:bot_skip_processing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_action", "bot_skip_processing"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot action none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:bot_action:none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_action", "none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/bot_action/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_action

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- bot_action

<a id="section"></a>

Type: `"single"`. Computed.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

## Direct properties

- [bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/bot_action/bot_skip_processing/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/bot_action/none/): complete subsection reference.
