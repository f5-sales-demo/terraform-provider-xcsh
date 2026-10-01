---
page_title: "rule_list.rules.spec.waf_action"
subcategory: "Security"
description: "rule_list.rules.spec.waf_action for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2482, "body_sha256": "sha256:5b0ecb9a21bed780f7bf5045138d0f42880fb3f2e069015dcaca67bcfa07b6be", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "path": "docs/guides/data-sources--service_policy--properties--rule_list--rules--spec--waf_action.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.waf_action for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- [rule_list](data-sources--service_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md)
- rule_list.rules.spec.waf_action

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

- [app_firewall_detection_control](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control.md): complete subsection reference.

- [none](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--none.md): complete subsection reference.

- [waf_skip_processing](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--waf_skip_processing.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--app_firewall_detection_control.md)
- [rule_list.rules.spec.waf_action.none](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--none.md)
- [rule_list.rules.spec.waf_action.waf_skip_processing](data-sources--service_policy--properties--rule_list--rules--spec--waf_action--waf_skip_processing.md)
- [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
