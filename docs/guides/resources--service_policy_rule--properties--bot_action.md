---
page_title: "bot_action"
subcategory: ""
description: "bot_action for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1598, "body_sha256": "sha256:8ffd68062c8457e8ecd13917027dcbdca36ff21b4bd99a011809bca1ea7b9f37", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "xcsh-docs:resources:service_policy_rule:properties:bot_action:none"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:bot_action", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "docs/guides/resources--service_policy_rule--properties--bot_action.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/bot_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_action for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_action

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- bot_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_skip_processing",
    "none")}
```

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

Terraform syntax:

```terraform
bot_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bot_skip_processing](resources--service_policy_rule--properties--bot_action--bot_skip_processing.md): complete subsection reference.

- [none](resources--service_policy_rule--properties--bot_action--none.md): complete subsection reference.

## Next pages

- [bot_action.bot_skip_processing](resources--service_policy_rule--properties--bot_action--bot_skip_processing.md)
- [bot_action.none](resources--service_policy_rule--properties--bot_action--none.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
