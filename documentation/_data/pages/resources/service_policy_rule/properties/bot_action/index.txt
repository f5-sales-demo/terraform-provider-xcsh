---
page_title: "bot_action"
subcategory: ""
description: "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing."
xcsh_docs: {"aliases": ["bot action"], "body_bytes": 1481, "body_sha256": "sha256:f6050aab00d8a6e4c8d63d3798c1cd11ffd4a2fd53dcd8e90623f375ced5c52d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "xcsh-docs:resources:service_policy_rule:properties:bot_action:none"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:bot_action", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/bot_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2103123012221321-1103321133011201-1313232020222001-2012201222102131-0300222221031123-2133203022221201-3200132333302332-1223010122321102", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:none", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_action"], "schema_version": 1, "sections": [{"aliases": ["bot action bot skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_action", "bot_skip_processing"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot action none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_action", "none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/bot_action/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_action

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- bot_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_skip_processing",
    "none")}
```

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

Terraform syntax:

```terraform
bot_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/bot_skip_processing/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/none/): complete subsection reference.
