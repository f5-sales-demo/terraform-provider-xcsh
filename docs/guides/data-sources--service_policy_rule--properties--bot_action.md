---
page_title: "bot_action"
subcategory: ""
description: "bot_action for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1231, "body_sha256": "sha256:df02c6a6545dbd8a7e399d80a4f71876f8dd60cb2bd704f31dc7f9224bd5f3f1", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:properties:bot_action", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:bot_action:bot_skip_processing", "xcsh-docs:data-sources:service_policy_rule:properties:bot_action:none"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:bot_action", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "docs/guides/data-sources--service_policy_rule--properties--bot_action.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/bot_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_action for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_action

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- bot_action

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

- [bot_skip_processing](data-sources--service_policy_rule--properties--bot_action--bot_skip_processing.md): complete subsection reference.

- [none](data-sources--service_policy_rule--properties--bot_action--none.md): complete subsection reference.

## Next pages

- [bot_action.bot_skip_processing](data-sources--service_policy_rule--properties--bot_action--bot_skip_processing.md)
- [bot_action.none](data-sources--service_policy_rule--properties--bot_action--none.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
