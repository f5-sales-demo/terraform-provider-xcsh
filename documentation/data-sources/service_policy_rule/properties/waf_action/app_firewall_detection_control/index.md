---
page_title: "waf_action.app_firewall_detection_control"
subcategory: ""
description: "waf_action.app_firewall_detection_control for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3152, "body_sha256": "sha256:1cc51451f2c6f89b6bc01035042ae2d0e9a5df68d910e8d37248d6e48a164385", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action", "path": "documentation/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["waf_action", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_action.app_firewall_detection_control for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_attack_type_contexts/)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_bot_name_contexts/)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_signature_contexts/)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/exclude_violation_contexts/)
- [waf_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
