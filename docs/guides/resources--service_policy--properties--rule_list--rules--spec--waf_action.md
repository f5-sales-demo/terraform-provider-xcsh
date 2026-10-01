---
page_title: "rule_list.rules.spec.waf_action"
subcategory: "Security"
description: "rule_list.rules.spec.waf_action for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2926, "body_sha256": "sha256:fc3d50178c2bb24a8724ae333c718546d5b44b680f11949ffccc22256bff9879", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:none", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing"], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--waf_action.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/waf_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.waf_action for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- rule_list.rules.spec.waf_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "none"),
  validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingObjectAttributes("none",
    "waf_skip_processing")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

Terraform syntax:

```terraform
waf_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_firewall_detection_control](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control.md): complete subsection reference.

- [none](resources--service_policy--properties--rule_list--rules--spec--waf_action--none.md): complete subsection reference.

- [waf_skip_processing](resources--service_policy--properties--rule_list--rules--spec--waf_action--waf_skip_processing.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](resources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control.md)
- [rule_list.rules.spec.waf_action.none](resources--service_policy--properties--rule_list--rules--spec--waf_action--none.md)
- [rule_list.rules.spec.waf_action.waf_skip_processing](resources--service_policy--properties--rule_list--rules--spec--waf_action--waf_skip_processing.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [xcsh_service_policy](../resources/service_policy.md)
