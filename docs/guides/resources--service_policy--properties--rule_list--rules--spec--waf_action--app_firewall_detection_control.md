---
page_title: "rule_list.rules.spec.waf_action.app_firewall_detection_control"
subcategory: "Security"
description: "rule_list.rules.spec.waf_action.app_firewall_detection_control for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3174, "body_sha256": "sha256:89a5803571f69a0881d9638ab72ff782e05952567f517a709a3323c2527f3301", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.waf_action.app_firewall_detection_control for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action.app_firewall_detection_control

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [rule_list.rules.spec.waf_action](resources--service_policy--properties--rule_list--rules--spec--waf_action.md)
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

- [exclude_attack_type_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md): complete subsection reference.

- [exclude_bot_name_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md): complete subsection reference.

- [exclude_signature_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts.md): complete subsection reference.

- [exclude_violation_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_signature_contexts.md)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control--exclude_violation_contexts.md)
- [rule_list.rules.spec.waf_action](resources--service_policy--properties--rule_list--rules--spec--waf_action.md)
- [xcsh_service_policy](../resources/service_policy.md)
