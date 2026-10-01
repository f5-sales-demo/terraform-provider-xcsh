---
page_title: "rule_list.rules.spec.bot_action"
subcategory: "Security"
description: "rule_list.rules.spec.bot_action for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1734, "body_sha256": "sha256:74356d27c115a588011e66f7240a6c90e0f27a125e9ac2766b489aecf4f86b13", "canonical_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:bot_action", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:bot_action:bot_skip_processing", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:bot_action:none"], "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:bot_action", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "path": "docs/guides/data-sources--service_policy--properties--rule_list--rules--spec--bot_action.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "bot_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/bot_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.bot_action for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.bot_action

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md)
- [Property reference](data-sources--service_policy--reference.md)
- [rule_list](data-sources--service_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md)
- rule_list.rules.spec.bot_action

<a id="section"></a>

Type: `"single"`. Computed.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

## Direct properties

- [bot_skip_processing](data-sources--service_policy--properties--rule_list--rules--spec--bot_action--bot_skip_processing.md): complete subsection reference.

- [none](data-sources--service_policy--properties--rule_list--rules--spec--bot_action--none.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.bot_action.bot_skip_processing](data-sources--service_policy--properties--rule_list--rules--spec--bot_action--bot_skip_processing.md)
- [rule_list.rules.spec.bot_action.none](data-sources--service_policy--properties--rule_list--rules--spec--bot_action--none.md)
- [rule_list.rules.spec](data-sources--service_policy--properties--rule_list--rules--spec.md)
- [xcsh_service_policy](../data-sources/service_policy.md)
