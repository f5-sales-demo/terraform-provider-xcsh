---
page_title: "waf_action.app_firewall_detection_control"
subcategory: ""
description: "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria."
xcsh_docs: {"aliases": ["waf action app firewall detection control"], "body_bytes": 1910, "body_sha256": "sha256:45ce680555b84cebde6c7db45c86a60462e8f121c7cf03f86fddeecc7b2bf6ba", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action", "path": "documentation/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1100223002122311-3102112101011001-3323030310110133-2001211223312132-0323301301012311-0321211203211032-2320021103030123-3323102223202123", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_action", "app_firewall_detection_control"], "schema_version": 1, "sections": [{"aliases": ["waf action app firewall detection control exclude attack type contexts"], "anchor": "section", "description": "Exclude an entire attack type only in the named context. For migrated per-parameter exceptions, prefer this over signature-ID exclusions because one payload can trigger several signatures; unrelated parameters and attack types remain protected.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_attack_type_contexts"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf action app firewall detection control exclude bot name contexts"], "anchor": "section", "description": "Bot Names to be excluded for the defined match criteria.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf action app firewall detection control exclude signature contexts"], "anchor": "section", "description": "Signature IDs to be excluded for the defined match criteria.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_signature_contexts"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf action app firewall detection control exclude violation contexts"], "anchor": "section", "description": "Violations to be excluded for the defined match criteria.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_violation_contexts", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_violation_contexts"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action.app_firewall_detection_control

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/)
- waf_action.app_firewall_detection_control

<a id="section"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.
