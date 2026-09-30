---
page_title: "mitigation_type.rules.mitigation_action"
subcategory: ""
description: "mitigation_type.rules.mitigation_action for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 2175, "body_sha256": "sha256:0f518a799c9cca88511c5056992e85d11b65b1c265c36eaba0da570848a48df5", "canonical_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "child_ids": ["xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:block_temporarily", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:javascript_challenge"], "collection_id": "xcsh-docs:data-sources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "parent_id": "xcsh-docs:data-sources:malicious_user_mitigation:properties:mitigation_type:rules", "path": "docs/guides/data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type", "rules", "mitigation_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type.rules.mitigation_action for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# mitigation_type.rules.mitigation_action

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md)
- [Property reference](data-sources--malicious_user_mitigation--reference.md)
- [mitigation_type](data-sources--malicious_user_mitigation--properties--mitigation_type.md)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- mitigation_type.rules.mitigation_action

<a id="section"></a>

Type: `"single"`. Computed.

Supported actions that can be taken to mitigate malicious activity from a user.

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

## Direct properties

- [block_temporarily](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--block_temporarily.md): complete subsection reference.

- [captcha_challenge](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--captcha_challenge.md): complete subsection reference.

- [javascript_challenge](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--javascript_challenge.md): complete subsection reference.

## Next pages

- [mitigation_type.rules.mitigation_action.block_temporarily](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--block_temporarily.md)
- [mitigation_type.rules.mitigation_action.captcha_challenge](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--captcha_challenge.md)
- [mitigation_type.rules.mitigation_action.javascript_challenge](data-sources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--javascript_challenge.md)
- [mitigation_type.rules](data-sources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md)
