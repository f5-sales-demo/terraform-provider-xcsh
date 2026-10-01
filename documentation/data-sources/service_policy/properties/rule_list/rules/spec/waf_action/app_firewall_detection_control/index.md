---
page_title: "rule_list.rules.spec.waf_action.app_firewall_detection_control"
subcategory: "Security"
description: "rule_list.rules.spec.waf_action.app_firewall_detection_control for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3852, "body_sha256": "sha256:d78feeef9d850ff600c06b516b5e0227e163acf162377fbf7bf4697063ab2b89", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action", "path": "documentation/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/index.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.waf_action.app_firewall_detection_control for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action.app_firewall_detection_control

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/)
- rule_list.rules.spec.waf_action.app_firewall_detection_control

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

- [exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/): complete subsection reference.

- [exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/): complete subsection reference.

- [exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/): complete subsection reference.

- [exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_signature_contexts/)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/exclude_violation_contexts/)
- [rule_list.rules.spec.waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
