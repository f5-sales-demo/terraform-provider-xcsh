---
page_title: "waf_action.app_firewall_detection_control"
subcategory: ""
description: "waf_action.app_firewall_detection_control for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2504, "body_sha256": "sha256:035f9fae56dd24f9939b7176e9fad841fb95942b8cf81352a4237edd9ddb3d11", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action", "path": "docs/guides/resources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_action", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_action.app_firewall_detection_control for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# waf_action.app_firewall_detection_control

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [waf_action](resources--service_policy_rule--properties--waf_action.md)
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

- [exclude_attack_type_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md): complete subsection reference.

- [exclude_bot_name_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md): complete subsection reference.

- [exclude_signature_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md): complete subsection reference.

- [exclude_violation_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md): complete subsection reference.

## Next pages

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md)
- [waf_action](resources--service_policy_rule--properties--waf_action.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
