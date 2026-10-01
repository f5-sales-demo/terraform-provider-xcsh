---
page_title: "waf_action"
subcategory: ""
description: "waf_action for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2019, "body_sha256": "sha256:8da6b5486d96560151fb2130e573ad4d902972a2098cadf0b59f171e96aaf50c", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:none", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:waf_skip_processing"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "docs/guides/data-sources--service_policy_rule--properties--waf_action.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["waf_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/waf_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "waf_action for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- waf_action

<a id="section"></a>

Type: `"single"`. Computed.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

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

## Direct properties

- [app_firewall_detection_control](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md): complete subsection reference.

- [none](data-sources--service_policy_rule--properties--waf_action--none.md): complete subsection reference.

- [waf_skip_processing](data-sources--service_policy_rule--properties--waf_action--waf_skip_processing.md): complete subsection reference.

## Next pages

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--properties--waf_action--app_firewall_detection_control.md)
- [waf_action.none](data-sources--service_policy_rule--properties--waf_action--none.md)
- [waf_action.waf_skip_processing](data-sources--service_policy_rule--properties--waf_action--waf_skip_processing.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
