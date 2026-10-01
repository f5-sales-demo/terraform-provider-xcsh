---
page_title: "waf_action"
subcategory: ""
description: "waf_action for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2472, "body_sha256": "sha256:4e2571064a2a6c055b233152d4cefaa56d2ad6ce33ebe742eb6d63a2396bc3d4", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "xcsh-docs:resources:service_policy_rule:properties:waf_action:none", "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:waf_action", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "docs/guides/resources--service_policy_rule--properties--waf_action.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/waf_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_action for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- waf_action

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

- [app_firewall_detection_control](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md): complete subsection reference.

- [none](resources--service_policy_rule--properties--waf_action--none.md): complete subsection reference.

- [waf_skip_processing](resources--service_policy_rule--properties--waf_action--waf_skip_processing.md): complete subsection reference.

## Next pages

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md)
- [waf_action.none](resources--service_policy_rule--properties--waf_action--none.md)
- [waf_action.waf_skip_processing](resources--service_policy_rule--properties--waf_action--waf_skip_processing.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
