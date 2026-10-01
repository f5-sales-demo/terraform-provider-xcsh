---
page_title: "waf_action.app_firewall_detection_control"
subcategory: ""
description: "waf_action.app_firewall_detection_control for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2503, "body_sha256": "sha256:099262d1b2ca30c777b6cd3d131d3045a317325bcb9e371b7c58d4e0ba1ff6fb", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_attack_type_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_bot_name_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_signature_contexts", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control:exclude_violation_contexts"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "parent_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action", "path": "docs/guides/data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_action", "app_firewall_detection_control"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_action.app_firewall_detection_control for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action.app_firewall_detection_control

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- [waf_action](data-sources--service_policy_rule--properties--waf_action.md)
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

- [exclude_attack_type_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md): complete subsection reference.

- [exclude_signature_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md): complete subsection reference.

- [exclude_violation_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md): complete subsection reference.

## Next pages

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_attack_type_contexts.md)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_bot_name_contexts.md)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_signature_contexts.md)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control--exclude_violation_contexts.md)
- [waf_action](data-sources--service_policy_rule--properties--waf_action.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
