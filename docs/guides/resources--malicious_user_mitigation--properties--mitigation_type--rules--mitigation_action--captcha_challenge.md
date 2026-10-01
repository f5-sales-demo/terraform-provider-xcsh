---
page_title: "mitigation_type.rules.mitigation_action.captcha_challenge"
subcategory: ""
description: "mitigation_type.rules.mitigation_action.captcha_challenge for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 1470, "body_sha256": "sha256:bf9b96749c702df706d1822af6677a24f0cf5c2d5eed2230766aea3d8489a714", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "child_ids": [], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action:captcha_challenge", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules:mitigation_action", "path": "docs/guides/resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action--captcha_challenge.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mitigation_type", "rules", "mitigation_action", "captcha_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/rules/mitigation_action/captcha_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mitigation_type.rules.mitigation_action.captcha_challenge for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type.rules.mitigation_action.captcha_challenge

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
- [Property reference](resources--malicious_user_mitigation--reference.md)
- [mitigation_type](resources--malicious_user_mitigation--properties--mitigation_type.md)
- [mitigation_type.rules](resources--malicious_user_mitigation--properties--mitigation_type--rules.md)
- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md)
- mitigation_type.rules.mitigation_action.captcha_challenge

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for captcha challenge.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
captcha_challenge = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [mitigation_type.rules.mitigation_action](resources--malicious_user_mitigation--properties--mitigation_type--rules--mitigation_action.md)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
