---
page_title: "waf_action.app_firewall_detection_control"
subcategory: ""
description: "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria."
xcsh_docs: {"aliases": ["waf action app firewall detection control"], "body_bytes": 3252, "body_sha256": "sha256:86c0fa4555fa0f3fcd033afcf06cf4692d7d65e42ebbe3224e6ab4b44acd9588", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action", "path": "documentation/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2033330131132212-0031300303303301-2320030332203121-0211321133320003-2032302122131013-0103100312101110-1002121021012102-0121023131122100", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_action", "app_firewall_detection_control"], "schema_version": 1, "sections": [{"aliases": ["waf action app firewall detection control exclude attack type contexts"], "anchor": "section", "description": "Exclude an entire attack type only in the named context. For migrated per-parameter exceptions, prefer this over signature-ID exclusions because one payload can trigger several signatures; unrelated parameters and attack types remain protected.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_attack_type_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["waf action app firewall detection control exclude bot name contexts"], "anchor": "section", "description": "Bot Names to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "enforcement": "provider-schema", "group": "waf_action.app_firewall_detection_control.exclude_bot_name_contexts:RequiredListObjectAttributes:bot_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "type": "requires"}], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["waf action app firewall detection control exclude signature contexts"], "anchor": "section", "description": "Signature IDs to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-waf_action--app_firewall_detection_control--exclude_signature_contexts--signature_id", "enforcement": "provider-schema", "group": "waf_action.app_firewall_detection_control.exclude_signature_contexts:RequiredListObjectAttributes:signature_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "type": "requires"}], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_signature_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["waf action app firewall detection control exclude violation contexts"], "anchor": "section", "description": "Violations to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_violation_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control", "exclude_violation_contexts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action.app_firewall_detection_control

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/)
- waf_action.app_firewall_detection_control

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/)
- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
