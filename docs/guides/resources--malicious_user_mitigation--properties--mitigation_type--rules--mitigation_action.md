---
page_title: "mitigation_type.rules.mitigation_action"
subcategory: ""
description: "mitigation_type.rules.mitigation_action for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 2730, "body_sha256": "sha256:cc78c8866c28449a96d6a8df3d693c933a6941d90874dbb9e494a20b9d6b100b", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge"], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "docs/guides/resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type", "rules", "mitigation_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type.rules.mitigation_action for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules.mitigation_action

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
- [Property reference](resources--malicious_user_mitigation--reference.md)
- [mitigation_type](resources--malicious_user_mitigation--properties--mitigation_type.md)
- [mitigation_type.rules](resources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- mitigation_type.rules.mitigation_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Supported actions that can be taken to mitigate malicious activity from a user.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block_temporarily",
    "captcha_challenge"),
  validators.ConflictingObjectAttributes("block_temporarily",
    "javascript_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge",
    "javascript_challenge")}
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
  "x-ves-oneof-field-mitigation_action": "[\"block_temporarily\",\"captcha_challenge\",\"javascript_challenge\"]"
}
```

Terraform syntax:

```terraform
mitigation_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block_temporarily](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--block_temporarily.md): complete subsection reference.

- [captcha_challenge](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--captcha_challenge.md): complete subsection reference.

- [javascript_challenge](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--javascript_challenge.md): complete subsection reference.

## Next pages

- [mitigation_type.rules.mitigation_action.block_temporarily](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--block_temporarily.md)
- [mitigation_type.rules.mitigation_action.captcha_challenge](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--captcha_challenge.md)
- [mitigation_type.rules.mitigation_action.javascript_challenge](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--javascript_challenge.md)
- [mitigation_type.rules](resources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
