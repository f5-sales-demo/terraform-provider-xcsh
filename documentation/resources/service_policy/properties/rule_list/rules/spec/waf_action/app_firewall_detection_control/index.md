---
page_title: "rule_list.rules.spec.waf_action.app_firewall_detection_control"
subcategory: "Security"
description: "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria."
xcsh_docs: {"aliases": ["rule list rules spec waf action app firewall detection control"], "body_bytes": 3943, "body_sha256": "sha256:76f79552a840b36e889b2637dab0060bb6a5ed4c017f1cda7787162b57b2b46f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3022321321013201-0020033302122032-2311011301300323-1202110132110022-1233301221333112-0030120230100322-3103323111103001-1210013301233011", "registry_path": "docs/guides/resources--service_policy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control"], "schema_version": 1, "sections": [{"aliases": ["exclude attack type contexts"], "anchor": "section", "description": "Exclude an entire attack type only in the named context. For migrated per-parameter exceptions, prefer this over signature-ID exclusions because one payload can trigger several signatures; unrelated parameters and attack types remain protected.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control", "exclude_attack_type_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["exclude bot name contexts"], "anchor": "section", "description": "Bot Names to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts--bot_name", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts:RequiredListObjectAttributes:bot_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control", "exclude_bot_name_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["exclude signature contexts"], "anchor": "section", "description": "Signature IDs to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_signature_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts--signature_id", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts:RequiredListObjectAttributes:signature_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_signature_contexts", "type": "requires"}], "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control", "exclude_signature_contexts"], "syntax": "block", "type": "object"}, {"aliases": ["exclude violation contexts"], "anchor": "section", "description": "Violations to be excluded for the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_violation_contexts", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control", "exclude_violation_contexts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action.app_firewall_detection_control

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/)
- rule_list.rules.spec.waf_action.app_firewall_detection_control

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

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/)
- [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
